// Package port exposes an unused port when execute: Get(n int)
package port

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	minPort   = 10000
	maxPorts  = 65535
	blockSize = 1024
	maxBlocks = 16
	attempts  = 10
)

var (
	port      int
	firstPort int
	once      sync.Once
	mu        sync.Mutex
	lnLock    sync.Mutex
)

// Get returns n ports that are free to use, panicing if it can't succeed.
func Get(n int) []int {
	ports, err := GetWithErr(n)
	if err != nil {
		panic(err)
	}
	return ports
}

// GetS returns n ports as strings that are free to use, panicing if it can't succeed.
func GetS(n int) []string {
	ports, err := GetSWithErr(n)
	if err != nil {
		panic(err)
	}
	return ports
}

// GetS return n ports (as strings) that are free to use.
func GetSWithErr(n int) ([]string, error) {
	ports, err := GetWithErr(n)
	if err != nil {
		return nil, err
	}
	var portsStr []string
	for _, port := range ports {
		portsStr = append(portsStr, strconv.Itoa(port))
	}
	return portsStr, nil
}

// Get returns n ports that are free to use.
func GetWithErr(n int) (ports []int, err error) {
	mu.Lock()
	defer mu.Unlock()

	if n > blockSize-1 {
		return nil, fmt.Errorf("port-allocator: block size is too small for ports requested")
	}

	once.Do(initialize)

	for len(ports) < n {
		port++
		if port > maxPorts {
			return ports, fmt.Errorf("port-allocator: ports has allocated but not enough want: %d, got: %d", n, len(ports))
		}
		if port < firstPort+1 {
			port = firstPort + 1
		}

		ln, err := listen(port)
		if err != nil {
			continue
		}
		ln.Close()

		ports = append(ports, port)
	}

	return ports, nil
}

func initialize() {
	if uint(maxPorts-minPort)+minPort > maxPorts {
		panic("port-allocator: minimum port is too big")
	}
	srcFromSeed := rand.NewSource(time.Now().UnixNano())
	var err error
	// Find first port
	for i := 0; i < attempts; i++ {
		firstPort = int(rand.New(srcFromSeed).Int31n(int32(maxPorts-minPort))) + minPort
		_, err = listen(firstPort)
		if err != nil {
			continue
		}
		return
	}
	panic("port-allocator: can't allocated port block")
}

func listen(port int) (*net.TCPListener, error) {
	return net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: port})
}
