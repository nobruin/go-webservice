package main

import (
	"fmt"

	"github.com/nobruin/api/firstwebservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Benjamin",
		LastName:  "Marins",
	}

	fmt.Println(u)
}
