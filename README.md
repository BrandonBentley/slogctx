# slogctx

slogctx is a simple context logger implementation for the slog logger introduced in go 1.21

To use this module in your project run:
``` bash
go get github.com/BrandonBentley/slogctx
```

There are a few ways to use this module.

The recommended way is to set the slog default logger to use the `JSONContextHandler`. This way all logs will use the context handler, which if no values are in the context, behaves as a `slog.JSONHandler`. See the code below for example usage:

```go
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/BrandonBentley/slogctx"
)

func main() {
	slog.SetDefault(slog.New(slogctx.NewJSONContextHandler(os.Stdout, nil)))

	ctx := context.Background()

	ctx = slogctx.WithAttrs(
		ctx,
		slog.String("someKey", "someVal"),
	)

	slog.InfoContext(
		ctx,
		"I have some info for you",
	)
}
```

Output

```json
{
  "time": "2025-04-11T01:22:31.706502-06:00",
  "level": "INFO",
  "msg": "I have some info for you",
  "someKey": "someVal"
}
```
---
### sloginit
You can also use the sloginit package provided to initialize the default slog logger with the `slogctx.JSONContextHandler` and not require any code for initializing

```go
package main

import (
	"context"
	"log/slog"

	"github.com/BrandonBentley/slogctx"
	_ "github.com/BrandonBentley/slogctx/sloginit"
)

func main() {
	ctx := context.Background()

	ctx = slogctx.WithAttrs(
		ctx,
		slog.String("Key1", "Val1"),
	)

	slog.InfoContext(
		ctx,
		"I have some info for you",
	)
}
```

Output

```json
{
  "time": "2025-04-11T01:43:33.14336-06:00",
  "level": "INFO",
  "msg": "I have some info for you",
  "Key1": "Val1"
}
```

### No More Duplicates!

A significant update: `slogctx.ContextHandler` will now remove duplicate keys, keeping the last provided value

```go
package main

import (
	"context"
	"log/slog"

	"github.com/BrandonBentley/slogctx"
	_ "github.com/BrandonBentley/slogctx/sloginit"
)

func main() {
	firstContext := slogctx.WithAttrs(
		context.Background(),
		slog.String("Key1", "Val1"),
	)

	ctx := firstContext
	slog.ErrorContext(
		ctx,
		"Original Value",
	)

	ctx = slogctx.WithAttrs(
		ctx,
		slog.String("key2", "val2"),
		slog.String("Key1", "IChangedIt"),
	)

	slog.InfoContext(
		ctx,
		"I Changed the first one",
	)

	ctx = slogctx.With(
		ctx,
		"key2", 7,
	)

	slog.InfoContext(
		ctx,
		"You can see the second one is now a number",
	)

	slog.WarnContext(
		firstContext,
		"the first context is unchanged",
	)
}

```

Output

```json
{
    "time": "2025-04-11T01:53:29.677325-06:00",
    "level": "ERROR",
    "msg": "Original Value",
    "Key1": "Val1"
}
{
    "time": "2025-04-11T01:53:29.677446-06:00",
    "level": "INFO",
    "msg": "I Changed the first one",
    "Key1": "IChangedIt",
    "key2": "val2"
}
{
    "time": "2025-04-11T01:53:29.677449-06:00",
    "level": "INFO",
    "msg": "You can see the second one is now a number",
    "Key1": "IChangedIt",
    "key2": 7
}
{
    "time": "2025-04-11T01:53:29.677472-06:00",
    "level": "WARN",
    "msg": "the first context is unchanged",
    "Key1": "Val1"
}
```
