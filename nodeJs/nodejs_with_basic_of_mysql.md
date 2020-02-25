## Node.jsでMySQLに接続

- MySQLのセットアップを行う

  ```sh
  $ brew install mysql
  # MySQLの起動を行う
  $ mysql server.start
  # Rootユーザでログインを行う
  $ mysql -uroot
  # スキーマの作成を行う
  mysql> create sub;
  # ユーザの作成を行う
  mysql> create user 'sub_user'@localhost identified by 'sub';
  # ユーザの確認を行う
  mysql> mysql> select user,host from mysql.user where user = 'sub_user';
  +----------+-----------+
  | user     | host      |
  +----------+-----------+
  | sub_user | localhost |
  +----------+-----------+
  1 row in set (0.01 sec)
  # GRANT OPTION と PROXY 以外の全ての権限を設定
  mysql > GRANT ALL PRIVILEGES ON sub.* TO sub_user@localhost;
  mysql > exit;
  # sub_userでログインを行うためパスワードを入力する
  $ mysql -u sub_user -p
  # スキーマの変更を行う
  mysql> use sub;
  # テーブルの作成
  mysql> create table sub.LANGUAGE_TBL(id varchar(4) not null primary key, name varchar(30));
  # データの挿入
  mysql> insert into node_database.user value('L001','Node.jS');
  # テーブルにデータが挿入されたことを確認する
  mysql> select * from sub.LANGUAGE_TBL;
  +------+---------+
  | id   | name    |
  +------+---------+
  | L001 | Node.js |
  +------+---------+
  1 row in set (0.00 sec)
  ```
- Node.jsからMySQLに接続を行う
  ```js
  var mysql = require("mysql");
  var connection = mysql.createConnection({
      host : 'localhost',
      user : 'sub_user',
      password : 'sub'
  });
  connection.connect(function(err){
      if (err) {
          console.log('[Connecting Error]' + err.stack);
          connection.end();
          return;
      }
      console.log('Connecting was Succeed!');
      connection.end();
  });

  ```
  ここで以下のエラーと遭遇

  ```
  [Connecting Error]Error: ER_NOT_SUPPORTED_AUTH_MODE: Client does not support authentication protocol requested by server; consider upgrading MySQL client
  ```
  <details>
  <summary>パスワードの認証方式がMySQL側で変更されていたが、Node側のパッケージでは対応されていなかったために発生したエラー</summary>

  ```js
  [Connecting Error]Error: ER_NOT_SUPPORTED_AUTH_MODE: Client does not support authentication protocol requested by server; consider upgrading MySQL client
      at Handshake.Sequence._packetToError (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/sequences/Sequence.js:47:14)
      at Handshake.ErrorPacket (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/sequences/Handshake.js:123:18)
      at Protocol._parsePacket (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/Protocol.js:291:23)
      at Parser._parsePacket (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/Parser.js:433:10)
      at Parser.write (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/Parser.js:43:10)
      at Protocol.write (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/Protocol.js:38:16)
      at Socket.<anonymous> (/Users/hacknatural/nodeJS/node_modules/mysql/lib/Connection.js:88:28)
      at Socket.<anonymous> (/Users/hacknatural/nodeJS/node_modules/mysql/lib/Connection.js:526:10)
      at Socket.emit (events.js:321:20)
      at addChunk (_stream_readable.js:294:12)
      --------------------
      at Protocol._enqueue (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/Protocol.js:144:48)
      at Protocol.handshake (/Users/hacknatural/nodeJS/node_modules/mysql/lib/protocol/Protocol.js:51:23)
      at Connection.connect (/Users/hacknatural/nodeJS/node_modules/mysql/lib/Connection.js:116:18)
      at Object.<anonymous> (/Users/hacknatural/nodeJS/connect-in-mysql.js:8:12)
      at Module._compile (internal/modules/cjs/loader.js:1157:30)
      at Object.Module._extensions..js (internal/modules/cjs/loader.js:1177:10)
      at Module.load (internal/modules/cjs/loader.js:1001:32)
      at Function.Module._load (internal/modules/cjs/loader.js:900:14)
      at Function.executeUserEntryPoint [as runMain] (internal/modules/run_main.js:74:12)
      at internal/main/run_main_module.js:18:47
  ```
  </details>

  ```sql
  # MySQL側のパスワード認証方式をNativeに変更する
  mysql> ALTER USER 'sub_user'@'localhost' IDENTIFIED WITH mysql_native_password BY 'sub';
  Query OK, 0 rows affected (0.01 sec)```
  ```

  再度スクリプトを実施する

  ```sh
  $ node first-connect-in-mysql.js
  Connecting was Succeed!

- Expressからの接続も確かめる

  ```js
  var express = require('express');
  var router = express.Router();
  var mysql = require('mysql');

  var connection = mysql.createConnection({
      host : 'localhost',
      user : 'sub_user',
      password : 'sub',
      database : 'sub'
  });

  router.get('/', function(req, res, next) {
    connection.query('select * from sub.LANGUAGE_TBL',function(error , results , fields){
        if (error) {
            console.log("Not Connected To DataBase");
            connection.end();
            return;
        }
        res.send(results);
        console.log(results);
        connection.end();     
    });
  });

  module.exports = router;
  ```
  webサーバの起動を行いレコードの確認を行う
  ```sh
  $ npm start

  > express-web-app@0.0.0 start /Users/hacknatural/nodeJS/express-web-app
  > node ./bin/www

  [
    RowDataPacket { id: 'L001', name: 'Node.js' },
    RowDataPacket { id: 'L002', name: 'Java' }
  ]
  ```
