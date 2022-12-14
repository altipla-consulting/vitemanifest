
# vitemanifest

[![Go Reference](https://pkg.go.dev/badge/github.com/altipla-consulting/vitemanifest.svg)](https://pkg.go.dev/github.com/altipla-consulting/vitemanifest)

Read the Vite manifest from Go.


## Install

```shell
go get github.com/altipla-consulting/vitemanifest
```


## Usage

Read the manifest from its standard location when starting up the application:

```go
var manifest vitemanifest.File

func init() {
	manifest, err = vitemanifest.Read()
	if err != nil {
		log.Fatal(err)
	}
}
```

You can also specify a custom location for the manifest file:

```go
manifest, err := vitemanifest.ReadPath("mycustom/file")
if err != nil {
  log.Fatal(err)
}
```


## Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using `make gofmt`.


## License

[MIT License](LICENSE)
