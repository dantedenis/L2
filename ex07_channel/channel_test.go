package ex07_channel

import (
	"fmt"
	"testing"
	"time"
)

var sig = func(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func TestMergeGor(t *testing.T) {
	start := time.Now()
	or := mergeRoutine

	<-or(
		sig(1*time.Second),
		sig(3*time.Second),
		sig(2*time.Second),
		sig(5*time.Second),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}

func TestMergeRecur(t *testing.T) {

	start := time.Now()
	or := mergeRecursive

	<-or(
		sig(1*time.Second),
		sig(3*time.Second),
		sig(2*time.Second),
		sig(5*time.Second),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}
