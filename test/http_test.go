package test_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Reqeast(url string) (res string) {
	res = ""
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Context-Type", "json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	res = string(body)
	fmt.Println(res)
	return
}

func TestReq(t *testing.T) {
	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():

				fmt.Println("------------end")
				return
			default:
				Reqeast("http://127.0.0.1:8282/cart_map")
			}
		}(ctx)
	}

	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(100 * time.Second)

}

func Sort() {

}
