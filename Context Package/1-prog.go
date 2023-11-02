package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	start := time.Now()
	val, err := fetchData(ctx, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("time taken:",time.Since(start))

}
func fetchData(ctx context.Context, userID int) (int, error) {
	val, err := fetchThiredPartyStuffWichCanBeSlow()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func fetchThiredPartyStuffWichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500) //just for demo
	return 777, nil
}
