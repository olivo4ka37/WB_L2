package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
===Or channel===

Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}
Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))

*/

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
