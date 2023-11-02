package main

import (
	"context"
	"fmt"
)

func main() {
	key := "user"
	var val int
	ctx := context.WithValue(context.Background(), key, val)
	n, err := checkContext(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}

func checkContext(ctx context.Context) (any, error) {
	val := ctx.Value("use")
	if val == nil {
		return 0, fmt.Errorf("Key is not matched:")
	}
	return val, nil
}
