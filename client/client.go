package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"math"
	"sync"
	"time"
)

var (
	client       *resty.Client
	tickInterval time.Duration = 1 * time.Second
	sleepTime    time.Duration = 5 * time.Minute
	A                          = 30
	lock         sync.Mutex
	statistic    map[string]int
)

func send() {
	statistic = make(map[string]int)
	client = resty.New()
	ticker := time.NewTicker(tickInterval)
	i := 0
	done := make(chan struct{})
	go func() {
		time.Sleep(sleepTime)
		ticker.Stop()
		fmt.Println(statistic)
		close(done)
	}()
	for {
		select {
		case <-done:
			log.Println("over")
			return
		case <-ticker.C:
			lauch(Sin(int(sleepTime.Seconds()/tickInterval.Seconds()), i, A))
			i++

		}
	}

}
func lauch(i int) {
	for j := 0; j < i; j++ {
		go func() {
			rsp, err := client.R().
				EnableTrace().
				Get("http://127.0.0.1:51451/random")
			if err != nil {
				log.Println(err)
				return
			}
			lock.Lock()
			statistic[rsp.Status()]++
			lock.Unlock()
			//fmt.Println(rsp.String())
		}()
	}
}
func Sin(n int, i int, A int) int {
	return int(float64(A) * (math.Sin(2*math.Pi/float64(n)*float64(i)) + 1))
}
func ShowWave() string {
	n := int(sleepTime.Seconds() / tickInterval.Seconds())
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, Sin(n, i, A))
	}
	str := ""
	for i, v := range res {
		if i == len(res)-1 {
			str += fmt.Sprintf("%d", v)
			break
		}
		str += fmt.Sprintf("%d,", v)
	}
	return fmt.Sprintf("[%s]", str)
}
