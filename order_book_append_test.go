package orderbook_test

import (
	"encoding/json"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/buurzx/orderbook-depth"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAppendOrderBids(t *testing.T) {
	type input struct {
		OrderID string
		side    orderbook.Side
		amount  decimal.Decimal
		price   decimal.Decimal
	}

	type snapshot struct {
		Book *orderbook.OrderBook
		Err  string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "happy path - 1",
			input: input{
				OrderID: "100",
				side:    orderbook.Buy,
				amount:  decimal.NewFromInt(9),
				price:   decimal.NewFromInt(100),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [
						{
							"id": "100",
							"side": "buy",
							"amount": "5",
							"price": "100"
						}
					],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			err = book.Append(tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book: &book,
				Err:  errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestAppendOrderBidsAndRemove(t *testing.T) {
	type input struct {
		OrderID string
		side    orderbook.Side
		amount  decimal.Decimal
		price   decimal.Decimal
	}

	type snapshot struct {
		Book *orderbook.OrderBook
		Err  string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "happy path - 1",
			input: input{
				OrderID: "100",
				side:    orderbook.Buy,
				amount:  decimal.NewFromInt(0),
				price:   decimal.NewFromInt(100),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [
						{
							"id": "100",
							"side": "buy",
							"amount": "5",
							"price": "100"
						}
					],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			err = book.Append(tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book: &book,
				Err:  errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestAppendOrderEmptyBids(t *testing.T) {
	type input struct {
		OrderID string
		side    orderbook.Side
		amount  decimal.Decimal
		price   decimal.Decimal
	}

	type snapshot struct {
		Book *orderbook.OrderBook
		Err  string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "happy path - 1",
			input: input{
				OrderID: "100",
				side:    orderbook.Buy,
				amount:  decimal.NewFromInt(9),
				price:   decimal.NewFromInt(100),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			err = book.Append(tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book: &book,
				Err:  errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestAppendOrderEmptyAsks(t *testing.T) {
	type input struct {
		side   orderbook.Side
		amount decimal.Decimal
		price  decimal.Decimal
	}

	type snapshot struct {
		Book *orderbook.OrderBook
		Err  string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "happy path - 1",
			input: input{
				side:   orderbook.Sell,
				amount: decimal.NewFromInt(5),
				price:  decimal.NewFromInt(550),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			err = book.Append(tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book: &book,
				Err:  errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestProcessPostOnlyOrderValidations(t *testing.T) {
	type input struct {
		side   orderbook.Side
		amount decimal.Decimal
		price  decimal.Decimal
	}

	type snapshot struct {
		Book *orderbook.OrderBook
		Err  string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "invalid order id",
			input: input{
				side:   orderbook.Sell,
				amount: decimal.NewFromInt(5),
				price:  decimal.NewFromInt(550),
			},
		},
		{
			name: "order already exists",
			input: input{
				side:   orderbook.Sell,
				amount: decimal.NewFromInt(5),
				price:  decimal.NewFromInt(550),
			},
		},
		{
			name: "invalid trader id",
			input: input{
				side:   orderbook.Sell,
				amount: decimal.NewFromInt(5),
				price:  decimal.NewFromInt(550),
			},
		},
		{
			name: "invalid amount",
			input: input{
				side:   orderbook.Sell,
				amount: decimal.NewFromInt(0),
				price:  decimal.NewFromInt(550),
			},
		},
		{
			name: "invalid price",
			input: input{
				side:   orderbook.Sell,
				amount: decimal.NewFromInt(5),
				price:  decimal.NewFromInt(0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [
						{
							"side": "buy",
							"amount": "5",
							"price": "500"
						},
						{
							"side": "buy",
							"amount": "1",
							"price": "400"
						},
						{
							"side": "buy",
							"amount": "0.5",
							"price": "300"
						}
					],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			err = book.Append(tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book: &book,
				Err:  errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}
