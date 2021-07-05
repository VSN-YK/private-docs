# Private Registoryの構築

[参考文献](https://qiita.com/Brutus/items/da63d23be32d505409c6)

```
private Registoryとは
privateなdockerイメージを一元管理するためのリポジトリサービス
dockerhubとlocalの間に位置しプライベートで管理したいdockerイメージをストックできる。
```

プライベートレジストリの概要図

![image](https://user-images.githubusercontent.com/60165356/81753909-cd427280-94ef-11ea-93f6-8bd34ded3af8.png)

## Private Registoryの作成

- [ ] dockerhubに公式のイメージがあがっているので`docker pull`する。

  ```sh
  # registryというイメージがdcokerhub上に存在するか確認する
  $ docker search registry > /dev/null && echo $?
  0
  # registryがdockerhub上に孫存在しているのを確認したたんめ、pullする。
  $ docker pull registry
  # イメージが正常にpullできたか確認する
  $ docker images | grep registry
  registry            latest              708bc6af7e5e        3 months ago        25.8MB
  ```

- [ ] pullしてきたregistryのイメージを元にコンテナを5000ポートに指定して起動する

  ```sh
  $ docker run -d -p 5000:5000 registry
  799daf3f7e88e345df916f74a53eb8ce4772f74361187d95d239f9b6c0c6ae8d
  ```
  これでPrivate Registryは構築できたので実際にイメージをpushしてみる

- [ ] Private Registryにイメージをpushするためにタグを修正する。

  ```sh
  # private registryにpushするためにはタグのフォーマットが以下である必要がある。
  $ docker tag <service on your registry:port>/<repo>/<image:tag>
  # 実際にイメージのtagを変更した時のコマンド例
  $ docker tag centos localhost:5000/repo/centos-v1:latest
  ```
- [ ] Private Registryにイメージをpushする

  ```sh
  $ docker push localhost:5000/repo/centos-v1:latest
  ```
- [ ] 実際にpushできたかローカルのイメージを削除し、レジストリからpullしてみる

  ```sh
  # ローカルのcentosイメージを削除
  $ docker rmi <centos-image>
  # プライベートレジストリからimageをpull
  $ docker pull localhost:5000/centos-image/centos-img-v1:latest
  latest: Pulling from centos-image/centos-img-v1
  8a29a15cefae: Pull complete
  a266ea50f064: Pull complete
  867f42fb7a24: Pull complete
  dc4eb2cea33d: Pull complete
  Digest: sha256:931b49d2e43089ceed8cc605c1cea9d3eec8de3b8d48317166000d37d7bc0552
  Status: Downloaded newer image for localhost:5000/centos-image/centos-img-v1:latest
  localhost:5000/centos-image/centos-img-v1:latest
  ```
  ローカルホストで建てたプライベートレジストリから無事にcentosのイメージをpullできましたここで一つ疑問がわきました。レジストリのイメージをUIで可視化してくれるアプリはないのかと...そしたらなんとありました。!!といっても[参考にした記事](https://qiita.com/Brutus/items/da63d23be32d505409c6)にはバッチリ書いてありましたけど

## ブラウザ上でレジストリをみる

- [ ] [本家のREADMEに記述されている](https://github.com/kwk/docker-registry-frontend)`docker run`コマンドで`docker-registry-frontend`コンテナを起動する。

  ```sh
  $ docker run \
  -d \
  -e ENV_DOCKER_REGISTRY_HOST=<ローカルホストのIP> \
  -e ENV_DOCKER_REGISTRY_PORT=<ポートフォワーディングした時のものを指定> \
  -p 8080:80 \
  # imageをpullしてきてない場合は自動的に必要なイメージをpullしてきてくれる。
  konradkleine/docker-registry-frontend:v2
  ```
- [ ] コンテナ正常に起動できているか確認する。

  ```sh
  $ docker ps | grep frontend
  c7915b47e70b        konradkleine/docker-registry-frontend:v2   "/bin/sh -c $START_S…"   13 hours ago        Up 13 hours         443/tcp, 0.0.0.0:8080->80/tcp   registry_browser
  ```
- [ ] frontend.appがサービスている:8080ポートにアクセスを行う

  [localhost:8080でアクセスを行って/homeにリダイレクトされるのは仕様](https://github.com/kwk/docker-registry-frontend/blob/acd3f1f55f0063c6496d26dd9a65d5d010dd6ff1/app/app.js#L87)
  ![image](https://user-images.githubusercontent.com/60165356/81879089-b28dfd80-95c4-11ea-8957-24ab1d489d26.png)

  ※　レジストリのイメージが反映されるまでに少し時間がかかります。
  
