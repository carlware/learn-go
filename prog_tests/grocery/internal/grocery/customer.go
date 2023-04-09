package grocery

import "errors"

const (
	CustomerTypeA = "A"
	CustomerTypeB = "B"
)

var (
	ErrUnknownType = errors.New("unknown type")
)

type Customer struct {
	Type        string
	ArrivalTime int
	Items       int
}

func CreateCustomer(kind string, time, items int) (*Customer, error) {
	if kind != CustomerTypeB && kind != CustomerTypeA {
		return nil, ErrUnknownType
	}
	return &Customer{
		Type:        kind,
		ArrivalTime: time,
		Items:       items,
	}, nil
}

type sortCustomersByTimeItemsAndType []*Customer

func (c sortCustomersByTimeItemsAndType) Len() int      { return len(c) }
func (c sortCustomersByTimeItemsAndType) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c sortCustomersByTimeItemsAndType) Less(i, j int) bool {
	if c[i].ArrivalTime != c[j].ArrivalTime {
		return c[i].ArrivalTime < c[j].ArrivalTime
	}
	if c[i].Items != c[i].Items {
		return c[i].Items < c[j].Items
	}
	return c[i].Type < c[j].Type
}
