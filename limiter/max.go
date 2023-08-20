package limiter

import "time"

type MaxLimiter struct {
	C chan struct{}
}

func NewM(cap int) *MaxLimiter {
	c := make(chan struct{}, cap)
	go func() {
		ticker := time.NewTicker(1000 * time.Millisecond)
		for range ticker.C {
			// 每秒钟往令牌桶中添加20个令牌
			for i := 0; i < 20; i++ {
				select {
				case c <- struct{}{}:
				default:
					// 如果令牌桶已满，则丢弃多余的令牌
				}
			}
		}
	}()
	return &MaxLimiter{
		C: c,
	}
}
func (l *MaxLimiter) Allow() bool {
	select {
	case <-l.C:
		return true
	default:
		return false
	}
}
