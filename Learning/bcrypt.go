package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123` // Example password
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPassword1 := `password123`
	// loginPassword2 := `password`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1))
	if err != nil {
		fmt.Println("Password Incorrect - Error:", err)
		return
	}
	fmt.Println("You're logged in!")
}
