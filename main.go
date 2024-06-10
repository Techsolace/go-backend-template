package main

import (
	"context"
	"encoding/json"
	"fmt"
	"techsolace/prisma/db"
)

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}

func Run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}
	ctx := context.Background()

	addUser, err := client.User.CreateOne(
		db.User.Mail.Set("samarth1234"),
		db.User.Name.Set("Samarth"),
	).Exec(ctx)
	if err != nil {
		return err
	}
	result, _ := json.MarshalIndent(addUser, "", " ")
	fmt.Printf("added user %v", result)
	return nil
}
