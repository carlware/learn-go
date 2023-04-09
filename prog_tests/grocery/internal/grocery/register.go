package grocery

type Register struct {
	ID              int
	// Describe total of customers waiting in line
	customersInLine int
	// Describe total of items for last Customer
	lastCustomerItems int
	// Array of times when cashier finish checkout for customer
	finishTimes []int
	// total time for cashier finish dispatching all customers
	endTime int
	// time to checkout an item
	n int
}

func CreateRegister(timeToCheckout int, id int) *Register {
	return &Register{
		ID: id,
		n:  timeToCheckout,
	}
}

func (r *Register) RemoveCustomersIfCheckoutFinish(t int) {
	idx := 0
	for _, ft := range r.finishTimes {
		if ft > t {
			break
		}
		r.customersInLine -= 1
		idx += 1
	}
	r.finishTimes = r.finishTimes[idx:]
}

func (r *Register) AddCustomer(c *Customer) {
	endTime := c.ArrivalTime
	if r.endTime > 0 {
		endTime = r.endTime
	}

	timeToFinish := (r.n * c.Items) + endTime
	r.customersInLine += 1
	r.lastCustomerItems = c.Items
	r.endTime = timeToFinish
	r.finishTimes = append(r.finishTimes, timeToFinish)
}

// useful for getting the total time
type sortRegistersByEndTime []*Register

func (r sortRegistersByEndTime) Len() int           { return len(r) }
func (r sortRegistersByEndTime) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r sortRegistersByEndTime) Less(i, j int) bool { return r[i].endTime < r[j].endTime }

type sortRegistersByNumberOfCustomersInLine []*Register

func (r sortRegistersByNumberOfCustomersInLine) Len() int      { return len(r) }
func (r sortRegistersByNumberOfCustomersInLine) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r sortRegistersByNumberOfCustomersInLine) Less(i, j int) bool {
	if r[i].customersInLine != r[j].customersInLine {
		return r[i].customersInLine < r[j].customersInLine
	}
	return r[i].ID < r[j].ID
}

type sortRegistersByLastCustomerItems []*Register

func (r sortRegistersByLastCustomerItems) Len() int      { return len(r) }
func (r sortRegistersByLastCustomerItems) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r sortRegistersByLastCustomerItems) Less(i, j int) bool {
	return r[i].lastCustomerItems < r[j].lastCustomerItems
}
