package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

func main() {
	ctx := context.Background()
	start := time.Now()
	val, err := fetchUserData(ctx, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("time taken:", time.Since(start))
}
func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	respo := make(chan Response)

	go func() {
		val, errs := fetchThiredPartyStuffWhichCanBeSlow()
		respo <- Response{
			value: val,
			err:   errs,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("Taking Longer then usual from thired party")
		case resp := <-respo:
			return resp.value, resp.err
		}
	}

}

func fetchThiredPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 100) //just for demo
	return 777, nil
}
