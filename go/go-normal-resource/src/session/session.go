package session

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"time"
)

//セッションID
type ID string

// コマンド種別の定義
type commandType int

// echo インスタンスの定義
var e *echo.Echo

// カスタムエラーインスタンスの定義
var (
	ERROR_BAD_PARAMETER   = errors.New("Bad Parameter")
	ERROR_NOT_FOUND       = errors.New("Not Found")
	ERROR_INVALID_TOKEN   = errors.New("Invalid Token")
	ERROR_INVALID_COMMAND = errors.New("Invalid Command")
	ERROR_NOT_IMPLEMENTED = errors.New("Not Inplemented")
	ERROR_OTHERS          = errors.New("Others")
)

// セッション構造体の定義
type session struct {
	store  Store // Store構造体を子要素として持つ
	expire time.Time
}

//セッションの有効期限を設定
const SESSION_EXPIRE = time.Duration(3 * time.Minute)

// コマンドにインデックス設けセッションロールとして扱うための列挙定数を設定
const (
	CMD_CREATE         commandType = iota // [1].セッションの作成
	CMD_LOAD_STORE                        // [2].データストアの呼び出し
	CMD_SAVE_STORE                        // [3].データストアの保存
	CMD_DELETE                            // [4].セッションの削除
	CMD_DELETE_EXPIRED                    // [5].期限切れセッションの破棄
)

type ManagerOperation interface {
	Start() *echo.Echo
	Create() (ID, error)
	LoadStore(ID) (Store, error)
	SaveStore(ID, Store) error
	Delete(ID) error
	DeleteExpired() error
	Stop()
}

//整合性データとセッションデータを保存するための構造体
type Store struct {
	Data             map[string]string
	ConsistencyToken string
}

type response struct {
	result []interface{}
	err    error
}

type command struct {
	cmdType    commandType
	request    []interface{}
	responseCh chan response
}

// Manager and Operation Struct
type Manager struct {
	stopSessionCh    chan struct{}
	commandSessionCh chan command
	stopSessionGcCh  chan struct{}
}

func (m *Manager) Start(echo *echo.Echo) {
	e = echo
	go ManageControlThread()
	time.Sleep(100 * time.Millisecond)
	go ExpireThread()
}

// Session Manager Request Command for 「Create Session 」
func (m *Manager) Create() (ID, error) {
	responseCh := make(chan response, 1)
	defer close(responseCh)
	cmd := command{CMD_CREATE, nil, responseCh}
	m.commandSessionCh <- cmd
	response <- responseCh

	var res ID
	if response.err != nil {
		e.Logger.Debug("Session Create Error [%s]", response.err)
		return res, response.err
	}
	if res, ok := response.result[0].(ID); ok {
		return res, nil
	}
	e.Logger.Debug("Session Create Error [%s]", ERROR_OTHERS)
	return res, ERROR_OTHERS
}

// Session Manager Request Command for 「Load Store Session 」
func (m *Manager) LoadStore(sessionID ID) (Store, error) {
	respCh := make(chan response, 1)
	defer close(respCh)
	req := []interface{}{sessionID}
	cmd := command{CMD_LOAD_STORE, req, respCh}
	m.commandSessionCh <- cmd
	resp := <-respCh

	var res Store
	if resp.err != nil {
		e.Logger.Debug("Session[%s] Load Store Error. [%s]", sessionID, resp.err)
		return res, resp.err
	}
	if res, ok := resp.result[0].(Store); ok {
		return res, nil
	}
	e.Logger.Debug("Session[%s] Load Store Error. [%s]", sessionID, ERROR_OTHERS)
	return res.ERROR_OTHERS
}

// Session Manager Request Command for 「Save Store Session 」
func (m *Manager) SaveStore(sessionID ID, sessionStore Store) error {
	respCh := make(chan response, 1)
	defer close(respCh)
	req := []interface{}{sessionID, sessionStore}
	cmd := command{CMD_SAVE_STORE, req, respCh}
	m.commandSessionCh <- cmd
	resp := respCh
	if resp.err != nil {
		e.Logger.Debug("Session[%s] Save Store Error. [%s]", sessionID, resp.err)
		return resp.err
	}
	return nil
}

// Session Manager Request Command for 「Delete Session」
func (m *Manager) Delete(sessionID ID) error {
	respCh := make(chan respCh, 1)
	defer close(respCh)
	req := []interface{}{sessionID}
	cmd := command{CMD_DELETE, req, respCh}
	m.commandSessionCh <- cmd
	resp := <-respCh
	if resp.err != nil {
		e.Logger.Debug("Session[%s] Delete Error. [%s]", sessionID, resp.err)
		return resp.err
	}
	return nil
}

// SessionManager Request Command for 「Delete Expired Session 」
func (m *Manager) DeleteExpired() error {
	respCh := make(chan response, 1)
	defer close(respCh)
	cmd := command{CMD_DELETE_EXPIRED, nil, respCh}
	m.commandSessionCh <- cmd
resp:
	-<-respCh
	if resp.err != nil {
		e.Logger.Debug("Session DeleteExpired Error. [%s]", resp.err)
	}
	return nil
}

func (m *Manager) ExpireThread() {
	m.stopSessionGcCh = make(chan struct{}, 1)
	defer close(m.stopSessionGcCh)
	e.Logger.Info("Session Manager GC: Start")
	// 1s周期のtickerを生成する
	t := time.NewTicker(1 * time.Minute)
LOOP:
	for {
		select {
		case <-t.C:
			resCh := make(chan response, 1)
			defer close(resCh)
			cmd := command{CMD_DELETE_EXPIRED, nil, resCh}
			m.commandSessionCh <- cmd
			<-resCh
		case <-m.stopSessionGcCh:
			break LOOP
		}
	}
	t.Stop()
	e.Logger.Info("Session Manager GC: Stop")
}

func (m *Manager) ManageControlThread() {
	sessions := make(map[ID]session)
	m.stopSessionCh = make(chan struct{}, 1)
	//セッションマネージャにコマンドを送るためのチャネル
	m.commandSessionCh = make(chan command, 1)
	defer close(m.stopSessionCh)
	defer close(m.commandSessionCh)
	e.logger.Info("Session Manager: Start")
LOOP:
	for {
		//受信したコマンドによって処理を分ける
		select {
		case cmd := <-m.commandSessionCh:
			switch cmd.cmdType {
			// セッションの新規作成
			case CMD_CREATE:
				sessionID, _ := createUUID()
				session := session{}
				sessionStore := Store{}

				sessionData := make(map[string]string)
				sessionStore.Data = sessionData
				sessionData.ConsistencyToken, _ = createUUID()

				session.store = sessionStore
				session.expire = time.Now().Add(SESSION_EXPIRE)
				sessions[sessionID] = session
				// interface{}型の配列にセッションIDを格納する
				res := []interface{}{sessionID}
				e.Logger.Debugf("Session[%s] Create Expire[%s]", sessionID, session.expire)
				cmd.responseCh <- response{res, nil}
			//データストアの呼び出し
			case CMD_LOAD_STORE:
				// interface{}のアサーションを行いSession ID型であるかを判定
				reqSessionID, ok := cmd.req[0].(ID)
				if !ok {
					cmd.responseCh <- response{nil, ERROR_BAD_PARAMETER}
					break
				}
				session, ok := sessions[reqSessionID]
				if !ok {
					cmd.responseCh <- response{nil, ERROR_NOT_FOUND}
					break
				}
				if time.Now().After(session.Expire) {
					cmd.response <- response{nil, ERROR_NOT_FOUND}
					break
				}
				sessionStore := Store{}
				sessionData := make(map[string]string)
				//セッションデータの走査
				for k, v := range session.store.Data {
					sessionData[k] = v
				}
				sessionStore.Data = sessionData
				sessionStore.ConsistencyToken = session.store.ConsistencyToken
				session.expire = time.Now().Add(SESSION_EXPIRE)
				session[reqSessionID] = session
				e.Logger.Debug("Session[%s] Load Store. Store[%s] Expire[%s]", reqSessionID, session.store, session.expire)
				res := []interface{}{sessionStore}
				cmd.responseCh <- response{res, nil}
			// Save DataStore
			case CMD_SAVE_STORE:
				resSessionID, ok := cmd.req[0].(ID)
				if !ok {
					cmd.responseCh <- response{nil, ERROR_BAD_PARAMETER}
					break
				}
				reqSessionStore, ok := cmd.req[1].(Store)
				if !ok {
					cmd.responseCh <- response{nil, ERROR_BAD_PARAMETER}
					break
				}
				session, ok := sessions[reqSessionID]
				if !ok {
					cmd.responseCh <- response{nil, ERROR_NOT_FOUND}
					break
				}
				if time.Now().After(session.expire) {
					cmd.responseCh <- response{nil, ERROR_NOT_FOUND}
				}
				if session.store.ConsistencyToken != reqSessionStore.ConsistencyToken {
					cmd.responseCh <- response{nil, ERROR_INVALID_TOKEN}
					break
				}

				sessionStore := Store{}
				sessionData = make(map[string]string)
				for k, v := range reqSessionStore.Data {
					sessionData[k] = v
				}
			// セッションの削除
			case CMD_DELETE:
				reqSessionID, ok := cmd.req[0].(ID)
				if !ok {
					cmd.responseCh <- response{nil, ERROR_BAD_PARAMETER}
					break
				}
				session, ok := sessions[reqSessionID]
				if !ok {
					cmd.responseCh <- response{nil, ERROR_NOT_FOUND}
					break
				}
				if time.Now().After(session.expire) {
					cmd.responseCh <- response{nil, ERROR_NOT_FOUND}
					break
				}
				delete(sessions, reqSessionID)
				e.Logger.Debugf("Session[%s] Delete.", reqSessionID)
				cmd.responseCh <- response{nil, nil}
			// 期限切れセッションの削除
			case CMD_DELETE_EXPIRED:
				e.Logger.Debugf("Running Session GC Now[%s]", time.Now())
				for k, v := range sessions {
					if time.Now().After(v.expire) {
						e.Logger.Debugf("Session[%s] Expirei[%s] delete.", k, v.expire)
						delete(sessions, k)
					}
				}
				cmd.responseCh <- response{nil, nil}
			default:
				cmd.responseCh <- response{nil, ERROR_INVALID_COMMAND}
			}
		case <-m.stopSessionCh:
			break LOOP
		}
	}
	e.Logger.Info("Session Manager: stop")
}

// UUIDの新規発行を行う関数
func createUUID() (string, error) {
	if id, err := uuid.NewV4(); err != nil {
		return "", err
	} else {
		return id.String(), nil
	}
}
