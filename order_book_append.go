package orderbook

import (
	"github.com/shopspring/decimal"
)

// Append processes a post only order.
func (ob *OrderBook) Append(side Side, amount, price decimal.Decimal) error {
	defer func() {
		ob.Unlock()
	}()

	ob.Lock()

	if amount.LessThanOrEqual(decimal.Zero) {
		order := NewOrder(price.String(), side, amount, price)
		listElement, ok := ob.orders[order.id]
		if !ok {
			return nil
		}

		if side == Buy {
			ob.bids.Remove(listElement)
		} else {
			ob.asks.Remove(listElement)
		}

		return nil
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return ErrInvalidPrice
	}

	order := NewOrder(price.String(), side, amount, price)

	listElement, ok := ob.orders[order.id]

	if side == Buy {
		if ok {
			ob.bids.UpdateAmount(listElement, amount)
		} else {
			ob.orders[order.id] = ob.bids.Append(order)
		}

		return nil
	}

	if ok {
		ob.asks.UpdateAmount(listElement, amount)
	} else {
		ob.orders[order.id] = ob.asks.Append(order)
	}

	return nil
}
