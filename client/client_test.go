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
	res, total := ShowWave(MakeWave(5, 0, 10))
	fmt.Println(res)
	fmt.Println("total: ", total)
}
func TestLocal(t *testing.T) {
	client = resty.New()
	//client.R().
	ftotal := 0
	wg := sync.WaitGroup{}
	for j := 0; j < 3; j++ {
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				rsp, err := client.R().Get("http://192.168.49.2/app/random")
				if err != nil {
					log.Println(err)
				}
				fmt.Println(i, " : ", rsp.String())
				if rsp.StatusCode() == 500 {
					ftotal++
				}
			}(i)
		}
		wg.Wait()
		time.Sleep(1 * time.Second)
	}
	fmt.Println("ftotal", ftotal)
}
func TestMakeWave(t *testing.T) {
	fmt.Println(MakeWave(10, 0, 10))
}
