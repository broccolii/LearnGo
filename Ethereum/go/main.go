package main

import (
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"fmt"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
)

func main() {
	GetBalance("0d7100ae851cc268b0009c425be758d780ff61a1")
}

func GetBalance(address string) {
	client, err := connectToRPC()

	if err != nil {
		panic(err.Error())
	}

	balance, err := client.BalanceAt(context.TODO(), common.HexToAddress(address), nil)
	fmt.Printf("balance: %v \n",balance)
	if err != nil {
		panic(err.Error())
	} else {
		a := big.NewInt(int64(math.Pow10(18)))
		result := balance.Div(balance, a)
		fmt.Printf("余额: %v \n",result)
	}
}

func connectToRPC() (*ethclient.Client, error) {
	client, err := rpc.Dial("http://127.0.0.1:8545")

	if err != nil {
		fmt.Printf("rpc dial error = %v", err)
		return nil, err
	}

	conn := ethclient.NewClient(client)

	return conn, nil
}