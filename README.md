
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
package main

import (
  "log"

  "github.com/altipla-consulting/vitemanifest"
)

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
package main

import (
  "log"

  "github.com/altipla-consulting/vitemanifest"
)

func init() {
	manifest, err := vitemanifest.ReadPath("mycustom/file")
	if err != nil {
    log.Fatal(err)
	}

  log.Println(manifest)
}
```
