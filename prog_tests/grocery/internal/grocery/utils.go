package grocery

import (
	"sort"
)

const (
	CheckoutTimeForTraineeCashier     = 2
	CheckoutTimeForExperiencedCashier = 1
)

func CreateRegisters(total int) []*Register {
	rs := []*Register{}
	idx := 1
	for i := 0; i < total-1; i++ {
		rs = append(rs, CreateRegister(CheckoutTimeForExperiencedCashier, idx))
		idx++
	}
	rs = append(rs, CreateRegister(CheckoutTimeForTraineeCashier, idx))
	return rs
}

func CalculateTotalTime(totalCashiers int, cs []*Customer) int {
	customers := sortCustomersByTimeItemsAndType(cs)
	sort.Sort(customers)

	registers := CreateRegisters(totalCashiers)

	for _, customer := range customers {
		var r *Register

		// if some customers finish purchasing; remove them
		// TODO: improve this
		for _, re := range registers {
			re.RemoveCustomersIfCheckoutFinish(customer.ArrivalTime)
		}

		if customer.Type == CustomerTypeA {
			r = getRegisterForCustomerA(registers)
		} else if customer.Type == CustomerTypeB {
			r = getRegisterForCustomerB(registers)
		}

		// if for some reason we get nil pointers don't panic
		if customer == nil || r == nil {
			continue
		}
		r.AddCustomer(customer)
	}
	return getTotalTime(registers)
}

func getRegisterForCustomerA(rs []*Register) *Register {
	sortedRegister := sortRegistersByNumberOfCustomersInLine(rs)
	sort.Sort(sortedRegister)
	return sortedRegister[0]
}

func getRegisterForCustomerB(rs []*Register) *Register {
	sortedRegister := sortRegistersByLastCustomerItems(rs)
	sort.Sort(sortedRegister)
	return sortedRegister[0]
}

func getTotalTime(rs []*Register) int {
	sortedRegister := sortRegistersByEndTime(rs)
	sort.Sort(sortedRegister)
	return sortedRegister[len(sortedRegister)-1].endTime
}
