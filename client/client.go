package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"httpserver/types"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	base = []int{
		30, 21, 29, 31, 40, 48, 53, 47, 37, 39, 31, 29, 17, 9, 20, 24, 27, 35, 41, 38,
		27, 31, 27, 26, 21, 13, 21, 18, 33, 35, 40, 36, 22, 24, 21, 20, 17, 14, 17, 19,
		26, 29, 40, 31, 20, 24, 18, 26, 17, 9, 17, 21, 28, 32, 46, 33, 23, 28, 22, 27,
		18, 8, 17, 21, 31, 34, 44, 38, 31, 30, 26, 32,
	}
)
var (
	client       *resty.Client
	tickInterval time.Duration = 3 * time.Second
	sleepTime    time.Duration = 2 * time.Hour
	A                          = 10
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
			lauch(base[i%len(base)] + rand.Intn(7) - 3)
			i++

		}
	}

}
func lauch(i int) {
	for j := 0; j < i; j++ {
		go func(i, j int) {
			product := types.Product{
				ID:        j,
				Name:      fmt.Sprintf("%d/%d", i, j),
				Quantity:  rand.Intn(100),
				Timestamp: time.Now().Format("01-02 03:04:05"),
			}
			rsp, err := client.R().
				EnableTrace().SetBody(product).
				Post("http://192.168.49.2/app/product/create")
			if err != nil {
				log.Println(err)
				return
			}
			if rsp.StatusCode()/100 != 2 && rsp.StatusCode()/100 != 5 {
				log.Println(rsp.String())
				return
			}
			lock.Lock()
			statistic[rsp.Status()]++
			lock.Unlock()
			//fmt.Println(rsp.String())
		}(i, j)
	}
}
func Sin(n int, i int, A int) int {
	return int(float64(A) * (math.Sin(2*math.Pi/float64(n)*float64(i)) + 1))
}
func ShowWave(res []int) (string, string) {
	total := 0
	str := ""
	for i, v := range res {
		total += v
		if i == len(res)-1 {
			str += fmt.Sprintf("%d", v)
			break
		}
		str += fmt.Sprintf("%d,", v)
	}
	return fmt.Sprintf("[%s]", str), fmt.Sprintf("%d", total)
}

func MakeWave(high, low, n float64) []int {
	res := make([]int, 0)
	for i := 0; i < int(n); i++ {
		res = append(res, int((high-low)/2*(math.Sin(2*math.Pi/n*float64(i))+1)))
	}
	return res
}
