package main

import (
	"context"
	"fmt"
	"github.com/AppleGrey/Kitex-Demo/kitex_gen/user"
	"github.com/AppleGrey/Kitex-Demo/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

func createUser(ctx context.Context, cli userservice.Client) error {
	req := &user.CreateUserRequest{
		Name:      "test",
		Gender:    user.Gender_Male,
		Age:       18,
		Introduce: "test",
	}

	resp, err := cli.CreateUser(ctx, req)
	if err != nil {
		fmt.Printf("create failed: %s\n", err.Error())
	}
	fmt.Println(resp)
	return err
}

func main() {
	var opts []client.Option
	// specify the address of the server
	opts = append(opts, client.WithHostPorts("localhost:8888"))
	// construct a client
	cli := userservice.MustNewClient("applegrey.demo", opts...)
	ctx := context.Background()
	// initialize one request
	err := createUser(ctx, cli)
	if err != nil {
		return
	}
}
