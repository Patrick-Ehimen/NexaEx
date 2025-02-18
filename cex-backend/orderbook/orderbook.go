package orderbook

import (
	"fmt"
	"sort"
	"time"
)

type Match struct {
    /// @notice The order that is being asked in the match
    Ask        *Order

    /// @notice The order that is being bid in the match
    Bid        *Order

    /// @notice The size filled in the match
    SizeFilled float64

    /// @notice The price at which the match occurred
    Price      float64
}

type Order struct {
	Size      float64
	Bid       bool
	Limit     *Limit
	Timestamp int64
}

type Orders []*Order

func (o Orders) Len() int           { return len(o) }
func (o Orders) Swap(i, j int)      {o[i], o[j] = o[j], o[i] }
func (o Orders) Less(i, j int) bool { return o[i].Timestamp < o[j].Timestamp }

// The `Limit` type represents a limit order with price, orders, and total volume information.
// @property {float64} Price - The `Price` property in the `Limit` struct represents the price at which
// the limit order is set. It is a floating-point number that specifies the desired price for buying or
// selling an asset.
// @property {Orders} Orders - Orders is a data structure that represents a collection of orders
// related to the limit. It could be a list or array of order objects containing information such as
// order ID, quantity, timestamp, and other relevant details.
// @property {float64} TotalVolume - TotalVolume represents the total volume of the limit order. It
// indicates the total quantity of the asset that the user wants to buy or sell at the specified price
// in the limit order.
type Limit struct {
	Price       float64
	Orders      Orders
	TotalVolume float64
}

type Limits []*Limit

type ByBestAsk struct {Limits}

func (a ByBestAsk) Len() int           { return len(a.Limits) }
func (a ByBestAsk) Swap(i, j int)      {a.Limits[i], a.Limits[j] = a.Limits[j], a.Limits[i] }
func (a ByBestAsk) Less(i, j int) bool { return a.Limits[i].Price < a.Limits[j].Price }

// The code defines a type named ByBestBid that includes a field of type Limits.
// @property {Limits}  - It looks like you have defined a new type `ByBestBid` with an embedded field
// `Limits`. This means that the `ByBestBid` type will inherit the fields and methods of the `Limits`
// type.
type ByBestBid struct {Limits}

func (b ByBestBid) Len() int           { return len(b.Limits) }
func (b ByBestBid) Swap(i, j int)      {b.Limits[i], b.Limits[j] = b.Limits[j], b.Limits[i] }
func (b ByBestBid) Less(i, j int) bool { return b.Limits[i].Price > b.Limits[j].Price }

// The NewOrder function creates a new order with the specified bid type, size, and timestamp.
func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}

// The `func (o *Order) String() string` method is a method defined on the `Order` struct in Go. It
// overrides the default `String()` method for the `Order` struct, allowing you to customize how an
// `Order` object is represented as a string when using `fmt.Println` or other formatting functions.
func (o *Order) String() string {
	return fmt.Sprintf("[size: %.2f]", o.Size)
}

// The NewLimit function creates a new Limit object with a specified price and an empty list of Orders.
func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}


// The `AddOrder` method of the `Limit` struct is used to add an order to the list of orders within
// that limit. Here's what each line of the method does:
func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

// The `DeleteOrder` method of the `Limit` struct is used to remove a specific order from the list of orders within that limit. Here's what each line of the method does:
func (l *Limit) DeleteOrder(o *Order) {
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
		}
	}
	o.Limit = nil
	l.TotalVolume -= o.Size

	sort.Sort(Orders(l.Orders))
}

// The Orderbook type in Go represents a collection of Ask and Bid limits with corresponding price
// levels.
// @property {[]*Limit} Asks - Asks is a slice of pointers to Limit structs, representing the list of
// asking prices in the order book.
// @property {[]*Limit} Bids - The `Bids` property in the `Orderbook` struct is a slice of pointers to
// `Limit` structs. It likely represents the list of buy orders in the order book, with each `Limit`
// struct containing information about a specific buy order.
// @property AskLimits - The `AskLimits` property in the `Orderbook` struct is a map that associates a
// price (float64) with a `Limit` struct representing an ask order at that price. This map allows for
// quick lookup of ask orders based on their price in the order book.
// @property BidLimits - The `BidLimits` property in the `Orderbook` struct is a map that associates a
// float64 key with a `Limit` value. This map is used to store bid limits in the order book, where the
// key represents the price level of the bid limit and the value is the corresponding `
type Orderbook struct {
	Asks []*Limit
	Bids []*Limit

	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
}

// The NewOrderbook function initializes a new Orderbook struct with empty slices for Asks and Bids and empty maps for AskLimits and BidLimits.
func NewOrderbook() *Orderbook {
	return &Orderbook{
		Asks:      []*Limit{},
		Bids:      []*Limit{},
		AskLimits: make(map[float64]*Limit),
		BidLimits: make(map[float64]*Limit),
	}
}

// The `PlaceOrder` method in the `Orderbook` struct is used to place an order in the order book at a specified price. Here's what the method does:
func (ob *Orderbook) PlaceOrder(price float64, o *Order) []Match {
	if o.Size > 0.0 {
		ob.add(price, o)
	}
	return []Match{}
}

// This `add` method in the `Orderbook` struct is responsible for adding an order to the order book at a specified price. Here's a breakdown of what the method does:
func (ob *Orderbook) add(price float64, o *Order) {
	var limit *Limit

	if o.Bid {
		limit = ob.BidLimits[price]
	} else {
		limit = ob.AskLimits[price]
	}

	if limit == nil {
		limit = NewLimit(price)
		limit.AddOrder(o)

		if o.Bid {
			ob.Bids = append(ob.Bids, limit)
			ob.BidLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AskLimits[price] = limit
		}
	}
}
