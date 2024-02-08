package limiter

import "sync"

type Limiter struct {
	limiter chan struct{}
	wg      sync.WaitGroup
}

func (l *Limiter) Add(num ...int) {
	n := 1
	if len(num) > 0 {
		n = num[0]
	}
	for i := 0; i < n; i++ {
		l.limiter <- struct{}{}
		l.wg.Add(1)
	}
}

func (l *Limiter) Done() {
	// 先释放 再 wg.Done 防止 Done 后 wg.Wait 立即返回 而其他的还没有释放完全
	<-l.limiter
	l.wg.Done()
}

func (l *Limiter) WaitAndClose() {
	l.wg.Wait()
	close(l.limiter)
}

func New(size int) *Limiter {
	return &Limiter{
		limiter: make(chan struct{}, size),
	}
}
