# dockerの知識

## Dockerfileの書き方について

- [ ] Dockerfileとは

  ```
  Dockerfile
  dockerイメージを自作で作成するためのビルドファイルのこと
  ```
  例: rubyの環境をコンテナに焼き付けるためのDockerfile

  https://qiita.com/yu-croco/items/6f9d9e210c59c65f54f6
  ```Dockerfile
  # ベースとなるイメージを指定する
  FROM ruby:2.5

  # コンテナ上のワーキングディレクトリを指定する
  WORKDIR /usr/src/

  # ディレクトリやファイルをコピーする
  # 左側がホストのディレクトリ、右側がコンテナ上のディレクトリ
  COPY ./sample.rb /usr/src/sample.rb

  # "docker build"時に実行される処理
  RUN echo "building..."

  # "docker run"実行時に実行される処理
  CMD ruby sample.rb
  ```
  [日本語のサポートも充実している](http://docs.docker.jp/engine/reference/builder.html#from)

  Dockerfile内でサポートされているコマンドについて

  `FROM コマンド`: 以降の命令で使う ベース・イメージ を指定する。

    書式について
    ```
    # イメージだけを指定する場合はtagとしてlatestが採用される
    FROM <image>

    # タグを指定してイメージを取得する例
    FROM <image>:<tag>

    # タグと同じ位置付けに分類されこちらはsha256のハッシュを指定
    FROM <image>@<digest>
    ```
  `MAINTAINER コマンド`: 生成するイメージの Author （作者）フィールドを指定

    書式について
    ```
    MAINTAINER <Author>
    ```
  `RUN コマンド` : Linuxでは`/bin/sh`, Windowsでは`cmd/`コマンドの実行を行う

  `CMD コマンド` : `docker run` 時に実行を行う

    書式について
    ```
    CMD ["./exec.sh" "hello" "world"]
    ```

   [同位体としてENTRY_POINTというものも存在する](https://qiita.com/hihihiroro/items/d7ceaadc9340a4dbeb8f)

  `ENV コマンド` : コンテナに環境変数をセットする

    書式
    ```
    ENV <key>=<value>
    ```
  `ADD コマンド` : ホスト側のファイルをdockerイメージの指定したディレクトリにコピする。また、圧縮ファイルをコピーする際には`tar -x` と同じ働きを持つ

    書式
    ```
    ADD <source> <dist>
    ```
  `COPY コマンド` : ホスト側のファイルをdockerイメージ側で指定したディレクトリにコピーする。`ADD コマンド`のように圧縮ファイルを展開する働きはもたない。

    書式
    ```
    COPY <source> <dist>
    ```

  `WORKDIR コマンド` : Dockerfile で `RUN` 、 `CMD` 、 `ENTRYPOINT` 、 `COPY` 、 `ADD` 命令実行時の作業ディレクトリ（working directory）を指定する。もし WORKDIR が存在しなければ、 Dockerfile 命令内で使用しなくても`ディレクトリを作成`します。

    書式
    ```
    WORKDIR /path/to/workdir
    ```

  イメージが沸くように実際に[centos](https://hub.docker.com/_/centos)のイメージをdockerfileで作成しその中で
    whichバイナリの所在を判定してみます。

    適当なdirを$HOMEディレクトリ直下に作成してその配下にDockerfileを設置します。
    ```sh
    Zsh-> ~/centos-docker-image-dir
    TML[ZsH] tree
    .
    └── Dockerfile

    0 directories, 1 file
    ```
    Dockerfileの中身
    ```
    # イメージにcentosを指定する
    FROM centos:latest

    # このdockerfileの著者を示す
    MAINTAINER Y.K

    # whichバイナリをyumでインストール
    RUN yum install -y which

    # whichバイナリの所在を判定する
    RUN if [[ $(which which) ]];then echo "Y";else echo "N";fi
    ```
    では実際にDockerfileをbuildしimageを作成してみます。
    ```sh
    # カンレトにDockerfileを配置した場合、ビルドは build .で行う
    TML[ZsH] docker build .
    Sending build context to Docker daemon  2.048kB
    Step 1/4 : FROM centos:latest
    latest: Pulling from library/centos
    8a29a15cefae: Pull complete
    Digest: sha256:fe8d824220415eed5477b63addf40fb06c3b049404242b31982106ac204f6700
    Status: Downloaded newer image for centos:latest
     ---> 470671670cac
    Step 2/4 : MAINTAINER Y.K
     ---> Running in 00da57c70751
    Removing intermediate container 00da57c70751
     ---> c25ee8b2c55d
    Step 3/4 : RUN yum install -y which
     ---> Running in 458953b84da8
    CentOS-8 - AppStream                            3.1 MB/s | 7.0 MB     00:02
    CentOS-8 - Base                                 2.7 MB/s | 2.2 MB     00:00
    CentOS-8 - Extras                                13 kB/s | 5.9 kB     00:00
    Dependencies resolved.
    ================================================================================
     Package         Architecture     Version                Repository        Size
    ================================================================================
    Installing:
     which           x86_64           2.21-10.el8            BaseOS            49 k

    Transaction Summary
    ================================================================================
    Install  1 Package

    Total download size: 49 k
    Installed size: 83 k
    Downloading Packages:
    which-2.21-10.el8.x86_64.rpm                    558 kB/s |  49 kB     00:00
    --------------------------------------------------------------------------------
    Total                                           112 kB/s |  49 kB     00:00
    warning: /var/cache/dnf/BaseOS-f6a80ba95cf937f2/packages/which-2.21-10.el8.x86_64.rpm: Header V3 RSA/SHA256 Signature, key ID 8483c65d: NOKEY
    CentOS-8 - Base                                 1.6 MB/s | 1.6 kB     00:00
    Importing GPG key 0x8483C65D:
     Userid     : "CentOS (CentOS Official Signing Key) <security@centos.org>"
     Fingerprint: 99DB 70FA E1D7 CE22 7FB6 4882 05B5 55B3 8483 C65D
     From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial
    Key imported successfully
    Running transaction check
    Transaction check succeeded.
    Running transaction test
    Transaction test succeeded.
    Running transaction
      Preparing        :                                                        1/1
      Installing       : which-2.21-10.el8.x86_64                               1/1
      Running scriptlet: which-2.21-10.el8.x86_64                               1/1
      Verifying        : which-2.21-10.el8.x86_64                               1/1

    Installed:
      which-2.21-10.el8.x86_64

    Complete!
    Removing intermediate container 458953b84da8
     ---> 4d9f40c5ccfb
    Step 4/4 : RUN if [[ $(which which) ]];then echo "Y";else echo "N";fi
     ---> Running in 3113706cb0ab
    Y
    Removing intermediate container 3113706cb0ab
     ---> c3d97ff513cc
    Successfully built c3d97ff513cc
    ```
    無事にイメージを作成することができたので実際にイメージを確認してみます。
    ```sh
    docker images
    REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
    <none>              <none>              c3d97ff513cc        About an hour ago   270MB
    first-nginx         latest              bc9cde674744        12 hours ago        127MB
    nginx               latest              602e111c06b6        2 weeks ago         127MB
    centos              latest              470671670cac        3 months ago        237MB
    ```
    一番上に今回image名やtagを指定していないので`<none>`できてますね。
    せっかくなのでイメージ名とtagを付与してみます。
    ```sh
    $ docker tag <image-id> <image>:tag
    ```
    無事にリネームできていることが確認できました。
    ```sh
    TML[ZsH] docker images | grep centos-img
    centos-img          latest              c3d97ff513cc        2 hours ago         270MB
    ```
    それではコンテナを起動します。
    ```sh
    $ docker run -d -it --name centos-container centos-img:latest
    # アクティブなコンテナを確認する
    $ docker ps
    TML[ZsH] docker ps
    CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
    e8cd5e95f78d        centos-img:latest   "/bin/bash"         3 seconds ago       Up 2 seconds                            centos-container
    ```
    コンテナにログインして実際にwhichコマンドを正常に授受できるか試してみる
    ```sh
    $ docker exec -it $(docker ps -q) /bin/bash
    [root@e8cd5e95f78d /] which --help >/dev/null && echo $?
    0
    ```
