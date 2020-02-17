## Node.js 入門

### 参考文献
- [Qiita MacにNode.jsをインストール](https://qiita.com/kyosuke5_20/items/c5f68fc9d89b84c0df09)

### Node.jsの特徴

- サーバーサイドの JavaScript 実行環境
- Google V8 JavaScript エンジンを使用しており、高速。
- npm (Node Package Manager) と呼ばれるパッケージ管理システムを同梱。
- [ノンブロッキングI/O](https://nodejs.org/ja/docs/guides/blocking-vs-non-blocking/) と イベントループ アーキテクチャにより、**10K問題** (クライアント1万台レベルになると性能が極端に悪化する問題) に対応。
- 通信やファイルの読み書きをノンブロッキングI/Oで処理するため、スレッドが I/O 待ちになる頻度が少なく、効率的。

### Node.jsのインストール

#### Mac編
- HomeBrewのインストール(割愛)
- nodebrewのインストール

  ```sh
  # nodebrewのsrcを管理するためのdirを作成
  $ mkdir -p .nodebrew/src
  $ brew install nodebrew
  $ nodebrew -v && echo "Installed"
  ```

- Node.jsのインストール

  ```sh
  # どんなバージョンがあるのか気になったのでファイルに書き出してみた
  $ node_ver_list=($(nodebrew ls-remote)) &&  for v in "${node_ver_list[@]}"; do echo $v ;done > version.txt
  # 安定版のバイナリを取得しコマンドの結果が正常である場合はインストールしたバージョンを表示する
  $ nodebrew install-binary stable && nodebrew ls
    v7.1.0
    current: none # インストールしたバージョンを有効化する必要がある。
  $ nodebrew use $(nodebrew ls)
    v12.16.0
    current: v12.16.0
  # バイナリのPATHを通す(.zshrcに追記する例)
  $ echo 'export PATH=$HOME/.nodebrew/current/bin:$PATH' >> ~/.zshrc
  # ターミナルの再起動
  $ exec $SHELL -l
  # nodeのPATHが通っているか確認を行う
  $ node -v
  ```

#### Windows編

### Node.jsを使用してみる
```sh
# nodejsのsrcを管理するためのdirを作成
$ mkdir ~/NodeJs && cd NodeJs
$ vim node_entry.js
  console.log("Hello Node.js")
$ node node_entry.js
```
