package resources

import (
	"bytes"
	"fmt"
	"sync"
)

const MEGABYTE = 1048576.0

var (
	memoryMutex  sync.Mutex
	memoryBuffer []byte
	memoryCount  float64 = 0
	memoryLimit  float64 = 0
)

func SetMemoryUsage(value float64) {
	if value > 0.0 {
		value = value / 100.0
	}

	memoryMutex.Lock()
	memoryCount = value * memoryLimit * MEGABYTE
	fmt.Printf("memory=%.2f bytes(%.1f%%)\n", memoryCount, value*100.0)
	memoryBuffer = bytes.Repeat([]byte{70}, int(memoryCount))
	memoryMutex.Unlock()
}

func ResetMemoryUsage() {
	memoryMutex.Lock()
	memoryBuffer = []byte{}
	memoryCount = 0
	memoryMutex.Unlock()
}

func SetMemoryLimit(limit float64) {
	memoryMutex.Lock()
	memoryLimit = limit
	memoryMutex.Unlock()
}

func GetMemoryUsage() float64 {
	memoryMutex.Lock()
	defer memoryMutex.Unlock()
	if memoryCount <= 0.0 || memoryLimit <= 0.0 {
		return 0.0
	}

	return ((memoryCount / memoryLimit) / MEGABYTE) * 100.0
}
