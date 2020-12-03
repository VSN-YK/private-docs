## PlantUMLをWindows with ATOMで使ってみる

  <table>
    <tr>
      <th colspan="2">PlantUMLの動作に必要なソフトウェア</th>
    </tr>
    <tr>
      <td>JDK8(1.8)</td>
      <td>javaのプログラムを開発したりするときに使う<br>
      とはいいつつも主役はJREです</td>
    </tr>
    <tr>
      <td>Graphiz dot</td>
      <td>PlantUMLでチャートを実現する時の基盤を<br>
      提供してくれます</td>
    </tr>
    <tr>
      <td>Plantuml.jar</td>
      <td>PlantUMLの基幹プルグラム(.javaの集合体)を<br>
      圧縮したもので、主にplantumlのソースコードをpng画像に変換する際に使用します。</td>
    </tr>
  </table>

- [ ] [1.Graphiz.Dotのインストールを行う](https://www2.graphviz.org/Packages/development/windows/10/msbuild/Release/Win32/)

- [ ] 2.CMDでお手元のjavaのバージョンを確認し、1.8.xであることを確かめる
[インストールしていない人は本家であるOracleのHPからダウンロードをおこなってください。](https://www.oracle.com/jp/java/technologies/javase/javase-jdk8-downloads.html)

  ```shell
  java -version
  ```
- [ ] 3.Plantuml.jarを以下のサイトから入手する。

  https://plantuml.com/ja/download

- [ ] Atomエディタをインストールする

  https://atom.io/

- [ ] Atomエディタ内で必要なパッケージをインストール

  - PlantUML-preview
  - language-plantuml

- [ ] インストールを行った`plantuml-preview`パッケージのコンフィグを変更する
