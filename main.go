package main

import (
	"fmt"

	"github.com/rafaph/notte-auth/domain/entities"
)

func main() {
	user, err := entities.NewUser("batata")
	fmt.Println(user)
	fmt.Println(err)
}
