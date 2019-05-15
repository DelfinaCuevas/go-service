package usecase

import (
	"reflect"
	"testing"
)

func TestNewSomething(t *testing.T) {
	type args struct {
		repo Repo
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
			if got := NewSomething(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomething_DoSomething(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		uc      *Something
		args    args
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DoSomething(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Something.DoSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Something.DoSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomething_DoSomethingWithRepo(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		uc      *Something
		args    args
		want    map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DoSomethingWithRepo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Something.DoSomethingWithRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Something.DoSomethingWithRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
