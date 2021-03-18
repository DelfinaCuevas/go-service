package usecase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eiizu/go-service/entity"
	"github.com/eiizu/go-service/store"
	"github.com/eiizu/go-service/usecase/mocks"
	"github.com/golang/mock/gomock"
)

func TestNewUsers(t *testing.T) {
	type args struct {
		db StoreUser
	}

	tests := []struct {
		name string
		args args
		want *Users
	}{
		{
			name: "Succesful NewUser usecase",
			args: args{
				db: &store.Store{},
			},
			want: &Users{
				store: &store.Store{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsers(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_GetUser(t *testing.T) {
	type args struct {
		key string
	}
	type fields struct {
		store func(m *mocks.MockStoreUser)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "Succesful GetUser",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {
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
				key: "alma.delfina.cuevas@gmail.com",
			},
			want: &entity.User{
				Id:       1,
				Name:     "Alma Delfina",
				Lastname: "Cuevas Lopez",
				Email:    "alma.delfina.cuevas@gmail.com",
				Address:  "Av.Isla Gomera 3166",
				Phone:    "3317724811",
			},
			wantErr: false,
		}, {
			name: "Succesful GetUser",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {},
			},
			args: args{
				key: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreUser(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Users{
				store: mockStore,
			}

			got, err := uc.GetUser(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_GetUsers(t *testing.T) {

	type fields struct {
		store func(m *mocks.MockStoreUser)
	}

	tests := []struct {
		name    string
		fields  fields
		want    []entity.User
		wantErr bool
	}{
		{
			name: "Succesful GetUsers",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {
					m.EXPECT().GetUsers().
						Return([]entity.User{
							{
								Id:       1,
								Name:     "Alma Delfina",
								Lastname: "Cuevas Lopez",
								Email:    "alma.delfina.cuevas@gmail.com",
								Address:  "Av.Isla Gomera 3166",
								Phone:    "3317724811",
							}, {
								Id:       3,
								Name:     "Miguel",
								Lastname: "Padilla",
								Email:    "miguel@gmail.com",
								Address:  "Av.Isla Gomera 3166",
								Phone:    "3317724811",
							},
						}, nil)
				},
			},
			want: []entity.User{
				{
					Id:       1,
					Name:     "Alma Delfina",
					Lastname: "Cuevas Lopez",
					Email:    "alma.delfina.cuevas@gmail.com",
					Address:  "Av.Isla Gomera 3166",
					Phone:    "3317724811",
				}, {
					Id:       3,
					Name:     "Miguel",
					Lastname: "Padilla",
					Email:    "miguel@gmail.com",
					Address:  "Av.Isla Gomera 3166",
					Phone:    "3317724811",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreUser(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Users{
				store: mockStore,
			}

			got, err := uc.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_CreateUser(t *testing.T) {
	type args struct {
		data entity.User
	}

	type fields struct {
		store func(m *mocks.MockStoreUser)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "Succesful GetUser",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {
					m.EXPECT().CreateUser(entity.User{
						Id:       1,
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
				data: entity.User{
					Id:       1,
					Name:     "Alma Delfina",
					Lastname: "Cuevas Lopez",
					Email:    "alma.delfina.cuevas@gmail.com",
					Address:  "Av.Isla Gomera 3166",
					Phone:    "3317724811",
				},
			},
			want: &entity.User{
				Id:       1,
				Name:     "Alma Delfina",
				Lastname: "Cuevas Lopez",
				Email:    "alma.delfina.cuevas@gmail.com",
				Address:  "Av.Isla Gomera 3166",
				Phone:    "3317724811",
			},
			wantErr: false,
		}, {
			name: "Succesful GetUser error name",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {},
			},
			args: args{
				data: entity.User{
					Id:       1,
					Name:     "",
					Lastname: "Cuevas Lopez",
					Email:    "alma.delfina.cuevas@gmail.com",
					Address:  "Av.Isla Gomera 3166",
					Phone:    "3317724811",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful GetUser error email",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {},
			},
			args: args{
				data: entity.User{
					Id:       1,
					Name:     "Alma Delfina",
					Lastname: "Cuevas Lopez",
					Email:    "",
					Address:  "Av.Isla Gomera 3166",
					Phone:    "3317724811",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful GetUser errir address",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {},
			},
			args: args{
				data: entity.User{
					Id:       1,
					Name:     "Alma Delfina",
					Lastname: "Cuevas Lopez",
					Email:    "alma.delfina.cuevas@gmail.com",
					Address:  "",
					Phone:    "3317724811",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful GetUser error phone",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {},
			},
			args: args{
				data: entity.User{
					Id:       1,
					Name:     "Alma Delfina",
					Lastname: "Cuevas Lopez",
					Email:    "alma.delfina.cuevas@gmail.com",
					Address:  "Av.Isla Gomera 3166",
					Phone:    "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreUser(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Users{
				store: mockStore,
			}
			got, err := uc.CreateUser(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_UpdateUser(t *testing.T) {
	type args struct {
		data entity.User
	}
	type fields struct {
		store func(m *mocks.MockStoreUser)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "Succesful UpdatetUser",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {
					m.EXPECT().UpdateUser(entity.User{
						Id:       1,
						Name:     "Alma Delfina",
						Lastname: "Cuevas Lopez",
						Email:    "alma@gmail.com",
						Address:  "Av.Isla Gomera",
						Phone:    "3317724811",
					}).
						Return(&entity.User{
							Id:       1,
							Name:     "Alma Delfina",
							Lastname: "Cuevas Lopez",
							Email:    "alma@gmail.com",
							Address:  "Av.Isla Gomera",
							Phone:    "3317724811",
						}, nil)
				},
			},
			args: args{
				data: entity.User{
					Id:       1,
					Name:     "Alma Delfina",
					Lastname: "Cuevas Lopez",
					Email:    "alma@gmail.com",
					Address:  "Av.Isla Gomera",
					Phone:    "3317724811",
				},
			},
			want: &entity.User{
				Id:       1,
				Name:     "Alma Delfina",
				Lastname: "Cuevas Lopez",
				Email:    "alma@gmail.com",
				Address:  "Av.Isla Gomera",
				Phone:    "3317724811",
			},
			wantErr: false,
		}, {
			name: "Succesful UpdatetUser error email",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {},
			},
			args: args{
				data: entity.User{
					Id:       1,
					Name:     "Alma Delfina",
					Lastname: "Cuevas Lopez",
					Email:    "",
					Address:  "Av.Isla Gomera",
					Phone:    "3317724811",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreUser(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Users{
				store: mockStore,
			}
			got, err := uc.UpdateUser(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_DeleteUser(t *testing.T) {
	type args struct {
		key string
	}
	type fields struct {
		store func(m *mocks.MockStoreUser)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Succesful DeleteUser",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {
					m.EXPECT().DeleteUser("alma.delfina.cuevas@gmail.com").
						Return(&entity.User{}, nil)
				},
			},
			args: args{
				key: "alma.delfina.cuevas@gmail.com",
			},
			want:    "The user was erased",
			wantErr: false,
		}, {
			name: "Succesful DeleteUser",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {},
			},
			args: args{
				key: "",
			},
			want:    "",
			wantErr: true,
		}, {
			name: "Succesful DeleteUser",
			fields: fields{
				store: func(m *mocks.MockStoreUser) {
					m.EXPECT().DeleteUser("alma.delfina.cuevas@gmail.com").
						Return(nil, fmt.Errorf("Something went wrong"))

				},
			},
			args: args{
				key: "alma.delfina.cuevas@gmail.com",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreUser(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Users{
				store: mockStore,
			}

			got, err := uc.DeleteUser(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
