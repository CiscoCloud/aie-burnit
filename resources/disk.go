package resources

import (
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
	ResetDiskUsage()

	diskCount = value
	running = true

	go func() {
		for {
			diskMutex.Lock()

			if running == false {
				continue
			}

			d1 := randBytes(diskCount * MEGABYTE)
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

func GetDiskUsage() int64 {
	diskMutex.Lock()
	defer diskMutex.Unlock()
	if diskCount <= 0.0 {
		return 0.0
	}

	return diskCount
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
