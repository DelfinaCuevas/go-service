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

func TestNewUsers(t *testing.T) {
	type args struct {
		user UserUseCase
	}
	tests := []struct {
		name string
		args args
		want *Users
	}{
		{
			name: "Succesful NewUsers controller",
			args: args{
				user: &usecase.Users{},
			},
			want: &Users{
				UseCase: &usecase.Users{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsers(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_GetUser(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockUserUseCase)
	}

	e := echo.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful GetUser",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().GetUser("alma.delfina.cuevas@gmail.com").
						Return(&entity.User{
							Id:       1,
							Name:     "Alma Delfina",
							Lastname: "Cuevas Lopez",
							Email:    "alma.delfina.cuevas@gmail.com",
							Address:  "Av.Isla Gomera 3166",
							Phone:    "3317724811",
						}, nil)
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
			name: "Succesful GetUser",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().GetUser("alma.delfina.cuevas@gmail.com").
						Return(&entity.User{}, fmt.Errorf("Something"))
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

			mockUC := mocks.NewMockUserUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Users{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/users/:id")
			eCtx.SetParamNames("id")
			eCtx.SetParamValues("alma.delfina.cuevas@gmail.com")

			if err := c.GetUser(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Users.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsers_GetUsers(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockUserUseCase)
	}

	e := echo.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful GetUsers",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().GetUsers().Return([]entity.User{
						{
							Id:       1,
							Name:     "Alma Delfina",
							Lastname: "Cuevas Lopez",
							Email:    "alma.delfina.cuevas@gmail.com",
							Address:  "Av.Isla Gomera 3166",
							Phone:    "3317724811",
						},
						{
							Id:       1,
							Name:     "Alma Delfina",
							Lastname: "Cuevas Lopez",
							Email:    "alma.delfina.cuevas@gmail.com",
							Address:  "Av.Isla Gomera 3166",
							Phone:    "3317724811",
						},
						{
							Id:       1,
							Name:     "Alma Delfina",
							Lastname: "Cuevas Lopez",
							Email:    "alma.delfina.cuevas@gmail.com",
							Address:  "Av.Isla Gomera 3166",
							Phone:    "3317724811",
						},
						{
							Id:       1,
							Name:     "Alma Delfina",
							Lastname: "Cuevas Lopez",
							Email:    "alma.delfina.cuevas@gmail.com",
							Address:  "Av.Isla Gomera 3166",
							Phone:    "3317724811",
						},
					}, nil)
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
			name: "Succesful GetUsers",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().GetUsers().
						Return([]entity.User{}, fmt.Errorf("Something went wrong"))
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

			mockUC := mocks.NewMockUserUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Users{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/users")

			if err := c.GetUsers(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Users.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsers_CreateUser(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockUserUseCase)
	}

	e := echo.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful CreateUsers",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().CreateUser(entity.User{
						Name:     "Alma Delfina",
						Lastname: "Cuevas Lopez",
						Email:    "alma.delfina.cuevas@gmail.com",
						Address:  "Av.Isla Gomera 3166",
						Phone:    "3317724811",
					}).
						Return(&entity.User{
							Id:       1,
							Name:     "Alma Delfina",
							Lastname: "Cuevas Lopez",
							Email:    "alma.delfina.cuevas@gmail.com",
							Address:  "Av.Isla Gomera 3166",
							Phone:    "3317724811",
						}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(`{`+""+`
						"email" : "alma.delfina.cuevas@gmail.com",`+""+`
						"name" : "Alma Delfina",`+""+`
						"lastname" : "Cuevas Lopez",`+""+`
						"address":"Av.Isla Gomera 3166",`+""+`
						"phone": "3317724811"`+""+`}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Succesful CreateUsers error invalid json",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {},
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
			name: "Succesful CreateUsers error not added",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().CreateUser(entity.User{
						Name:     "Alma Delfina",
						Lastname: "Cuevas Lopez",
						Email:    "alma.delfina.cuevas@gmail.com",
						Address:  "Av.Isla Gomera 3166",
						Phone:    "3317724811",
					}).
						Return(&entity.User{}, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(`{`+""+`
					"email" : "alma.delfina.cuevas@gmail.com",`+""+`
					"name" : "Alma Delfina",`+""+`
					"lastname" : "Cuevas Lopez",`+""+`
					"address":"Av.Isla Gomera 3166",`+""+`
					"phone": "3317724811"`+""+`}`),
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

			mockUC := mocks.NewMockUserUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Users{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/users")

			if err := c.CreateUser(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Users.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsers_UpdateUser(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockUserUseCase)
	}
	e := echo.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful UpdateUsers",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().UpdateUser(entity.User{
						Name:     "Alma Delfina",
						Lastname: "Cuevas Lopez",
						Email:    "alma.delfina.cuevas@gmail.com",
						Address:  "Av.Isla Gomera 3166",
						Phone:    "3317724811",
					}).Return(&entity.User{
						Id:       1,
						Name:     "Alma Delfina",
						Lastname: "Cuevas Lopez",
						Email:    "alma.delfina.cuevas@gmail.com",
						Address:  "Av.Isla Gomera 3166",
						Phone:    "3317724811",
					}, nil)
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPut,
					"/",
					strings.NewReader(`{`+""+`
					"name" : "Alma Delfina",`+""+`
					"lastname" : "Cuevas Lopez",`+""+`
					"address":"Av.Isla Gomera 3166",`+""+`
					"phone": "3317724811"`+""+`}`),
				),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		}, {
			name: "Succesful UpdateUsers error invalid json",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {},
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
			name: "Succesful UpdateUsers",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().UpdateUser(entity.User{
						Name:     "Alma Delfina",
						Lastname: "Cuevas Lopez",
						Email:    "alma.delfina.cuevas@gmail.com",
						Address:  "Av.Isla Gomera 3166",
						Phone:    "3317724811",
					}).Return(&entity.User{}, fmt.Errorf("Something went wrong"))
				},
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPut,
					"/",
					strings.NewReader(`{`+""+`
					"name" : "Alma Delfina",`+""+`
					"lastname" : "Cuevas Lopez",`+""+`
					"address":"Av.Isla Gomera 3166",`+""+`
					"phone": "3317724811"`+""+`}`),
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

			mockUC := mocks.NewMockUserUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Users{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/users/:id")
			eCtx.SetParamNames("id")
			eCtx.SetParamValues("alma.delfina.cuevas@gmail.com")

			if err := c.UpdateUser(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Users.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsers_DeleteUser(t *testing.T) {
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type fields struct {
		usecase func(m *mocks.MockUserUseCase)
	}

	e := echo.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Succesful DeleteUser Controller",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().DeleteUser("alma.delfina.cuevas@gmail.com").
						Return("The user was erased", nil)
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
			name: "Succesful DeleteUser error",
			fields: fields{
				usecase: func(m *mocks.MockUserUseCase) {
					m.EXPECT().DeleteUser("alma.delfina.cuevas@gmail.com").
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

			mockUC := mocks.NewMockUserUseCase(mockCtrl)
			if tt.fields.usecase != nil {
				tt.fields.usecase(mockUC)
			}

			c := &Users{
				UseCase: mockUC,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			eCtx.SetPath("/users/:id")
			eCtx.SetParamNames("id")
			eCtx.SetParamValues("alma.delfina.cuevas@gmail.com")

			if err := c.DeleteUser(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Users.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
