package main

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v11"
)

func main() {
    client := gocloak.NewClient("http://localhost:8080")
	fmt.Println(client)
	ctx := context.Background()
	fmt.Println(ctx)
	jwt, err := client.LoginAdmin(ctx, "test", "password", "master")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(jwt)
}