package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"
)

func main() {
	Slices()

	Defers()

	Interfaces()

	ChanDeadlock()

	ErrorInterface()

	RelocateSlice()

	MergeChanErr()
}

// 1.
func Slices() {
	// таки образом инициализируется массив
	a := [5]int{76, 77, 78, 79, 80}

	// а тут мы создаем срез от массива а
	var b []int = a[1:4]
	fmt.Println(b)
}

// 2.
func Defers() {
	fmt.Println(testDef())
	fmt.Println(anotherTestDef())
}

func testDef() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTestDef() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

// 3.
func Interfaces() {
	var temp interface{}
	err := foo()
	fmt.Println(err)
	// эта операция будет true тогда, и только тогда, когда оба поля в интерфейсе будут нил 
	// в интерфейсе 2 поля: tab и data : информация о конкретном типа и ссылка на данные (соответственно)
	fmt.Println(err == nil)

	// Наглядно увидим это с помощью пакета reflect
	fmt.Println("Type:", reflect.TypeOf(err), "Data:", reflect.ValueOf(err))
	fmt.Println("Type:", reflect.TypeOf(temp), "Data:", reflect.ValueOf(temp))
	temp = err
	fmt.Println("Type:", reflect.TypeOf(temp), "Data:", reflect.ValueOf(temp))
}

func foo() error {
	var err *os.PathError = nil
	return err
}

// 4.
func ChanDeadlock() {
	ch := make(chan int)
	go func() {
		// add close canal
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// Deadlock - канал не закрыт, ждем значения, а никто туда не пишет
	for n := range ch {
		fmt.Println(n)
	}
}

//5.
func ErrorInterface() {
	var err error
	// error это интерфейс, аналогично как в 4, мы записываем в поле тип, но поле дата по прежнему nil
	err = test()
	// и сравнивая интерфейс с nil мы получим false
	fmt.Println(err != nil, err == nil)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("ok")
}

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	return nil
}

//6.
func RelocateSlice() {
	// Слайс это структура с 3 полями: указатель на массив, длинна и емкость
	var s = []string{"1", "2", "3"}
	// передаем слайс-структуру
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	// ображаемся к 0 элементу массиву и меняешь значение
	i[0] = "3"
	// релоцируем массив, тем самым в копии слайса i уже указатель на другую область памяти
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
	// до конца программы мы работает с новым массивом, не затронув тот, который был в s после релокации
}

//7.
func MergeChanErr() {
	a := asChan(1, 2, 4, 5, 5)
	b := asChan(2, 4, 6, 7, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	/*
		go func() {
			for {
				select {
					// бесконечно пытаемся читать с канала, которые закрыты, необходимо проверить,
					// удается ли считать с канала, и перестать с него считывать если он закрыт
					case v := <-a:
						c <-v
					case v := <-b:
						c <-v
				}
			}
		}()
	*/

	// рабочий вариант: отслеживаем вазможность считывания с канала и меняем значение
	// плюс не забываем закрыть канал, что бы не было Deadlock'a
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
