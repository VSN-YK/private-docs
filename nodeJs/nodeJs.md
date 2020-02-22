## Node.js 実践

### アプリケーションフォルダにnpmを使用しPackageをインストールする

- npmの初期化を行う(package.jsonの作成)

    ```sh
    # githubのnodejsディレクト直下でnpmの初期化を行う
    $ cd ~/Documents/private-docs/nodeJs && npm init

    This utility will walk you through creating a package.json file.
    It only covers the most common items, and tries to guess sensible defaults.

    See `npm help json` for definitive documentation on these fields
    and exactly what they do.

    Use `npm install <pkg>` afterwards to install a package and
    save it as a dependency in the package.json file.

    Press ^C at any time to quit.
    package name: (nodejs) node-app-package
    version: (1.0.0)
    description:
    entry point: (index.js)
    test command:
    git repository:
    keywords:
    author: Y.K
    license: (ISC)
    About to write to /Users/developer01/Documents/private-docs/nodeJs/package.json:

    {
      "name": "node-app-package",
      "version": "1.0.0",
      "description": "",
      "main": "index.js",
      "scripts": {
        "test": "echo \"Error: no test specified\" && exit 1"
      },
      "author": "Y.K",
      "license": "ISC"
    }

    Is this OK? (yes)

    ls -la | grep *.json
    -rw-r--r--   1 developer01  staff   215  2 17 11:31 package.json

   ```

- jqueryのpackageをnpmでinstall

  ```sh
  # --saveオプションを付与することでpackage.jsonに依存関係を持たせる
  $ npm install jquery --save
  npm notice created a lockfile as package-lock.json. You should commit this file.
  npm WARN node-app-package@1.0.0 No description
  npm WARN node-app-package@1.0.0 No repository field.

  + jquery@3.4.1
  added 1 package from 1 contributor and audited 1 package in 0.384s
  found 0 vulnerabilities
  # jqueryのpackageがインストールできたかを確認する
  $ npm list
  node-app-package@1.0.0 /Users/developer01/Documents/private-docs/nodeJs
  └── jquery@3.4.1
  ```

### Express パッケージを使用しWeb Appを作成　

<details>
<summary>Express パッケージとは?</summary>

  ```
  Expressは、nodeを使ったWebアプリケーションを作成するためのワークフレーム
  ```
  </details>

#### 参考文献

- [MacでExpressを使用する](https://rikson.net/install-express/)

  ```sh
  # express packageをグローバルとして適応する
  $ npm install -g express
  ```

  **※ express パッケージはexpress プロジェクトをgenerateする機能はない**

  WEBサーバーを立ち上げてみる

  express.js

  ```js
  // express モジュールをimport
  var express = require('express');
  var app = express();
  //GETメソッドでコンテキストPATHに/が渡ったきた場合にハンドルされfunction()呼び出される
  app.get('/' , function(req , res){
      res.send('Hello Node.js From Express');
  });

  app.listen(8082 , function() {
      console.log('Running Server Port On 8082')
  });
  ```
  実行した際に以下のエラーが表示される場合は `npm root -g`でモジュールの場所を出力し
  環境変数として設定するとfixする場合がある。

  ```sh
  $ echo "NODE_PATH=$(npm root -g)" >> ~/.zshrc
  ```

  サーバの始動

  ```sh
  $ node express.js
  ```
  別ターミナルでcurlコマンドを実行しWebサーバが始動状態にあることを確認する

  ```sh
  $ curl http://localhost:8082/
  Hello Node.js From Express
  ```

### Expressのプロジェクトを利用しテンプレートWebアプリの動作を検証する

- Expressの雛形プロジェクトの生成を行ってくれるCLI(`express-generator`)をインストールし、
  express コマンドでプロジェクトを作成する

  ```sh
  $ npm install -g express-generator
  $ express express-web-app
  ```
  <table>
    <tr>
      <th>Dir Name</th>
      <th>Roles</th>
    </tr>
    <tr>
      <td>public</td>
      <td>css,imgなどの静的なファイルを配置するディレクトリ</td>
    </tr>

    <tr>
      <td>routes</td>
      <td>コンテキストのPATHを元にRouterとして機能するディレクトリ</td>
    </tr>

    <tr>
      <td>views</td>
      <td>Viewを担うディレクトリ</td>
    </tr>


  </table>

  プロジェクトの構成についてtreeコマンドで確認する。

  ```sh
  $ cd express-web-app
  $ tree
  .
  ├── app.js
  ├── bin
  │   └── www
  ├── package.json
  ├── public
  │   ├── images
  │   ├── javascripts
  │   └── stylesheets
  │       └── style.css
  ├── routes
  │   ├── index.js
  │   └── users.js
  └── views
      ├── error.jade
      ├── index.jade
      └── layout.jade
  ```
  package-lock.jsonにexpress-web-app内で使用していパッケージとの依存関係を記述

  ```sh
  $ npm install -d
  ```
  ちなみにpackage-lock.jsonには以下のように各パッケージとの依存関係が既述されている
  ```sh
  # acorn パッケージとの依存関係
  $ cat  package-lock.json | jq .dependencies.acorn
  {
    "version": "2.7.0",
    "resolved": "https://registry.npmjs.org/acorn/-/acorn-2.7.0.tgz",
    "integrity": "sha1-q259nYhqrKiwhbwzEreaGYQz8Oc="
  }
  ```
  appの起動を行う

  ```sh
  $ npm start

  > express-web-app@0.0.0 start /Users/hacknatural/nodeJS/express-web-app
  > node ./bin/www  
  ```
  別ターミナルからcurlコマンドを実行し、WebServerとの接続を確かめる

  ```sh
  $ curl -sSf http://localhost:3000/
  <!DOCTYPE html><html><head><title>Express</title><link rel="stylesheet" href="/stylesheets/style.css"></head><body><h1>Express</h1><p>Welcome to Express</p></body></html>
  ```

### Expressで簡易的なWebAppを作成する

今回作成を行うWebアプリはテキストフィールドに入力されたValueをPOSTで送信し
その内容をConfirmとして表示する簡易的なアプリである。

**アプリの構成**

```sh
.
├── app.js
├── bin
│   └── www
├── package.json
├── public
│   ├── images
│   ├── javascripts
│   └── stylesheets
│       └── style.css
├── routes
│   ├── index.js
│   ├── users.js
│   └── wellcome.js
└── views
    ├── confirm.jade
    ├── error.jade
    ├── index.jade
    ├── layout.jade
    └── wellcome.jade

7 directories, 12 files
```
**test.appのルーティングの設定を行うめアプリケーションのインスタンスを作成する**

```js
var app = require('./routes/test');
console.log(app)
```
appのインスタンス情報を出力したもの
```
[Function: router] {
  params: {},
  _params: [],
  caseSensitive: undefined,
  mergeParams: undefined,
  strict: undefined,
  stack: [
    Layer {
      handle: [Function: bound dispatch],
      name: 'bound dispatch',
      params: undefined,
      path: undefined,
      keys: [],
      regexp: /^\/?$/i,
      route: [Route]
    }
  ]
}
```

### GETパラメーダを受け取るコード(Expressを使用しない例)

- Server側のソース

  ```js
  var http = require("http");
  var url = require("url");
  // Serverのインスタンスを生成する
  var server = http.createServer(function(req , res){
  	console.log("[Debug]" + req);
  	var url_parse = url.parse(req.url , true);
  	console.log(url_parse);
  	res.end();
  }).listen(8083);
  ```

- 別ターミナルからリクエスト送信し、実際にパラメータを取得する。

  ```sh
  $ echo $0
    -zsh

  $ curl -X GET http://localhost:8083/node-test\?name=nodeJS

  Url {
  protocol: null,
  slashes: null,
  auth: null,
  host: null,
  port: null,
  hostname: null,
  hash: null,
  search: '?name=nodeJS',
  query: [Object: null prototype] { name: 'nodeJS' },
  pathname: '/node-test',
  path: '/node-test?name=nodeJS',
  href: '/node-test?name=nodeJS'
  }
  ```
- [参考文献](https://qiita.com/Coolucky/items/4b35622d32d6eae2ebc4)
