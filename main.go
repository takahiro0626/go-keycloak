package main

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v11"
)

func main() {
    client := gocloak.NewClient("http://localhost:8080")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "kc", "kc", "master")
	if err != nil {
		fmt.Println(err.Error())
	}

	user := gocloak.User{
		FirstName: gocloak.StringP("Bob"),
		LastName:  gocloak.StringP("Uncle"),
		Email:     gocloak.StringP("something@really.wrong"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP("CoolGuy"),
	   }
	fmt.Println(user.FirstName)

	_, err = client.CreateUser(ctx, token.AccessToken, "test", user)
    if err != nil {
        fmt.Println("Failed to create user")
    }
}