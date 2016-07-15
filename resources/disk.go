package resources

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	diskMutex sync.Mutex
	diskCount int64 = 0
	running   bool  = false
)

func SetDiskUsage(value int64) {
	if value > 0.0 {
		value = value / 100.0
	}

	ResetDiskUsage()

	running = true

	go func() {
		for {
			diskMutex.Lock()

			if running == false {
				continue
			}

			diskCount = value * MEGABYTE
			fmt.Printf("disk=%.2f bytes(%.1f%%)\n", diskCount, value*100.0)

			d1 := randBytes(diskCount)
			ioutil.WriteFile("/tmp/tempfile.dat", d1, 0644)

			diskMutex.Unlock()

			time.Sleep(time.Duration(1) * time.Second)
		}
	}()
}

func ResetDiskUsage() {
	diskMutex.Lock()

	running = false

	diskCount = 0

	diskMutex.Unlock()
}

func GetDiskUsage() float64 {
	diskMutex.Lock()
	defer diskMutex.Unlock()
	if diskCount <= 0.0 {
		return 0.0
	}

	return float64((diskCount / MEGABYTE)) * 100.0
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func randBytes(n int64) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return b
}
