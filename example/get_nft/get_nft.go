package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {

	cli := sui.NewSuiClient("https://fullnode.testnet.sui.io/")
	object, err := cli.GetObject(context.Background(), models.GetObjectRequest{
		ObjectID: "0xc74788debd5e9e79e5be1197a63e3e5f0d80af5eea0a3c2d6a4105814c4e7c5a",
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, content := range object.Data.Content.Fields.Attributes.Fields.Contents {
		fmt.Println(content.Fields)
	}
}
