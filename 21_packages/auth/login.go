package auth

import "fmt"

func LoginWithCredentials(email string, password string) {
	fmt.Println("Logging in with email:", email, "and password:", password)
}