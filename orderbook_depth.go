package orderbook

// Depth returns the depth.
func (ob *OrderBook) Depth() *Depth {
	defer ob.RUnlock()
	ob.RLock()

	asks := make([]*PriceLevel, 0)
	level := ob.asks.MinPriceQueue()

	for level != nil {
		asks = append(asks, NewPriceLevel(level.price, level.amount))
		level = ob.asks.GreaterThan(level.price)
	}

	bids := make([]*PriceLevel, 0)
	level = ob.bids.MaxPriceQueue()

	for level != nil {
		bids = append(bids, NewPriceLevel(level.price, level.amount))
		level = ob.bids.LessThan(level.price)
	}

	return &Depth{bids, asks}
}
