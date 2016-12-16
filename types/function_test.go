package types

import (
	"reflect"
	"testing"
)

var (
	foo       = NewFunction(INT32, FLOAT)
	fooReturn = INT32
	fooArgs   = []Type{FLOAT}

	bar       = NewFunction(FLOAT, INT64)
	barReturn = FLOAT
	barArgs   = []Type{INT64}

	baz       = NewFunction(FLOAT, INT64, INT64, INT32)
	bazReturn = FLOAT
	bazArgs   = []Type{INT64, INT64, INT32}
)

func TestNewFunction(t *testing.T) {
	type args struct {
		returnType Type
		argTypes   []Type
	}
	tests := []struct {
		name string
		args args
		want Function
	}{
		{"foo", args{fooReturn, fooArgs}, foo},
		{"bar", args{barReturn, barArgs}, bar},
		{"baz", args{bazReturn, bazArgs}, baz},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFunction(tt.args.returnType, tt.args.argTypes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_String(t *testing.T) {
	type fields struct {
		returnType Type
		argTypes   []Type
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"foo", fields{fooReturn, fooArgs}, "i32 (float)"},
		{"bar", fields{barReturn, barArgs}, "float (i64)"},
		{"baz", fields{bazReturn, bazArgs}, "float (i64, i64, i32)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				returnType: tt.fields.returnType,
				argTypes:   tt.fields.argTypes,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("Function.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Arguments(t *testing.T) {
	type fields struct {
		returnType Type
		argTypes   []Type
	}
	tests := []struct {
		name   string
		fields fields
		want   []Type
	}{
		{"foo", fields{fooReturn, fooArgs}, fooArgs},
		{"bar", fields{barReturn, barArgs}, barArgs},
		{"baz", fields{bazReturn, bazArgs}, bazArgs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				returnType: tt.fields.returnType,
				argTypes:   tt.fields.argTypes,
			}
			if got := f.Arguments(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Function.Arguments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_ReturnType(t *testing.T) {
	type fields struct {
		returnType Type
		argTypes   []Type
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{"foo", fields{fooReturn, fooArgs}, fooReturn},
		{"bar", fields{barReturn, barArgs}, barReturn},
		{"baz", fields{bazReturn, bazArgs}, bazReturn},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				returnType: tt.fields.returnType,
				argTypes:   tt.fields.argTypes,
			}
			if got := f.ReturnType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Function.ReturnType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Equal(t *testing.T) {
	type fields struct {
		returnType Type
		argTypes   []Type
	}
	type args struct {
		n Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"foo == foo", fields{fooReturn, fooArgs}, args{foo}, true},
		{"foo == bar", fields{fooReturn, fooArgs}, args{bar}, false},
		{"foo == baz", fields{fooReturn, fooArgs}, args{baz}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				returnType: tt.fields.returnType,
				argTypes:   tt.fields.argTypes,
			}
			if got := f.Equal(tt.args.n); got != tt.want {
				t.Errorf("Function.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
