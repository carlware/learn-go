package grocery

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	type args struct {
		kind  string
		time  int
		items int
	}
	tests := []struct {
		name    string
		args    args
		want    *Customer
		wantErr bool
	}{
		{
			name: "should raise error",
			args: args{
				kind:  "X",
				time:  1,
				items: 1,
			},
			wantErr: true,
		},
		{
			name: "should pass, type B",
			args: args{
				kind:  "B",
				time:  1,
				items: 1,
			},
		},
		{
			name: "should pass, type A",
			args: args{
				kind:  "B",
				time:  1,
				items: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := CreateCustomer(tt.args.kind, tt.args.time, tt.args.items); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_Sort(t *testing.T) {
	tests := []struct {
		name string
		req  []*Customer
		want []*Customer
	}{
		{
			name: "sort by time only",
			req: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 2,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 3,
					Items:       2,
				},
			},
			want: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 2,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 3,
					Items:       2,
				},
			},
		},
		{
			name: "sort by time only",
			req: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 3,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 2,
					Items:       2,
				},
			},
			want: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 2,
					Items:       2,
				},
				{
					Type:        "A",
					ArrivalTime: 3,
					Items:       2,
				},
			},
		},
		{
			name: "sort by time and items",
			req: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       3,
				},
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       4,
				},
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       1,
				},
			},
			want: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       1,
				},
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       3,
				},
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       4,
				},
			},
		},
		{
			name: "sort by time, items and type",
			req: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       3,
				},
				{
					Type:        "B",
					ArrivalTime: 1,
					Items:       3,
				},
				{
					Type:        "B",
					ArrivalTime: 1,
					Items:       3,
				},
			},
			want: []*Customer{
				{
					Type:        "A",
					ArrivalTime: 1,
					Items:       3,
				},
				{
					Type:        "B",
					ArrivalTime: 1,
					Items:       3,
				},
				{
					Type:        "B",
					ArrivalTime: 1,
					Items:       3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(sortCustomersByTimeItemsAndType(tt.req))
			assert.ElementsMatch(t, tt.req, tt.want)
		})
	}
}
