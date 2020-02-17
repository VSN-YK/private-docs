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

### GETパラメーダを受け取るコード

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
  ```
  ```
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
