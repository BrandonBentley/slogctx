# slogctx

slogctx is a simple context logger factory implementation for the slog logger introduced in go 1.21

There are 2 ways to use this module. First you can create a `Factory` and pass it to all places that need it like so:
``` go
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/BrandonBentley/slogctx"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(handler)

	factory := slogctx.NewFactory(logger)

	ctx := context.Background()

	factory.GetContextLogger(ctx).Info(
		"here is a info message",
	)
}
```
Output
``` json
{"time":"2023-11-27T23:32:59.542452-07:00","level":"INFO","msg":"here is a info message"}
```

Or you can set the root factory and use the global functions:

``` go
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/BrandonBentley/slogctx"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(handler)

	factory := slogctx.NewFactory(logger)

	slogctx.SetRootLoggerFactory(factory)

	ctx := context.Background()

	slogctx.GetContextLogger(ctx).Info(
		"here is a info message from the root factory",
	)
}

```
Output
``` json
{"time":"2023-11-27T23:34:15.939646-07:00","level":"INFO","msg":"here is a info message from the root factory"}
```

or finally you can just use the default logger factory. If no default logger is set it will use whatever `slog.Default()` returns. The default `*slog.Logger` is a `TextHandler`

``` go
package main

import (
	"context"

	"github.com/BrandonBentley/slogctx"
)

func main() {
	ctx := context.Background()

	slogctx.GetContextLogger(ctx).Info(
		"here is a info message from the default root factory",
	)
}
```
Output
``` 
2023/11/27 23:37:45 INFO here is a info message from the default root factory
```