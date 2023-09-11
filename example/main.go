package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/buurzx/orderbook-depth"
	"github.com/shopspring/decimal"
)

func main() {
	book := orderbook.NewOrderBook("BTCUSDT")

	for i := 0; i < 10; i++ {
		// rand.Seed(time.Now().UnixNano())
		side := []orderbook.Side{orderbook.Buy, orderbook.Sell}[rand.Intn(2)]

		book.Append(side, decimal.NewFromInt(rand.Int63n(1000)), decimal.NewFromInt(rand.Int63n(1000)))
	}

	depth, _ := json.Marshal(book.Depth())
	var buf bytes.Buffer
	json.Indent(&buf, depth, "", "  ")
	fmt.Println(buf.String())
}
