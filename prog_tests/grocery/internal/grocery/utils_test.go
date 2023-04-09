package grocery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalculateTotalTime(t *testing.T) {

	t.Run("should get 7", func(t *testing.T) {
		c1, _ := CreateCustomer(CustomerTypeA, 1, 2)
		c2, _ := CreateCustomer(CustomerTypeA, 2, 1)
		cs := []*Customer{c1, c2}

		got := CalculateTotalTime(1, cs)
		assert.Equal(t, 7, got)
	})

	t.Run("should get 13", func(t *testing.T) {
		c1, _ := CreateCustomer(CustomerTypeA, 1, 5)
		c2, _ := CreateCustomer(CustomerTypeB, 2, 1)
		c3, _ := CreateCustomer(CustomerTypeA, 3, 5)
		c4, _ := CreateCustomer(CustomerTypeB, 5, 3)
		c5, _ := CreateCustomer(CustomerTypeA, 8, 2)
		cs := []*Customer{c1, c2, c3, c4, c5}

		got := CalculateTotalTime(2, cs)
		assert.Equal(t, 13, got)
	})

	t.Run("should get 6", func(t *testing.T) {
		c1, _ := CreateCustomer(CustomerTypeA, 1, 2)
		c2, _ := CreateCustomer(CustomerTypeA, 1, 2)
		c3, _ := CreateCustomer(CustomerTypeA, 2, 1)
		c4, _ := CreateCustomer(CustomerTypeB, 3, 2)
		cs := []*Customer{c1, c2, c3, c4}

		got := CalculateTotalTime(2, cs)
		assert.Equal(t, 6, got)
	})

	t.Run("should get 9", func(t *testing.T) {
		c1, _ := CreateCustomer(CustomerTypeA, 1, 2)
		c2, _ := CreateCustomer(CustomerTypeA, 1, 3)
		c3, _ := CreateCustomer(CustomerTypeA, 2, 1)
		c4, _ := CreateCustomer(CustomerTypeA, 2, 1)
		cs := []*Customer{c1, c2, c3, c4}

		got := CalculateTotalTime(2, cs)
		assert.Equal(t, 9, got)
	})

	t.Run("should get 11", func(t *testing.T) {
		c1, _ := CreateCustomer(CustomerTypeA, 1, 3)
		c2, _ := CreateCustomer(CustomerTypeA, 1, 5)
		c3, _ := CreateCustomer(CustomerTypeA, 3, 1)
		c4, _ := CreateCustomer(CustomerTypeB, 4, 1)
		c5, _ := CreateCustomer(CustomerTypeA, 4, 1)
		cs := []*Customer{c1, c2, c3, c4, c5}

		got := CalculateTotalTime(2, cs)
		assert.Equal(t, 11, got)
	})
}
