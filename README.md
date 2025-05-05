# HTTP Server

## Table of Content

- [Authors](#authors)
- [Examples](#examples)
- [Links](#links)

## Authors

| Name         | GitHub                                             |
|--------------|----------------------------------------------------|
| Klim Sidorov | [@entrlcom-klim](https://github.com/entrlcom-klim) |

## Examples

```go
package main

import (
	"net/http"

	"flida.dev/http-server"
)

func main() {
	// ❌.
	httpServer := &http.Server{
		Addr: "...",
	}

	// ✅.
	httpServer := http_server.NewDefaultHTTPServer()
	httpServer.Addr = "..."
}

```

## Links

* [So you want to expose Go on the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)
