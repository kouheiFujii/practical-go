- Go 言語では、複数のソースコードをまとめて一つのパッケージとして扱うことができます。
- Go 言語では、他のパッケージを import することで、そのパッケージの機能を使用することができます。
- 各パッケージ間で循環 import することは許されていません。循環 import するとエラーが発生します。
- Go 言語では、必要なパッケージ間の依存関係を明確にすることが推奨されます。これにより、各パッケージの責務が明確になり、再利用性が高まります。

循環参照のサンプルコード

```go
// main.go
package main

import (
	"./a"
)

func main() {
	a.Hello()
}

// a/a.go
package a

import (
	"../main"
)

func Hello() {
	println("Hello from package a")
	main.Main()
}
```

a パッケージと main パッケージで参照が循環しているのでエラーが生じる

```bash
import cycle not allowed: main -> a -> main
```
