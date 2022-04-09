package batches

type OrderLine struct {
	ID       int
	OrderID  string
	SKU      string
	Quantity int
	BatchID  int
}

func NewOrderLine(orderID, sku string, quantity int) OrderLine {
	return OrderLine{OrderID: orderID, SKU: sku, Quantity: quantity}
}
