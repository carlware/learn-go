package grocery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister_Clear(t *testing.T) {
	tests := []struct {
		name            string
		register        *Register
		time            int
		wantCustomers   int
		wantFinishTimes []int
	}{
		{
			name: "nothing happen if register does not have any customers",
			register: &Register{
				customersInLine:   0,
				lastCustomerItems: 0,
				finishTimes:       []int{},
				endTime:           0,
				n:                 1,
			},
			time:            3,
			wantCustomers:   0,
			wantFinishTimes: []int{},
		},
		{
			name: "customers should not be decreased if time is less than first finish checkout time for customer",
			register: &Register{
				customersInLine:   3,
				lastCustomerItems: 5,
				finishTimes:       []int{3, 5, 7},
				endTime:           7,
				n:                 1,
			},
			time:            2,
			wantCustomers:   3,
			wantFinishTimes: []int{3, 5, 7},
		},
		{
			name: "one customer should be removed",
			register: &Register{
				customersInLine:   3,
				lastCustomerItems: 5,
				finishTimes:       []int{3, 5, 7},
				endTime:           7,
				n:                 1,
			},
			time:            3,
			wantCustomers:   2,
			wantFinishTimes: []int{5, 7},
		},
		{
			name: "all customers should be removed",
			register: &Register{
				customersInLine:   3,
				lastCustomerItems: 5,
				finishTimes:       []int{3, 5, 7},
				endTime:           7,
				n:                 1,
			},
			time:            7,
			wantCustomers:   0,
			wantFinishTimes: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.register.RemoveCustomersIfCheckoutFinish(tt.time)
			assert.Equal(t, tt.wantCustomers, tt.register.customersInLine)
			assert.ElementsMatch(t, tt.wantFinishTimes, tt.register.finishTimes)
		})
	}
}

func TestRegister_AddCustomer(t *testing.T) {
	t.Run("add one customer to register", func(t *testing.T) {
		r := CreateRegister(1, 1)
		c1, _ := CreateCustomer(CustomerTypeA, 1, 3)
		r.AddCustomer(c1)

		assert.Equal(t, 4, r.endTime)
		assert.Equal(t, 3, r.lastCustomerItems)
		assert.Equal(t, 1, r.customersInLine)
		assert.ElementsMatch(t, []int{4}, r.finishTimes)
	})

	t.Run("add one customer to trainee register", func(t *testing.T) {
		r := CreateRegister(2, 1)
		c1, _ := CreateCustomer(CustomerTypeA, 1, 3)
		r.AddCustomer(c1)

		assert.Equal(t, 7, r.endTime)
		assert.Equal(t, 3, r.lastCustomerItems)
		assert.Equal(t, 1, r.customersInLine)
		assert.ElementsMatch(t, []int{7}, r.finishTimes)
	})

	t.Run("add two customer to register", func(t *testing.T) {
		r := CreateRegister(1, 1)
		c1, _ := CreateCustomer(CustomerTypeA, 1, 3)
		c2, _ := CreateCustomer(CustomerTypeA, 1, 4)
		r.AddCustomer(c1)
		r.AddCustomer(c2)

		assert.Equal(t, 8, r.endTime)
		assert.Equal(t, 4, r.lastCustomerItems)
		assert.Equal(t, 2, r.customersInLine)
		assert.ElementsMatch(t, []int{4, 8}, r.finishTimes)
	})
}
