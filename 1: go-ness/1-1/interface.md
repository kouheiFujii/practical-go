interface は er を用いることがある
`io.Reader` 等

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

複数のメソッドをもつ interface はその目的を説明する名前にする

```go
type ResponseWriter interface {
	Header() Header

	Write([]byte) (int, error)

	WriteHeader(statusCode int)
}
```
