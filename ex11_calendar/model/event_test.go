package model

import (
	"fmt"
	"testing"
	"time"
)

func TestAddEvent(t *testing.T) {
	ev := NewEventManager()
	mod1 := Model{Date: Date{time.Now()}, UserID: "eqweqw"}
	//mod2 := Model{Date: time.Now(), UserID: "f1231"}
	//mod3 := Model{Date: time.Now(), UserID: "12123123"}
	_, err := ev.CreateEvent(mod1)
	if err != nil {
		return
	}
	_, err = ev.CreateEvent(mod1)
	if err != nil {
		return
	}
	//ev.CreateEvent(mod2)
	//ev.CreateEvent(mod3)
	fmt.Println(ev)
}