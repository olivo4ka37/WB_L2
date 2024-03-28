package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1000*time.Millisecond),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(9999*time.Millisecond),
		sig(12*time.Second),
	)

	fmt.Printf("time after start %v", time.Since(start))
}

// or объединяет несколько каналов в один и возвращает его
func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}

	combined := make(chan interface{})

	go func() {
		defer close(combined)

		cases := make([]reflect.SelectCase, len(channels))
		for i, ch := range channels {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}

		for len(cases) > 0 {
			chosen, value, ok := reflect.Select(cases)
			if !ok { // Channel closed
				cases = append(cases[:chosen], cases[chosen+1:]...)
				continue
			}

			combined <- value.Interface()
		}
	}()

	return combined
}
