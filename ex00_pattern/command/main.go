package main

import "ex00_pattern/command/pkg"

/*
	Команда

	Поведенческий паттерн проектирования, который превращает запросы в объекты,
	позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь,
	логировать их, а также поддерживать отмену операций.

+:	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	Позволяет реализовать простую отмену и повтор операций.
	Позволяет реализовать отложенный запуск операций.
	Позволяет собирать сложные команды из простых.
	Реализует принцип открытости/закрытости.

-:	Усложняет код программы из-за введения множества дополнительных классов.
*/

func main() {
	queue := pkg.Broker{}
	win1 := pkg.NewWindow("TotalComm")
	win2 := pkg.NewWindow("IE5")
	win3 := pkg.NewWindow("This computer")

	openWin := pkg.OpenWindow{}
	closeWin := pkg.CloseWindow{}
	queue.AddCommand(openWin.Set(win1), openWin.Set(win2), closeWin.Set(win1), closeWin.Set(win2), openWin.Set(win3), closeWin.Set(win3))
	queue.RunCommand()
}
