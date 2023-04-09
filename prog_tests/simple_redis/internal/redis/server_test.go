package redis

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Foo struct {
	Name string
	Age int
}


func Test_server_Set(t *testing.T) {
	srv := New()

	type args struct {
		key     string
		untyped interface{}
	}
	tests := []struct {
		name    string
		server *server
		args    args
		wantErr bool
	}{
		{
			name: "should be pass",
			server: srv,
			args: args{
				key:     "name",
				untyped: "carlos",
			},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := srv.Set(tt.args.key, tt.args.untyped); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_server_Get(t *testing.T) {
	srv := New()

	srv.Set("name", "carlos")
	srv.Set("custom", Foo{Name: "carlos", Age: 30})

	tests := []struct {
		name    string
		key string
		server *server
		want    interface{}
		wantErr bool
	}{
		{
			name: "should fail, key does not exist",
			key: "does_not_exist",
			wantErr: true,
		},
		{
			name: "should pass, key exists",
			key: "name",
			want: "carlos",
		},
		{
			name: "should pass, with custom types",
			key: "custom",
			want: Foo{
				Name: "carlos",
				Age:  30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.Get(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	srv := New()

	srv.Set("name", "carlos")
	srv.Set("custom", Foo{Name: "carlos", Age: 30})

	tests := []struct {
		name string
		key string
		want interface{}
		wantErr bool
	}{
		{
			name: "should fail, does not exists",
			key: "does_not_exist",
			wantErr: true,
		},
		{
			name: "should pass",
			key: "name",
			want: "carlos",
		},
		{
			name: "should pass, custom types",
			key: "custom",
			want: Foo{
				Name: "carlos",
				Age:  30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.Delete(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() err = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}