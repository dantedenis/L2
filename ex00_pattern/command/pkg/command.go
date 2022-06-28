package pkg

import "fmt"

// что бы обощить наши структуры, которые по своему управляют чем-либо
type Executor interface {
	Exec()
}

// допустим это окно
type Window struct {
	name string
}

func NewWindow(name string) *Window {
	return &Window{name: name}
}

func (w *Window) Open() {
	fmt.Println("2xClick Open to window:", w.name)
}

func (w *Window) Close() {
	fmt.Println("Click close window:", w.name)
}

// Превращаем наши запросы в объекты
type OpenWindow struct {
	window *Window
}

func (o OpenWindow) Set(w *Window) *OpenWindow {
	o.window = w
	return &o
}

// И реализуем интерфейс
func (o *OpenWindow) Exec() {
	o.window.Open()
}

type CloseWindow struct {
	window *Window
}

func (c CloseWindow) Set(w *Window) *CloseWindow {
	c.window = w
	return &c
}

func (c *CloseWindow) Exec() {
	c.window.Close()
}

// Будет очередь событий и пару методов для добавления и выполнения
type Broker struct {
	commands []Executor
}

func (b *Broker) AddCommand(command ...Executor) {
	b.commands = append(b.commands, command...)
}

func (b *Broker) RunCommand() {
	for _, run := range b.commands {
		run.Exec()
	}
}
