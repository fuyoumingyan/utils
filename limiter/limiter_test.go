package limiter

import (
	"fmt"
	"testing"
	"time"
)

func T() {
	time.Sleep(3 * time.Second)
}

func TestName(t *testing.T) {
	start := time.Now()
	limiter := New(50)
	for i := 0; i < 100000000; i++ {
		limiter.Add()
		println(len(limiter.limiter))
		go func() {
			T()
			limiter.Done()
		}()
	}
	limiter.WaitAndClose()
	elapsed := time.Since(start)
	fmt.Printf("myFunction 执行时间：%s\n", elapsed)
}
