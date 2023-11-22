package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(time.Second)
	myErr := fmt.Errorf("my error")
	ctx, cancel := context.WithDeadlineCause(context.Background(), d, myErr)
	cancel()
	causedErr := context.Cause(ctx)
	fmt.Printf("causedErr: %v\n", causedErr) // causedErr: context canceled
}
