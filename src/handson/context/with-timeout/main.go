package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bc := context.Background()
	func() {
		// timeout case
		t := 50 * time.Millisecond
		ctx, cancel := context.WithTimeout(bc, t)
		defer cancel()
		do(ctx)
	}()
	func() {
		// over slept case
		t := 3000 * time.Millisecond
		ctx, cancel := context.WithTimeout(bc, t)
		defer cancel()
		do(ctx)
	}()
}

func do(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
