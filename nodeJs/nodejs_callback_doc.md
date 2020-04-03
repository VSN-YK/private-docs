# Node.js コールバックの基本について理解する

## そもそもコールバックとは?...

他の関数に自身を実行してもらう関数のことである。

ここではJSの中で有名な`setTimeout(function() , milis)`関数についてコールバックの挙動を確認していく

https://www.scollabo.com/banban/jsindex/sample/sample_254.html

```js
setTimeout(function(){ console.log('Hello JS World');},3000)
```
## コールバックの呼び出され方

- [ ] Step1 3秒(3000ms)待機する処理が開始される

- [ ] Step2 その後、consoleに`Hello JS World`が表示される

  ```js
  function(){ console.log('Hello JS World');}
  ```

  また、function()の部分は `=>` (アロー演算子)で代用できる

  ```js
  setTimeout(() => {console.log('Hello Js CallBack');},3000);
  ```
  実際の挙動

  ```sh
  $ node basicCallBackSample.js
  # 3秒後にHello Js CallBackが表示される
  Hello Js CallBack
  ```
  このように別の関数の処理が始動条件で呼び出される関数のことを`コールバック関数`と呼ぶ。

  かなり割愛したので詳細が知りたい方は以下の記事が参考になるかと思います。
  https://sbfl.net/blog/2019/02/08/javascript-callback-func/

  しかし、jsではCallBackHellと呼ばれるコールバック関数のネスト地獄の問題があり近年ではあまり使用されないようになってきました。
  そのため、最近ではPromiseを利用した非同期処理が主流となっております。

## Promiseを利用した非同期処理

### Promiseとは

[公式の定義](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Promis)

```
Promise オブジェクトは非同期処理の最終的な完了処理（もしくは失敗）およびその結果の値を表現します。
```
ちょっと抽象的な表現が多くてわかり憎かったのでコードに起こしてみました。
この例は先ほどの同様の`setTimeout(function() , milis)`を使用してPromiseの利点について解説いたします。

```js
function asyncronusFunction(num){
    //Promise オブジェクトのリターンを行う
    return new Promise(function(resolve ,reject) {
        var sw = Math.floor(Math.random() * Math.floor(num));
        setTimeout(function() {
            console.log(sw);
            if (sw == 1 || sw == 2) {
                resolve('Succees Promise Process');
            }else{
                reject('Faild');

            }
        },1000);
    });
}

asyncronusFunction(3).then(function (value) {
    console.log(value);
}).catch(function (error) {
    console.log(error);
});
```

処理の内容を解説(関係のある処理だけを抜粋してます)

- [ ] promise オブジェクトの作成(1sスリープした後、コンソールにメッセージを出力するコールバック)を行う

  ```js
  return new Promise(function(resolve ,reject) {
      var sw = Math.floor(Math.random() * Math.floor(num));
      setTimeout(function() {
          console.log(sw);
          if (sw == 1 || sw == 2) {
              //Return Expect of Value
              resolve('Succees Promise Process');
          }else{
              //coused Error
              reject('Faild');

          }
      },1000);
  });

  ```
- [ ] promiseオブジェクトの結果を受け取り次のイベントを始動
  numの値が1or2の場合はresolve()を返し、0,3の場合はreject()を返す仕様です。
  ```js
  asyncronusFunction(3).then(function (value) {
      console.log(value);
  }).catch(function (error) {
      console.log(error);
  });

  ```

**※ 定義や性格については後ほど述べます。**

まず、Promiseを理解する上で大切になるキーワードは`resolve`と`reject`です。

`resolve`というのはPromiseで定めた振る舞いが担保されている場合(今回の場合は3sスリープした後にコンソールに`Like [Sync Process]  with Promiss`が表示される)に結果の承諾を行うPromiseに付随したメソッドになります。
これにより、Promiseのプロセスは正常に実行できたことを通知することができます。

**正確には通知ではなく、PromiseオブジェクトのReturnを行う**

[MS公式のコード](https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/Promise/resolve)を
改良し`Promise.resolve()`の挙動を確認

```js
const promiseObj = Promise.resolve('Completed Promise Process');

console.log(promiseObj)

promiseObj.then(function(value) {
  console.log(value);
  // expected output: Completed Promise Process
});
```
Console
```
> [object Promise] //Promise オブジェクトがリターンされていることがわかる
> "Completed Promise Process"
```

ですので `reject`はその逆でエラーハンドリングを行うための制御メソッドになります。`reject()`を発動させることで `new Error()`を生成できます。

```js
reject('Faild');
```
つまり`catch((e =>)`できるということになります。

#### Promiseが保証する事項

- 現在の JavaScript イベントループの実行完了より前には、コールバックが決して呼び出されない。

- 非同期処理が完了もしくは失敗した後に then() により登録されたコールバックでも、上記のように呼び出される。

- then() を何回も呼び出して複数のコールバックを追加してもよく、それぞれのコールバックは追加順に独立して実行される。

#### Promiseが保有する状態

- pending: 初期状態
- fulfilled: 処理が成功して完了
- rejected: 処理が失敗

