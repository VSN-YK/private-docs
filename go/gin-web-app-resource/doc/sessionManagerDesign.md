## SessionManagerを作ってみる

[参考文献](https://leben.mobi/go/session-design/practice/web/)

構成について

```sh
.
├── main.go
└── sessions
    ├── manager.go
    └── session.go
```
では早速 <span style="color: red; ">manager.go</span>から実装していきます。
今回使用するパッケージについては以下とします。

<table>
	<tr>
		<th>パッケージ名</th>
		<th>機能概要</th>
	</tr>
	<tr>
		<td>crypto/rand</td>
		<td>疑似乱数生成を行う</td>
	</tr>
	<tr>
		<td>encoding/base64</td>
		<td>base64でエンコーディングを行う</td>
	</tr>
	<tr>
		<td>errors</td>
		<td>エラー管理を行う</td>
	</tr>
	<tr>
		<td>io</td>
		<td>入出力処理を行う</td>
	</tr>
	<tr>
		<td>net/http</td>
		<td>HTTPを扱うパッケージで、HTTPクライアントとHTTPサーバーを実装するために必要な機能が提供されている。</td>
	</tr>
</tabel>

- [ ] セッションマネージャ構造体の定義を行う

	Map型で定義を行っている理由についてはまずキーとして一意となるセッションIDを持たせ、
	そのキーに紐づくセッション情報を格納したいため

	```go
	type Manager struct {
		database map[string]interface{}
	}
	```
- [ ] `Manager構造体`の初期化を行うためのコンストラクタを作成する。

	コンストラクタであるため、このファイルが読み込まれるタイミングで初期化が行われる。

	```go
	var manager Manager

	func NewManager() *Manager {
		return &manager
	}
	```
- [ ] 新規セッションIDの払い出しを行う関数を作成する

	まずこの関数みて着目するべきところは`rand.Readerという`ioパッケージのインスタンス変数を利用していることです。
	なぜ`crypto/randパッケージ`がReaderインスタンス変数を利用できるかについては[Goの公式パッケージDOC](https://golang.org/pkg/crypto/rand/#pkg-variables)をみることで理解できます。あとは64バイトでゼロクリアされた配列にcyrpto/randで擬似乱数を生成してbase64で文字列にエンコードすることでセッションIDを払い出すことができます。

	ドキュメントの一部分を抜粋します。
	```go
	var Reader io.Reader

	Reader is a global, shared instance of a cryptographically secure random number generator.
	```
	上記からReaderは共有インスタンス変数としてcryptographicallyで利用できると謳っておりました。これにより`cyrpto/rand.Reader`の解釈ができると思います。

	```go
	func (m * Manager) PayLoadNewSessionId() string {
		b := make([]byte, 64)
		if _, err := io.ReadFull(rand.Reader, b); err := nil {
			return ""
		}
		return base64.URLEncoding.EncodeToString(b)
	}
	```

- [ ] 新規セッションの生成を行う

	まず、Newメソッドの第一引数である`*http.Request`についてみていきます。この形の解釈としは
	`net/httpパッケージの中で定義されているRequest構造体`を示していることになります。

	実際に[Request構造体がどのような形状になっているかは下記のコード](https://github.com/golang/go/blob/master/src/net/http/request.go#L108)を参照ください。

	次に`r.Cookie()メソッド`についてみていきます。まずこのメソッドのRequest構造体に紐づくものであることがわかると思いますので[メソッドの定義](	https://github.com/golang/go/blob/master/src/net/http/request.go#L422)をみます。

	下記はメソッドを抜粋したものです。
	ソースをみるとfor文でreadCookies関数から取得したCookieのリストを回していることがわかります。せっかくですのでもう少し詳しくみてみます。

	```go
	func (r *Request) Cookie(name string) (*Cookie, error) {
		for _, c := range readCookies(r.Header, name) {
			return c, nil
		}
		return nil, ErrNoCookie
		}
	```
	先ほどの`readCookies関数`の話に戻します。この関数の第一引数として渡している`r.Header`という値に注目します。これはリクエストヘッダと呼ばれるHTTP通信の際に付与されるヘッダーであり
	通信の大まかな要求が記述されているものとして捉えてください。

	では関数内部の話に戻します。
	関数の237行目に着目してください。
	`lines := h["Cookie"]`というコードでHeaderマップから"Cookieをキーとして`リクエストライン`が存在するか判定を行っております。

	**※リクエストラインとは...**

	[HTTPリクエストの中で記述されているリクエストヘッダーの一行目の部分を示し、HTTPリクエスト全体の概要が大まかに記述されております。](https://wa3.i-3-i.info/word1843.html)

	```go
	func readCookies(h Header, filter string) []*Cookie {
	lines := h["Cookie"]
	if len(lines) == 0 {
		return []*Cookie{}
	}

	cookies := make([]*Cookie, 0, len(lines)+strings.Count(lines[0], ";"))
	for _, line := range lines {
		line = strings.TrimSpace(line)

		var part string
		for len(line) > 0 { // continue since we have rest
			if splitIndex := strings.Index(line, ";"); splitIndex > 0 {
				part, line = line[:splitIndex], line[splitIndex+1:]
			} else {
				part, line = line, ""
			}
			part = strings.TrimSpace(part)
			if len(part) == 0 {
				continue
			}
			name, val := part, ""
			if j := strings.Index(part, "="); j >= 0 {
				name, val = name[:j], name[j+1:]
			}
			if !isCookieNameValid(name) {
				continue
			}
			if filter != "" && filter != name {
				continue
			}
			val, ok := parseCookieValue(val, true)
			if !ok {
				continue
			}
			cookies = append(cookies, &Cookie{Name: name, Value: val})
		}
	}
	return cookies
}
```

	```go
	func (m *Manager) New(r.*http.Request, cookieName string) (*Session , error) {
		cookie, err := r.Cookie(cookieName)
		if err != nil && m.Exists(cookie.Value) {
			return nil, errors.New("Session Id Was AllReady Issued")
		}

		session := NewSession(m, cookieName)
		session.Id = m.PayLoadNewSessionId()
		session.Request = r

		return session, nil
	}
	```
