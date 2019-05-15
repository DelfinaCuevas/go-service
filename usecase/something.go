package usecase

import (
	"strings"
)

// Repo -
type Repo interface {
	GetAccounts(string) (map[string]string, error)
}

// Something -
type Something struct{
	repo Repo
}

// NewSomething -
func NewSomething(repo Repo) *Something {
	return &Something{
		repo: repo,
	}
}

// DoSomething -
func (uc *Something) DoSomething(data string) (map[string]int, error) {
	resp := map[string]int{}

	arr := strings.Split(data, " ")
	for _, value := range arr {
		resp[value]++
	}

	return resp, nil
}

// DoSomethingWithRepo -
func (uc *Something) DoSomethingWithRepo(data string) (map[string]string, error) {
	resp := map[string]string{}

	resp, err:= uc.repo.GetAccounts("1")
	if err != nil {
		return resp, nil
	}

	return resp, nil
}
