package pkg

import (
	"fmt"
)

type Application struct {
	user IUser
	db DB
}

func NewApplication() *Application{
	return &Application{
		user: *NewDefaultUser("TestUser", "TestPass"),
		db: *NewPostgresConnection("TestParammeters"),
	}
}

func (app *Application) Run() {
	app.db.Authentification(app.user)
	fmt.Println("Application is run")
}