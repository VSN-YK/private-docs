## golangでレイヤードアーキテクチャについて勉強してみた

### 各レイヤのモデルについて(golangに依存した話ではない)

[golangではないですけどモデルの概念は変わらない](https://terasolunaorg.github.io/guideline/public_review/Overview/ApplicationLayering.html#id4)

- **アプリケーション層**
```
情報の入出力となるUIを提供したり、リクエスト情報をドメイン層や、他システムから呼び出し、表示用の出力を返す手続きを行うなど、
アプリケーションを構築するための層である。この層は、できるだけ薄く保たれるべきであり、ビジネスルールを含んではいけない。
```
- **ドメイン層**
```
ドメイン層は、アプリケーションのコアとなる層である。ビジネス上の解決すべき問題を表現し、
ビジネスオブジェクトや、ビジネスルールを含む(口座へ入金する場合に、残高が十分であるかどうかのチェックなど)。
ドメイン層は、他の層からは疎であり、再利用できる。
```
- **インフラストラクチャー層**
```
インフラストラクチャ層では、ドメイン層(Repositoryインタフェース)の実装を提供する。
データストア(RDBMSや、NoSQLなどのデータを格納する場所)への永続化や、メッセージの送信などを担う。
```

[Goでのモデリング](https://www.slideshare.net/pospome/go-80591000)


### レイヤ構造にするメリット

- コードの依存関係を整理できる
- レイヤ間の差し替えが容易になる
- レイヤ内のパッケージの凝集度を高められる

### 実装を行う中での注意点
- レイヤの責務を理解し明確にする　-> ここがシステムの改修基点となるから
  今まではHTTPの`POST`でリクエストしていたものを`PUT`に変わる(hundler起点)

  ```go
  type User struct {
    Name string
    Address string
  }

  ```


  上記の制約を元にして先人の知恵を借り、レイヤードアーキテクチャのモデルを作ってみる

### User-Infoを返すAPIを作成してみる
1.Domain層にモデルを作成   `domain/model/UserInfo.go`

**Domain層はシステムが担うビジネスロジックのモデルを定義**するところである。
今回の場合、ユーザの情報すなわちUserInfoを返したいのでその元となるモデルを作成する。

  ```go
  package model

  import (
    "time"
  )
  type UserInfo (){
    Id    string,
    Name  string,
    Age   int,
    Join  time.time,
  }
  ```


  [参考文献](https://yyh-gl.github.io/tech-blog/blog/go_web_api/)
