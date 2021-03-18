package entity

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type Book struct {
	Id        int    `json:"id"`
	Tittle    string `json:"tittle"`
	Author    string `json:"author"`
	Category  string `json:"category"`
	Pages     int    `json:"pages"`
	Copies    int    `json:"copies"`
	Available bool   `json:"available"`
}

type Loan struct {
	Uuid       string   `json:"id"`
	Loan_Book  []string `json:"loan_book"`
	Loan_User  string   `json:"loan_user"`
	Date_Begin string   `json:"date_begin"`
	Date_End   string   `json:"date_end"`
	State      string   `json:"state"`
	Coments    string   `json:"coments"`
}
