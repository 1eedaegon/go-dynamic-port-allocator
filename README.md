# go-dynamic-port-allocator

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
