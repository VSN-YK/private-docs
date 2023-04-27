# dockerについての入門書
[参考文献](https://knowledge.sakura.ad.jp/13265/)

## 1. dockerとは?
```
コンテナ仮想化を用いてアプリケーションを開発・配置・実行するためのオープンソースソフトウェアあるいはオープンプラットフォームである。
```
※ ここでポイントなるワードは**コンテナ仮想化**になります。

```
コンテナ仮想化とは
ホストマシンのカーネル(OSの中核)を利用し、プロセスやユーザなどを隔離することで、あたかも別のマシンが動いているかのように動かす仕組みのこと
```
![](https://cn.teldevice.co.jp/asset/migration-files/kcf/image/column/compass/20160202/docker01.jpg)

## 2. dockerのリソース管理と利点

### リソース管理について
dockerはミドルウェアのインストールや各種環境設定をコード化して管理する。
このモデルは`Infrastracture as Code(IaC)`という言葉で定義されている。 // comment 一般的には、IaCはミドルウェアだけじゃなくて、インフラ全般ですね。オンプレならOSやネットワーク、クラウドサービスなら例えばAWSならAWSリソースもそうですね。これらはミドルウェア（nginxとかTomcatとかmysqlとか）や各種環境設定（それらの設定ファイル）には含まれないので。

### IaCによるメリット
- [ ] コード化されたファイルを共有することで、どこでも誰でも同じ環境が作れる。
  ```
  dockerのインストールが済んでしまえば環境に依存することなくdockerのイメージをチームで共有して使用することができる。 // comment 上記コメントと同じ指摘ですが、「dockerのインストールも IaCでやることの一部だよ」ということですね。まぁプロジェクトによっては「OSまでは手動でインストールしよう、ここまでできたらあとはコードでやろう。それがうちのIaCのやり方だ」って決めるのもありですけどね。コスパとか制約とかあるしね。
  ```
- [ ] 作成した環境を配布しやすい。
  ```
  docker registoryにpushしておくことでdocker pullコマンドによるイメージの利用が可能
  ```
- [ ] スクラップ＆ビルドが容易にできる。

## 3.dockerのインストール
dockerのインストールについてはOS毎にバイナリが異なるのでそれぞれの端末に合ったものをインストールする。
- [docker for Mac](https://hub.docker.com/editions/community/docker-ce-desktop-mac)
- [docker for windows](https://hub.docker.com/editions/community/docker-ce-desktop-windows)

今回は`Mac`環境での構築を例にします。

dockerhubのアカウントを作成し、赤枠で囲っている`Get Docker`ボタンを押下する
ファイルを解凍するとインストーラーが起動する。
![image](https://user-images.githubusercontent.com/60165356/81518134-01932300-9378-11ea-9ad9-b467f2a65a0e.png)

`command` + `space` でアプリの検索ができるのでここで`docker`と入力しdocker.appが起動していることを確認する。
![image](https://user-images.githubusercontent.com/60165356/81519786-571dfe80-937d-11ea-947a-004160d9399d.png)

ターミナルから実際に`docker demon`が起動していることをdocker help コマンドで確かめる
```sh
$ docker --help
```
dockerhubから公式のイメージを取得する。
```sh
docker pull nginx
```
今回はnginxのイメージを取得し実際にWebサーバを起動します。
```sh
 docker run --name nginx-container -d -p 8080:81 nginx
```
コンテナが無事に生成されているか確認する。
```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                  NAMES
f20cb911dfff        nginx               "nginx -g 'daemon of…"   29 seconds ago      Up 28 seconds       0.0.0.0:8080->80/tcp   nginx-container
```
実際に`http://localhost:8080/`でサーバにリクエストを送ってみる。
![image](https://user-images.githubusercontent.com/60165356/81611656-643afc00-9416-11ea-9466-9600e594bfe8.png)

無事にリクエストを受け取ることができた。
このようにdockerによるコンテナ開発を行うことで楽にwebサーバなどを立ち上げることができます。また、今回は[dockerHub](https://hub.docker.com/_/nginx)にあるイメージを使用してコンテナを作成しましたが、もちろんDockerfileによるイメージのカスタマイズも可能です。
