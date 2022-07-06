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
	// ограничиваем рекусрию, разбиваем все каналы по парам, для мержа 2канала->1канал
	switch len(channels) {
	// завершаем тут рекусрию, если каналов не осталось вощвращаем закрытый
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return channels[0]
	// делим каналы по половине рекурсивно и в конечном итоге мержим попарно, при возврате вверх по рекусрии
	default:
		m := len(channels) / 2
		return mergeTwo(mergeRecursive(channels[:m]...),
			mergeRecursive(channels[m:]...))
	}
}

// Мержим 2 канала в 1
func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})

	// запускаем рутину
	go func() {
		// не забываем по завершению закрыть канал
		defer close(c)
		// чекаем значение канала, тк будем его изменять
		for a != nil || b != nil {
			select {
			// необходимо проерять на возможность чтения из канала, если не ок, то меняем значение канала
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			// аналогично, считываем и если не ок - меняем
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	// возвращаем смерженный канал
	return c
}
