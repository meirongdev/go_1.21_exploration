package main

import (
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	newCtx := context.WithoutCancel(ctx)
	cancel()
	<-ctx.Done()
	<-newCtx.Done() // fatal error: all goroutines are asleep - deadlock!
}
