package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {

	cli := sui.NewSuiClient("https://wallet-rpc.mainnet.sui.io/")
	objects, err := cli.GetAllNFT(context.Background(), "0x95014a81857fb00e6a711b3d24c570d34b57ed4e3d0a206dabbcf23d5aacdea9")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, object := range objects {

		fmt.Println(object)
	}
}
