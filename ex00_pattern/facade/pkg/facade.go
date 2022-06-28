package pkg

import (
	"fmt"
)

type Application struct {
	user IUser
	db   DB
}

// NewApplication Суть таковая, что имея несколько сервисов мы объединяем их в один монолит и работает с помощью него
func NewApplication() *Application {
	return &Application{
		user: *NewDefaultUser("TestUser", "TestPass"),
		db:   *NewPostgresConnection("TestParameters"),
	}
}

func (app *Application) Run() {
	err := app.db.Authentification(app.user)
	if err != nil {
		return
	}
	fmt.Println("Application is run")
}
