package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"math"
	"time"
)

var client *resty.Client

func send() {
	client = resty.New()
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Minute)
		ticker.Stop()
		close(done)
	}()
	for {
		select {
		case <-done:
			log.Println("over")
			return
		case <-ticker.C:
			lauch(Sin(300, i, 30))
			i++

		}
	}

}
func lauch(i int) {
	for j := 0; j < i; j++ {
		go func() {
			rsp, err := client.R().
				EnableTrace().
				Get("http://127.0.0.1:50103/random")
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
