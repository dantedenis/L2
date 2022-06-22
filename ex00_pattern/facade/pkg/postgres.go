package pkg

import (
	"fmt"
	"time"
)

type DB interface {
	ConnectionDB() error
	Authentification(user IUser) error
}

type Postgres struct {
	connectionPar string
}

func NewPostgresConnection(param string) *Postgres {
	res := &Postgres{connectionPar: param}
	res.ConnectionDB()
	return res
}

func (p Postgres) ConnectionDB() error {
	fmt.Printf("Succsess connection Postgres with param: (%s)\n", p.connectionPar)
	return nil
}

func (p Postgres) Authentification(user IUser) error {
	fmt.Printf("Name: <%s> with role: <%s> try to authentification....\n", user.GetLogin(), user.GetUserRole())
	for i:= 0; i < 5; i++{
		fmt.Printf(".")
		time.Sleep(500 * time.Millisecond)			
	}
	fmt.Println()
	fmt.Println("Succsess authentification")
	time.Sleep(1 * time.Second)
	return nil
}