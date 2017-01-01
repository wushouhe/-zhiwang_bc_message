package main

import (
	"github.com/ethereum/go-ethereum/rpc"
	"zhiwang_bc_message/geth/json"
	"zhiwang_bc_message/geth/subscribe"
	"fmt"
)

func main() {

	client, _ := rpc.Dial("http://139.196.178.168:8545")
	blockChan := make(chan *json.JsonHeader,100)
	fmt.Println("正在监听new Block.........")
	subscribe.ListenNewBlock(client, blockChan)
	for {
		select {
		case block := <-blockChan:

			fmt.Printf(`
			coinbase %s
			number %s
			difficulty %s
			gaslimit %s
			gasused %s
			nonce %#v
			parenthash %s
			txhash %s
			receipthash %s
			bloom %s
			extra %s
			MixDigest %s
			root %s
			time %s
			Transactions %v `,
				block.Coinbase.Hex(),
				block.Number.ToInt(),
				block.Difficulty.ToInt(),
				block.GasLimit.ToInt(),
				block.GasUsed.ToInt(),
				block.Nonce.Uint64(),
				block.ParentHash.Hex(),
				block.TxHash.Hex(),
				block.ReceiptHash.Hex(),
				block.Bloom.Big(),
				block.Extra.String(),
				block.MixDigest.Hex(),
				block.Root.Hex(),
				block.Time.ToInt(),
				block.Transactions)

			//fmt.Println(block.Transactions)


		}
	}

}





