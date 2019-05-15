package controller

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/unrolled/render"
)

func TestNewSomething(t *testing.T) {
	type args struct {
		uc SomethingUseCase
		r  *render.Render
	}
	tests := []struct {
		name string
		args args
		want *Something
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSomething(tt.args.uc, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomething_HandlerSomething(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		op   *Something
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.op.HandlerSomething(tt.args.w, tt.args.r)
		})
	}
}

func TestSomething_HandlerSomethingWithRepo(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Something
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.HandlerSomethingWithRepo(tt.args.w, tt.args.r)
		})
	}
}
