package ex07_channel

import (
	"sync"
)

// Решить проблему помогла статья на medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de

// Решение с помощью горутин
func mergeRoutine(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	// обязательно закрываем канал, иначе Deadlock
	defer close(out)
	wg := sync.WaitGroup{}
	//Проход по всем каналам, и перенаправляем их в out
	for _, c := range channels {
		wg.Add(1)
		//Передаем по значению, защитить от DataRace
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}
	// Ожидаем завершение работы группы каналов
	wg.Wait()
	return out
}

// Решение с помощью рекурсии
func mergeRecursive(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return channels[0]
	default:
		m := len(channels) / 2
		return mergeTwo(mergeRecursive(channels[:m]...),
			mergeRecursive(channels[m:]...))
	}
}

func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}
