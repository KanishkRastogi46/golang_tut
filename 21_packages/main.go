package main

import (
	"fmt"

	"github.com/KanishkRastogi46/golang_tut/21_packages/auth"
	"github.com/KanishkRastogi46/golang_tut/21_packages/users"
	"github.com/fatih/color"
)

func main() {
	// creating a user
	var user users.Users = users.Users{
		Id: "1",
		Email: "kanishk@gmail.com",
		Password: "secret",
	}

	// logging in the user
	auth.LoginWithCredentials(user.Email, user.Password)

	// getting the session
	session, err := auth.GetSession();
	if (err != nil) {
		panic(err)
	}
	fmt.Println("Session:", session)

	// using color package
	color.Yellow(user.Id)
	color.Red(user.Email)
	color.Magenta(user.Password)
}