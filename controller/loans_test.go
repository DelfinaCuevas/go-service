package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/eiizu/go-service/controller/mocks"
	"github.com/eiizu/go-service/entity"
	"github.com/eiizu/go-service/usecase"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestNewLoans(t *testing.T) {
	type args struct {
		loan LoansUseCase
	}
	tests := []struct {
		name string
		args args
		want *Loans
	}{
		{
			name: "Succesful NewLoans controller",
			args: args{
				loan: &usecase.Loans{},
			},
			want: &Loans{
				UseCaseLoan: &usecase.Loans{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoans(tt.args.loan); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoans_GetLoans(t *testing.T) {
	type args struct {
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockLoansUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesfull GetLoans",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {
					m.EXPECT().GetLoans(map[string]string{
						"book": "",
						"user": "",
						"uuid": "8d5a6717-77c2-4bc6-bd1e-d400883d3bd0",
					}).Return(map[string]entity.Loan{
						"8d5a6717-77c2-4bc6-bd1e-d400883d3bd0": {
							Uuid: "8d5a6717-77c2-4bc6-bd1e-d400883d3bd0",
							Loan_Book: []string{
								"Inmensamente pequeño",
								"Como soportar a un hombre",
							},
							Loan_User:  "alma@gmail.com",
							Date_Begin: "2021-02-10 11:36:45.3868275-06:00",
							Date_End:   "00/00/0000",
							State:      "Loan",
							Coments:    "En curso",
						},
					}, nil)
				},
			},
			args: args{
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Succesfull GetLoans",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {
					m.EXPECT().GetLoans(map[string]string{
						"book": "",
						"user": "",
						"uuid": "8d5a6717-77c2-4bc6-bd1e-d400883d3bd0",
					}).Return(nil, fmt.Errorf("Something went Wrong"))
				},
			},
			args: args{
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUC := mocks.NewMockLoansUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Loans{
				UseCaseLoan: mockUC,
			}

			q := make(url.Values)
			q.Set("uuid", "8d5a6717-77c2-4bc6-bd1e-d400883d3bd0")
			req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

			eCtx := e.NewContext(req, tt.args.rec)

			if err := c.GetLoans(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Loans.GetLoans() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoans_CreateLoan(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockLoansUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful CreateLoan",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {
					m.EXPECT().CreateLoan(entity.Loan{
						Loan_Book: []string{
							"Inmensamente pequeño",
							"Como soportar a un hombre",
						},
						Loan_User: "alma@gmail.com",
						Coments:   "En curso",
					}).Return(&entity.Loan{
						Uuid: "8d5a6717-77c2-4bc6-bd1e-d400883d3bd0",
						Loan_Book: []string{
							"Inmensamente pequeño",
							"Como soportar a un hombre",
						},
						Loan_User:  "alma@gmail.com",
						Date_Begin: "2021-02-10 11:36:45.3868275-06:00",
						Date_End:   "00/00/0000",
						State:      "Loan",
						Coments:    "En curso",
					}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(`{`+""+`
						"loan_book": [`+""+`
							"Inmensamente pequeño",`+""+`
							"Como soportar a un hombre"`+""+`
						],`+""+`
						"loan_user": "alma@gmail.com",`+""+`
						"coments": "En curso"`+""+`}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Succesful CreateLoan",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {},
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
			name: "Succesful CreateLoan",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {
					m.EXPECT().CreateLoan(entity.Loan{
						Loan_Book: []string{
							"Inmensamente pequeño",
							"Como soportar a un hombre",
						},
						Loan_User: "alma@gmail.com",
						Coments:   "En curso",
					}).Return(nil, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(`{`+""+`
						"loan_book": [`+""+`
							"Inmensamente pequeño",`+""+`
							"Como soportar a un hombre"`+""+`
						],`+""+`
						"loan_user": "alma@gmail.com",`+""+`
						"coments": "En curso"`+""+`}`),
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

			mockUC := mocks.NewMockLoansUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Loans{
				UseCaseLoan: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/loans")

			if err := c.CreateLoan(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Loans.CreateLoan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoans_UpdateLoan(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockLoansUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful UpdateLoan",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {
					m.EXPECT().UpdateLoan(entity.Loan{
						Uuid:    "23fc2204-13b5-4f78-a211-20214477b456",
						State:   "Devuelto",
						Coments: "Entrego en tiempo",
					}).Return(&entity.Loan{
						Uuid: "23fc2204-13b5-4f78-a211-20214477b456",
						Loan_Book: []string{
							"Inmensamente pequeño",
							"Como soportar a un hombre",
						},
						Loan_User:  "alma@gmail.com",
						Date_Begin: "2021-02-10 11:36:45.3868275-06:00",
						Date_End:   "2021-02-10 11:36:45.3868275-06:00",
						State:      "Devuelto",
						Coments:    "Entrego en tiempo",
					}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPut,
					"/",
					strings.NewReader(` {`+""+`
						"state": "Devuelto",`+""+`
						"coments": "Entrego en tiempo"`+""+`
					}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Succesful UpdateLoan",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {},
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
			name: "Succesful UpdateLoan",
			fields: fields{
				usecase: func(m *mocks.MockLoansUseCase) {
					m.EXPECT().UpdateLoan(entity.Loan{
						Uuid:    "23fc2204-13b5-4f78-a211-20214477b456",
						State:   "Devuelto",
						Coments: "Entrego en tiempo",
					}).Return(nil, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPut,
					"/",
					strings.NewReader(` {`+""+`
						"state": "Devuelto",`+""+`
						"coments": "Entrego en tiempo"`+""+`
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

			mockUC := mocks.NewMockLoansUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Loans{
				UseCaseLoan: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/loans/:id")
			eCtx.SetParamNames("id")
			eCtx.SetParamValues("23fc2204-13b5-4f78-a211-20214477b456")
			if err := c.UpdateLoan(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Loans.UpdateLoan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
