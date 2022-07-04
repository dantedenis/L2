package model

import (
	"errors"
	_ "log"
	"sync"
)

// EventManager эвент менеджер для хранения и регистрации эвентов пользователей
type EventManager struct {
	sync.RWMutex
	Items map[string][]Model
}

// NewEventManager конструктор менеджера
func NewEventManager() *EventManager {
	var e EventManager
	e.Items = make(map[string][]Model)
	return &e
}

func (e *EventManager) CreateEvent(model *Model) ([]Model, error) {
	temp := e.Items[model.UserID]
	var err error

	e.RLock()
	for _, ev := range temp {
		if ev.ID == model.ID {
			err = errors.New("model already exist")
			break
		}
	}
	e.RUnlock()

	if err != nil {
		return nil, err
	}

	e.Lock()
	e.Items[model.UserID] = append(e.Items[model.UserID], *model)
	e.Unlock()
	return e.Items[model.UserID], nil
}

func (e *EventManager) DeleteEvent(model *Model) []Model {
	temp := e.Items[model.UserID]

	// ищем наш эвент и удаляем его из слайса эвентов
	for i, t := range temp {
		if model.ID == t.ID {
			e.Lock()
			temp = append(temp[:i], temp[i+1:]...)
			e.Unlock()
		}
	}

	// перезаписываем адрес на новый слайс, на всякий
	e.Lock()
	e.Items[model.UserID] = temp
	e.Unlock()
	return temp
}

func (e *EventManager) UpdateEvent(model *Model) []Model {
	temp := e.Items[model.UserID]

	// ищем запись лочим мьют и перезаписываем данные
	for i, t := range temp {
		if model.ID == t.ID {
			e.Lock()
			temp[i].Date = model.Date
			temp[i].Title = model.Title
			e.Unlock()
			break
		}
	}

	return e.Items[model.UserID]
}

// EventForDay возврат ошибки - если будет изменяться функционал (была мысль принимать модель и компоратор)
func (e *EventManager) EventForDay(model *Model) ([]Model, error) {
	result := make([]Model, 0)
	models := e.Items[model.UserID]

	// сканируем эвенты и забираем все совпадающие по дате
	e.RLock()
	defer e.RUnlock()
	for _, mod := range models {
		if mod.Date == model.Date {
			result = append(result, mod)
		}
	}

	// возврат слайса результатов
	return result, nil
}

// реализация аналогична методу выше
func (e *EventManager) EventForWeek(model *Model) ([]Model, error) {
	result := make([]Model, 0)
	models := e.Items[model.UserID]
	targetYear, targetWeek := model.Date.ISOWeek()

	e.RLock()
	defer e.RUnlock()

	for _, mod := range models {
		year, week := mod.Date.ISOWeek()
		if targetYear == year && targetWeek == week {
			result = append(result, mod)
		}
	}

	return result, nil
}

// аналогично
func (e *EventManager) EventForMonth(model *Model) ([]Model, error) {
	result := make([]Model, 0)
	models := e.Items[model.UserID]
	targetYear, targetMonth := model.Date.Year(), model.Date.Month()

	e.RLock()
	defer e.RUnlock()

	for _, mod := range models {
		year, month := mod.Date.Year(), mod.Date.Month()
		if targetYear == year && targetMonth == month {
			result = append(result, mod)
		}
	}

	return result, nil
}
