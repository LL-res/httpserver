package client

import (
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

var client *resty.Client

func send() {
	client = resty.New()
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	done := make(chan struct{})
	go func() {
		time.Sleep(30 * time.Second)
		ticker.Stop()
		close(done)
	}()
	for {
		select {
		case <-done:
			log.Println("over")
			return
		case <-ticker.C:
			lauch(i)
			i++

		}
	}

}
func lauch(i int) {
	for j := 0; j < i; j++ {
		go func() {
			_, err := client.R().
				EnableTrace().
				Get("http://127.0.0.1:32345/random")
			if err != nil {
				log.Println(err)
			}
		}()
	}
}
