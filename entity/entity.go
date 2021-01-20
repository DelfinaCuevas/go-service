package entity

type User struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Book struct {
	Name     string `json:"name"`
	Pages    int    `json:"pages"`
	Category string `json:"category"`
}
