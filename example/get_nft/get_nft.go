package main

import (
	"context"
	"fmt"

	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/shoshinsquare/sui-go-sdk/sui"
)

func main() {

	cli := sui.NewSuiClient("https://wallet-rpc.mainnet.sui.io/")
	object, err := cli.GetObject(context.Background(), models.GetObjectRequest{
		ObjectID: "0x831705e3d61d56c173fb01a81248dc933406fc32ad598f61a1c965386e46a830",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(object)
}
