package usecase

import (
	"fmt"
)

// Status -
type Status struct {
	AppName string
}

// NewStatus -
func NewStatus(AppName string) *Status {
	return &Status{
		AppName: AppName,
	}
}

// Statusz -
func (*Status) Statusz() (string, error) {
	return "Libreria!", nil
}

// Healthz -
func (uc *Status) Healthz() (string, error) {
	resp := fmt.Sprintf("App Name: %s", uc.AppName)
	return resp, nil
}
