package types

import (
	"reflect"
	"testing"
)

var (
	void = NewVoid()
)

func TestNewVoid(t *testing.T) {
	tests := []struct {
		name string
		want Void
	}{
		{"void", void},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVoid(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVoid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoid_String(t *testing.T) {
	tests := []struct {
		name string
		t    Void
		want string
	}{
		{"void", void, "void"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Void{}
			if got := v.String(); got != tt.want {
				t.Errorf("Void.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoid_Equal(t *testing.T) {
	type args struct {
		a Type
	}
	tests := []struct {
		name string
		t    Void
		args args
		want bool
	}{
		{"void == void", void, args{void}, true},
		{"void == i32", void, args{INT32}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Void{}
			if got := v.Equal(tt.args.a); got != tt.want {
				t.Errorf("Void.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
