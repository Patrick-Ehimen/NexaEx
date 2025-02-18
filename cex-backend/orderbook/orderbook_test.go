package orderbook

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T){
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
	
}