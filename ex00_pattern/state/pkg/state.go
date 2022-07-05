package pkg

import (
	"fmt"
)

type State interface {
	GetName() string
	Freeze(*StateContext)
	Heat(*StateContext)
}

// Context
type StateContext struct {
	state State
}

// Конструктор контекса, по умолчанию состояние Solid
func NewStateContext() *StateContext{
	return &StateContext{state: NewSolidState()}
}

// Сеттер состояния
func (s *StateContext) SetState(state State) {
	fmt.Println("Changing state to", state.GetName(), "...")
	s.state = state
}

// Методы для изменения состояния
func (s *StateContext) Freeze() {
	fmt.Println("Freezing", s.state.GetName(), "...")
	s.state.Freeze(s)
}

func (s *StateContext) Heat() {
	fmt.Println("Heating", s.state.GetName(), "...")
	s.state.Heat(s)
}

// Solid state
type SolidState struct {
	StateContext
}

// Конструктор состояния
func NewSolidState() *SolidState {
	return &SolidState{}
}

// Метод для текущего состояния, который ничего не делает
func (s *SolidState) Freeze(state *StateContext) {
	fmt.Println("Nothing happens")
}

// Метод для изменяет состояние на следующее
func (s *SolidState) Heat(state *StateContext) {
	state.SetState(NewLiquidState())
}

func (s *SolidState) GetName() string {
	return "Solid"
}

// Поведение методом аналогичны SolidState

// LiquidState 
type LiquidState struct {
	StateContext
}

func NewLiquidState() *LiquidState {
	return &LiquidState{}
}

func (l *LiquidState) GetName() string {
	return "Liquid"
}

func (l *LiquidState) Freeze(state *StateContext){
	state.SetState(NewSolidState())
}

func (l *LiquidState) Heat(state *StateContext) {
	state.SetState(NewGasState())
}

// GasState
type GasState struct {
	StateContext
}

func NewGasState() *GasState {
	return &GasState{}
}

func (g *GasState) GetName() string {
	return "Gas"
}

func (g *GasState) Freeze(state *StateContext){
	state.SetState(NewLiquidState())
}

func (g *GasState) Heat(state *StateContext){
	fmt.Println("Nothing happens")
}

