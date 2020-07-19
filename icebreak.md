## IceBreakv9.0を実現するための調査ログ

### 0.前提条件

**※下記のユースケースに基づいての調査である。**

![](../../pu-resource/icebreak.png)

### 1.機能照会
- **1-1. ユーザマスタ編集機能[User Master Edit]**

  ```
  ユーザマスタテーブルへレコードの登録、削除、編集、更新を提供する機能である。
  ここで挙げているユーザとは待機社員を示す。この処理を実行できるアクターは管理者(Admin)に限る
  ```

- **1-2. カレンダマスタ編集機能[Calender Master Edit]**

  ```
  カレンダマスタテーブルへレコードの追加、削除、編集、更新を提供する機能である。
  ```

- **1-3. 勤怠情報管理機能[Attendance Master Edit]**

  ```
  ユーザが勤怠管理テーブルに勤怠情報の登録を行う機能である。
  ```

### 2.勤怠情報管理機能[Attendance Master Edit]の実現調査

- [ ] **勤怠情報管理機能[Attendance Master Edit]**

  現行のIceBreakWebSystemではICカードをR/Wに翳すことで`社員番号`と`社員名の取得`を行い、その情報を勤怠管理情報テーブルに格納する仕様となっているため。そのため今回のマイグレーションでもこのフローについては変えずに行いたい。そのためまずGo言語でも`R/Wを読み込むためのパッケージがサポートされているか`調査する必要がある。

  **根本的なアプローチの仕方について**
  ```
  特定のICカードのDLLをシステムコールか呼び出す方法で解決できそう
  ```
  尚、[Windows環境の場合](https://mattn.kaoriya.net/software/lang/go/20130805173059.htm)と[Linux環境の場合](https://github.com/cookieo9/goffi)で呼びだし方法が異なることもわかった。
そのため実際にコードに起こす際にはOSの判定処理を踏んでからDLLの読み込みを行う必要がある。

  **Windows環境下でDLLを読み込む例**

  下記の例では `user32.dll`をsyscallパッケージの`LoadDLL(path string)関数`で読み込み
  DLL内部でサポートされている[MessageBoxW関数](https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw)を実行している例です。

  ```go
  package main

  import (
      "log"
      "syscall"
      "unsafe"
  )

  func main() {
      dll, err := syscall.LoadDLL("user32.dll")
      if err != nil {
          log.Fatal(err)
      }
      defer dll.Release()

      proc, err := dll.FindProc("MessageBoxW")
      if err != nil {
          log.Fatal(err)
      }

      proc.Call(0,
          uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("メッセージ"))),
          uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("ボックス"))),
          0)
  ```

  **Linux環境でDLLを読み込む例**

  https://github.com/cookieo9/goffi

  ```go
  import "github.com/cookieo9/goffi/fcall"
  ...
    puts, _ := fcall.GetFunction("puts", SINT32, POINTER)
    cstr := fcall.CString("Hello, World!")
    defer fcall.Free(cstr)
    puts(cstr)
  ```
