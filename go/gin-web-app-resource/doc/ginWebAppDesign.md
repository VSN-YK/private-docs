## ginパッケージを用いWeb Appを作成

### 今回作成するアプリケーションの概要

```
localhost:8089/にアクセスを行うとログイン画面(login.html)に遷移を行う。
この場合は初回アクセスとなるため、ユーザに「ユーザID」「ユーザ名」「パスワード」「メール」
の登録を要求する。登録が完了すると登録完了画面(wellcome_member.html)に遷移し、ユーザに
アカウントの登録が完了したことを促す。

次回以降のアクセスについてはcookieにてセッション管理を行う。

注釈 セッション管理の詳細はアカウント管理ロジック※1を参照する。
```

#### アプリケーションの概要
![image](https://user-images.githubusercontent.com/60165356/79192989-b4995b00-7e64-11ea-915e-66d44c1014c3.png)

#### アカウント管理ロジック
**アカウント管理ロジックについては以下の機能を有する**
  - [ ]  ユーザのアカウントを作成する機能
  - [ ]  ユーザのアカウントを破棄する機能

**※1またこの時アカウントの管理にはCookieを利用してセッション管理を行う**

- [ ] **構成要素については以下を参照**

  <table>
    <tr>
      <th colspan="2"><center>構成要素</center></th>
    </tr>
    <tr>
      <th>パッケージ</th>
      <td>github.com/gin-contrib/sessions<br>github.com/gin-contrib/sessions/cookie</td>
    </tr>
    <tr>
      <th>セッション管理</th>  
      <td>cookie</td>
    </tr>
  </table>

- [ ] **セッション管理のシーケンスフローを下記に示す。**

  `Ex:DBにセッションIDを登録するフロー`

  ![](../image/sessionSequenceFlow.png)

- [ ] **Cookieを利用している実際のソース**

  セッションストアの初期化を行う

  ```go
  store := cookie.NewStore([]byte("secret"))
  //mysessionと呼ばれるセッションインスタンスが生成される。
  router.Use(sessions.Sessions("mysession", store))
  ```
  ```go
  //  クライアントから送信されるコンテキストを元にセッション情報の取得を行う。
  session := sessions.Default(ctx)
  ```
  <details>

  <summary>折角なのでもうすこし細かいレベルでソースを追ってみました。</summary>

  `.(Session)`で型変換を行っていることから`*gin.Context.MustGet(DefaultKey)`
  が`インタフェース型`を返していることがわかる。
  また、`*Context.MustGet()`メソッドのキー値として渡しているパラメータ([github.com/gin-contrib/sessions](https://github.com/gin-contrib/sessions/blob/0b3a2450f71448ee0facc69311f62926928b3d16/sessions.go#L13))は実際では[*Context.Get()メソッド](https://github.com/gin-gonic/gin/blob/4f208887e1231459672a2a9fc1b2aa40486825d4/context.go#L246)内部で `github.com/gin-contrib/sessions`キーに紐づくValueが存在するかどうかを判定してその結果をインタフェースとして返す仕様になっていた。

  ```go
  // shortcut to get session
  func Default(c *gin.Context) Session {
     return c.MustGet(DefaultKey).(Session)
  }
  ```
  c.MustGet(DefaultKey)から取得した実際のセッション情報
  ```sh
  # 実際にクライアントから取得したセッション情報
  {mysession 0xc0003d4300 0xc000350020 <nil> false 0xc0003de0e0}
  ```

  余談) 公式Docを確認するとやはりそうであった。

  ```go
  func (c *Context) MustGet(key string) interface{}
  ```

  [structの形状などについてはsessionパッケージを参照](https://github.com/gin-contrib/sessions/blob/master/sessions.go#L138)

  </details>

  ```go
  //  取得したセッション情報にパラメータのセットを行う
  if session.Get("LoginId") != "" {
    session.Set("LoginId", loginId)
    session.Save()
  }
  ```

  ```go
  //セッションの破棄
  session.Clear()
  session.Save()
  ```
