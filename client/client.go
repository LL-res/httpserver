package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"math"
	"time"
)

var (
	client       *resty.Client
	tickInterval time.Duration = 15 * time.Second
	sleepTime    time.Duration = 5 * time.Minute
)

func send() {
	client = resty.New()
	ticker := time.NewTicker(tickInterval)
	i := 0
	done := make(chan struct{})
	go func() {
		time.Sleep(sleepTime)
		ticker.Stop()
		close(done)
	}()
	for {
		select {
		case <-done:
			log.Println("over")
			return
		case <-ticker.C:
			lauch(Sin(int(sleepTime.Seconds()/tickInterval.Seconds()), i, 30))
			i++

		}
	}

}
func lauch(i int) {
	for j := 0; j < i; j++ {
		go func() {
			rsp, err := client.R().
				EnableTrace().
				Get("http://127.0.0.1:43787/random")
			if err != nil {
				log.Println(err)
			}
			fmt.Println(rsp.String())
		}()
	}
}
func Sin(n int, i int, A int) int {
	return A * int(math.Sin(2*math.Pi/float64(n*i))+1)
}
