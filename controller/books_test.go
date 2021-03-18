package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/eiizu/go-service/controller/mocks"
	"github.com/eiizu/go-service/entity"
	"github.com/eiizu/go-service/usecase"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestNewBooks(t *testing.T) {
	type args struct {
		book BooksUseCase
	}
	tests := []struct {
		name string
		args args
		want *Books
	}{
		{
			name: "Succesful NewBooks controller",
			args: args{
				book: &usecase.Books{},
			},
			want: &Books{
				UseCase: &usecase.Books{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBooks(tt.args.book); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooks_GetBook(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockBooksUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful GetBook",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().GetBook("Inmensamente pequeño").
						Return(&entity.Book{
							Id:        2,
							Tittle:    "Inmensamente pequeño",
							Author:    "Lidia Escoto",
							Category:  "Suspenso",
							Pages:     567,
							Copies:    1,
							Available: true}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodGet,
					"/",
					strings.NewReader(""),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Succesful GetBook",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().GetBook("Inmensamente pequeño").
						Return(&entity.Book{}, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodGet,
					"/",
					strings.NewReader(""),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUC := mocks.NewMockBooksUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Books{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/books/:id")
			eCtx.SetParamNames("id")
			eCtx.SetParamValues("Inmensamente pequeño")

			if err := c.GetBook(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Books.GetBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooks_GetBooks(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockBooksUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful GetBooks",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().GetBooks().
						Return([]entity.Book{
							{
								Id:        2,
								Tittle:    "Inmensamente pequeño",
								Author:    "Lidia Escoto",
								Category:  "Suspenso",
								Pages:     567,
								Copies:    1,
								Available: true}}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodGet,
					"/",
					strings.NewReader(""),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Succesful GetBooks",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().GetBooks().
						Return(nil, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodGet,
					"/",
					strings.NewReader(""),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUC := mocks.NewMockBooksUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Books{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/books")
			if err := c.GetBooks(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Books.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooks_CreateBook(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockBooksUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success CreateBook",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().CreateBook(entity.Book{
						Tittle:    "LA La song",
						Author:    "Charles Perrolt",
						Category:  "Infantil",
						Pages:     230,
						Copies:    4,
						Available: true,
					}).Return(&entity.Book{
						Id:        2,
						Tittle:    "LA La song",
						Author:    "Charles Perrolt",
						Category:  "Infantil",
						Pages:     230,
						Copies:    4,
						Available: true,
					}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(`{`+""+`
						"tittle": "LA La song",`+""+`
						"author": "Charles Perrolt",`+""+`
						"category": "Infantil",`+""+`
						"pages": 230,`+""+`
						"copies": 4,`+""+`
						"available": true`+""+`
					}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Success CreateBook",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(``),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Success CreateBook",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().CreateBook(entity.Book{
						Tittle:    "LA La song",
						Author:    "Charles Perrolt",
						Category:  "Infantil",
						Pages:     230,
						Copies:    4,
						Available: true,
					}).Return(&entity.Book{}, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(`{`+""+`
						"tittle": "LA La song",`+""+`
						"author": "Charles Perrolt",`+""+`
						"category": "Infantil",`+""+`
						"pages": 230,`+""+`
						"copies": 4,`+""+`
						"available": true`+""+`
					}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUC := mocks.NewMockBooksUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Books{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/books")

			if err := c.CreateBook(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Books.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooks_UpdateBook(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockBooksUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success Update",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().UpdateBook("LA La song", entity.Book{
						Tittle:    "La La song",
						Author:    "Charles Perrolt",
						Category:  "Infantil",
						Pages:     230,
						Copies:    4,
						Available: true,
					}).Return(&entity.Book{
						Id:        2,
						Tittle:    "La La song",
						Author:    "Charles Perrolt",
						Category:  "Infantil",
						Pages:     230,
						Copies:    4,
						Available: true,
					}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPut,
					"/",
					strings.NewReader(`{`+""+`
						"tittle": "La La song",`+""+`
						"author": "Charles Perrolt",`+""+`
						"category": "Infantil",`+""+`
						"pages": 230,`+""+`
						"copies": 4,`+""+`
						"available": true`+""+`
					}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Success Update",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPut,
					"/",
					strings.NewReader(``),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Success Update",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().UpdateBook("LA La song", entity.Book{
						Tittle:    "La La song",
						Author:    "Charles Perrolt",
						Category:  "Infantil",
						Pages:     230,
						Copies:    4,
						Available: true,
					}).Return(nil, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPut,
					"/",
					strings.NewReader(`{`+""+`
						"tittle": "La La song",`+""+`
						"author": "Charles Perrolt",`+""+`
						"category": "Infantil",`+""+`
						"pages": 230,`+""+`
						"copies": 4,`+""+`
						"available": true`+""+`
					}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUC := mocks.NewMockBooksUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Books{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/books/:id")
			eCtx.SetParamNames("id")
			eCtx.SetParamValues("LA La song")

			if err := c.UpdateBook(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Books.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooks_DeleteBook(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockBooksUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success Delete",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().DeleteBook("LA La song").
						Return("The book was erased", nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodDelete,
					"/",
					strings.NewReader(``),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Success Delete",
			fields: fields{
				usecase: func(m *mocks.MockBooksUseCase) {
					m.EXPECT().DeleteBook("LA La song").
						Return("", fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodDelete,
					"/",
					strings.NewReader(``),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUC := mocks.NewMockBooksUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Books{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/books/:id")
			eCtx.SetParamNames("id")
			eCtx.SetParamValues("LA La song")
			if err := c.DeleteBook(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Books.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
