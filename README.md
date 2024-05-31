# HTTP server

## Table of Content

- [Examples](#examples)
- [License](#license)
- [Links](#links)

## Examples

```go
package main

import (
	"context"

	"entrlcom.dev/http-server"
)

func main() {
	ctx := context.Background()

	// HTTP server.
	httpServer := http_server.NewHTTPServer()
	httpServer.Addr = ""     // TODO: set address.
	httpServer.Handler = nil // TODO: set handler.

	if err := httpServer.ListenAndServe(); err != nil {
		// TODO: Handle error.
		return
	}

	// ...

	// Shutdown HTTP server.
	if err := httpServer.Shutdown(ctx); err != nil {
		// TODO: Handle error.
		return
	}

	// ...
}

```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Links

* [So you want to expose Go on the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)
