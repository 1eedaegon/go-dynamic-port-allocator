# go-dynamic-port-allocator

[![Go](https://pkg.go.dev/badge/github.com/1eedaegon/go-dynamic-port-allocator.svg)](https://pkg.go.dev/github.com/1eedaegon/go-dynamic-port-allocator)
![workflow](https://github.com/1eedaegon/go-dynamic-port-allocator/actions/workflows/go.yml/badge.svg)

A Go library for allocate ports, dynamically.

## Example

```go
ports := port.Get(3)
// ports is something like []int{10000, 10001, 10002}
```

Or

```go
ports := port.GetS(3)
// ports is something like []string{"10000", "10001", "10002"}
```

## License

[MIT](LICENSE)
