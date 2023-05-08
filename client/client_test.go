package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	send()
	//client = resty.New()
	////client.R().
	//rsp, err := client.R().
	//	EnableTrace().
	//	Get("http://127.0.0.1:35477/random")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(string(rsp.Body()))
}
func TestShowWave(t *testing.T) {
	res := ShowWave()
	fmt.Println(res)
}
func TestLocal(t *testing.T) {
	client = resty.New()
	//client.R().
	wg := sync.WaitGroup{}
	for j := 0; j < 3; j++ {
		for i := 0; i < 21; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				rsp, err := client.R().Get("http://127.0.0.1:50821/random")
				if err != nil {
					log.Println(err)
				}
				fmt.Println(i, " : ", rsp.String())
			}(i)
		}
		wg.Wait()
		time.Sleep(15 * time.Second)
	}

}
