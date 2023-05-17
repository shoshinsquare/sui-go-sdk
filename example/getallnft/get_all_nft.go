package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {

	cli := sui.NewSuiClient("https://wallet-rpc.mainnet.sui.io/")
	objects, err := cli.GetAllNFT(context.Background(), "0x7fb8193cdeab49a8411270fc2e152f913962eac89bb1a33a0d2276520dd3ba69")
	if err != nil {
		fmt.Println(err)
	}

	for _, object := range objects {
		fmt.Println(object)
	}
}
