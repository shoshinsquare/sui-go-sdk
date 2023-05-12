package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {

	cli := sui.NewSuiClient("https://wallet-rpc.mainnet.sui.io/")
	object, err := cli.GetLatestCheckpointSequenceNumber(context.Background(), models.GetLatestCheckpointSequenceNumberRequest{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(object)
}
