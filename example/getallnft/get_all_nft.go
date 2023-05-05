package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {

	cli := sui.NewSuiClient("https://wallet-rpc.mainnet.sui.io/")
	objects, err := cli.GetAllNFT(context.Background(), "0x1131dea66ccff11bc001162ec0a7fb823574e71654fb0843ca4cd6f2982b16fb")
	if err != nil {
		fmt.Println(err)
	}

	for _, object := range objects {
		fmt.Println(object)
	}
}
