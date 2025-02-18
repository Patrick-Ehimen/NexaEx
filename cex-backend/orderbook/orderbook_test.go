package orderbook

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 1067)
	buyOrderB := NewOrder(true, 3400)
	buyOrderC := NewOrder(true, 1007)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder((buyOrderC))

	fmt.Println(l)
}

func TestOrderbook(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 1067)
	buyOrderB := NewOrder(true, 1067)

	ob.PlaceOrder(80_000, buyOrderA)
	ob.PlaceOrder(810_000, buyOrderB)

	// for i := 0; i < len(ob.Bids); i++ {
	// 	ob.PlaceOrder(80_000, NewOrder(true, 1067))
	// }

	fmt.Printf("%+v", ob)
	// fmt.Println(ob.Bids[0])
}
