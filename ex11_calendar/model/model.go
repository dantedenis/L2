package model

import (
	"time"
)

type Date struct {
	time.Time
}

type Model struct {
	Date   Date   `json:"date"`
	UserID string `json:"user_id"`
	//Лучше будет переделать на uuid
	ID    string `json:"event_id"`
	Title string `json:"title"`
}

// Parse Обертка для парсера даты
func Parse(t string) (Date, error) {
	timeReq, err := time.Parse("2006-01-02", t)
	return Date{timeReq}, err
}

// UnmarshalJSON Имплементируем интерфейс для анмаршела у времени
func (d *Date) UnmarshalJSON(date []byte) error {
	if string(date) == "" || string(date) == "null" {
		*d = Date{time.Now()}
		return nil
	}

	tm, err := time.Parse(`"`+"2006-01-02"+`"`, string(date))
	*d = Date{tm}
	return err
}
