# go-dynamic-port-allocator

[![Go](https://pkg.go.dev/badge/github.com/1eedaegon/go-dynamic-port-allocator.svg)](https://pkg.go.dev/github.com/1eedaegon/go-dynamic-port-allocator)
[![CI](https://github.com/1eedaegon/go-dynamic-port-allocator/actions/workflows/go.yml/badge.svg)](https://github.com/1eedaegon/go-dynamic-port-allocator/actions/workflows/go.yml)
[![CodeQL](https://github.com/1eedaegon/go-dynamic-port-allocator/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/1eedaegon/go-dynamic-port-allocator/actions/workflows/codeql.yml)

A Go library for allocate ports, dynamically.

## Example

```go
import (
	...
	port "github.com/1eedaegon/go-dynamic-port-allocator"
	...
)

ports := port.Get(3)
// ports is something like []int{10000, 10001, 10002}
```

Or

```go
import (
	...
	port "github.com/1eedaegon/go-dynamic-port-allocator"
	...
)

ports := port.GetS(3)
// ports is something like []string{"10000", "10001", "10002"}
```

## License

[MIT](LICENSE)
