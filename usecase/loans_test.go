package usecase

import (
	"reflect"
	"testing"

	"github.com/eiizu/go-service/entity"
	"github.com/eiizu/go-service/store"
	"github.com/eiizu/go-service/usecase/mocks"
	"github.com/golang/mock/gomock"
)

func TestNewLoans(t *testing.T) {
	type args struct {
		db StoreLoan
	}
	tests := []struct {
		name string
		args args
		want *Loans
	}{
		{
			name: "Succesful NewLoan UseCase",
			args: args{
				db: &store.Store{},
			},
			want: &Loans{
				store: &store.Store{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoans(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoans_GetLoans(t *testing.T) {
	type args struct {
		parameters map[string]string
	}
	type fields struct {
		store func(m *mocks.MockStoreLoan)
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantLoan map[string]entity.Loan
		wantErr  bool
	}{
		{
			name: "Succesful GetLoans",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {
					m.EXPECT().GetLoans().
						Return(map[string]entity.Loan{
							"23fc2204-13b5-4f78-a211-20214477b456": {
								Uuid: "23fc2204-13b5-4f78-a211-20214477b456",
								Loan_Book: []string{
									"Inmensamente pequeño",
									"Como soportar a un hombre",
								},
								Loan_User:  "alma@gmail.com",
								Date_Begin: "2021-02-10 11:37:46.6019703-06:00",
								Date_End:   "00/00/0000",
								State:      "Loan",
								Coments:    "En curso",
							},
							"3925204d-6bb8-4e93-861d-ca4e5265e195": {
								Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
								Loan_Book: []string{
									"Inmensamente pequeño",
									"Como soportar a un hombre",
									"Holaa",
								},
								Loan_User:  "alma.delfina.cuevas@gmail.com",
								Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
								Date_End:   "00/00/0000",
								State:      "On loan",
								Coments:    "En curso",
							},
						}, nil)
				},
			},
			args: args{
				parameters: map[string]string{
					"book": "",
					"user": "",
					"uuid": "",
				},
			},
			wantLoan: map[string]entity.Loan{
				"23fc2204-13b5-4f78-a211-20214477b456": {
					Uuid: "23fc2204-13b5-4f78-a211-20214477b456",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
					},
					Loan_User:  "alma@gmail.com",
					Date_Begin: "2021-02-10 11:37:46.6019703-06:00",
					Date_End:   "00/00/0000",
					State:      "Loan",
					Coments:    "En curso",
				},
				"3925204d-6bb8-4e93-861d-ca4e5265e195": {
					Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
						"Holaa",
					},
					Loan_User:  "alma.delfina.cuevas@gmail.com",
					Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
					Date_End:   "00/00/0000",
					State:      "On loan",
					Coments:    "En curso",
				},
			},
			wantErr: false,
		}, {
			name: "Succesful GetLoan",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {
					m.EXPECT().GetLoan(map[string]string{
						"book": "Holaa",
						"user": "",
						"uuid": "",
					}).Return(map[string]entity.Loan{
						"3925204d-6bb8-4e93-861d-ca4e5265e195": {
							Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
							Loan_Book: []string{
								"Inmensamente pequeño",
								"Como soportar a un hombre",
								"Holaa",
							},
							Loan_User:  "alma.delfina.cuevas@gmail.com",
							Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
							Date_End:   "00/00/0000",
							State:      "On loan",
							Coments:    "En curso",
						},
					}, nil)
				},
			},
			args: args{
				parameters: map[string]string{
					"book": "Holaa",
					"user": "",
					"uuid": "",
				},
			},
			wantLoan: map[string]entity.Loan{
				"3925204d-6bb8-4e93-861d-ca4e5265e195": {
					Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
						"Holaa",
					},
					Loan_User:  "alma.delfina.cuevas@gmail.com",
					Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
					Date_End:   "00/00/0000",
					State:      "On loan",
					Coments:    "En curso",
				},
			},
			wantErr: false,
		}, {
			name: "Succesful GetLoan",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {
					m.EXPECT().GetLoan(map[string]string{
						"book": "",
						"user": "alma.delfina.cuevas@gmail.com",
						"uuid": "",
					}).Return(map[string]entity.Loan{
						"3925204d-6bb8-4e93-861d-ca4e5265e195": {
							Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
							Loan_Book: []string{
								"Inmensamente pequeño",
								"Como soportar a un hombre",
								"Holaa",
							},
							Loan_User:  "alma.delfina.cuevas@gmail.com",
							Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
							Date_End:   "00/00/0000",
							State:      "On loan",
							Coments:    "En curso",
						},
					}, nil)
				},
			},
			args: args{
				parameters: map[string]string{
					"book": "",
					"user": "alma.delfina.cuevas@gmail.com",
					"uuid": "",
				},
			},
			wantLoan: map[string]entity.Loan{
				"3925204d-6bb8-4e93-861d-ca4e5265e195": {
					Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
						"Holaa",
					},
					Loan_User:  "alma.delfina.cuevas@gmail.com",
					Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
					Date_End:   "00/00/0000",
					State:      "On loan",
					Coments:    "En curso",
				},
			},
			wantErr: false,
		}, {
			name: "Succesful GetLoan",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {
					m.EXPECT().GetLoan(map[string]string{
						"book": "",
						"user": "",
						"uuid": "3925204d-6bb8-4e93-861d-ca4e5265e195",
					}).Return(map[string]entity.Loan{
						"3925204d-6bb8-4e93-861d-ca4e5265e195": {
							Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
							Loan_Book: []string{
								"Inmensamente pequeño",
								"Como soportar a un hombre",
								"Holaa",
							},
							Loan_User:  "alma.delfina.cuevas@gmail.com",
							Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
							Date_End:   "00/00/0000",
							State:      "On loan",
							Coments:    "En curso",
						},
					}, nil)
				},
			},
			args: args{
				parameters: map[string]string{
					"book": "",
					"user": "",
					"uuid": "3925204d-6bb8-4e93-861d-ca4e5265e195",
				},
			},
			wantLoan: map[string]entity.Loan{
				"3925204d-6bb8-4e93-861d-ca4e5265e195": {
					Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
						"Holaa",
					},
					Loan_User:  "alma.delfina.cuevas@gmail.com",
					Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
					Date_End:   "00/00/0000",
					State:      "On loan",
					Coments:    "En curso",
				},
			},
			wantErr: false,
		}, {
			name: "Succesful GetLoan_",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {
					m.EXPECT().GetLoan_(map[string]string{
						"book": "Inmensamente pequeño",
						"user": "alma.delfina.cuevas@gmail.com",
						"uuid": "",
					}).Return(map[string]entity.Loan{
						"3925204d-6bb8-4e93-861d-ca4e5265e195": {
							Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
							Loan_Book: []string{
								"Inmensamente pequeño",
								"Como soportar a un hombre",
								"Holaa",
							},
							Loan_User:  "alma.delfina.cuevas@gmail.com",
							Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
							Date_End:   "00/00/0000",
							State:      "On loan",
							Coments:    "En curso",
						},
					}, nil)
				},
			},
			args: args{
				parameters: map[string]string{
					"book": "Inmensamente pequeño",
					"user": "alma.delfina.cuevas@gmail.com",
					"uuid": "",
				},
			},
			wantLoan: map[string]entity.Loan{
				"3925204d-6bb8-4e93-861d-ca4e5265e195": {
					Uuid: "3925204d-6bb8-4e93-861d-ca4e5265e195",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
						"Holaa",
					},
					Loan_User:  "alma.delfina.cuevas@gmail.com",
					Date_Begin: "2021-02-23 10:48:44.4527775-06:00",
					Date_End:   "00/00/0000",
					State:      "On loan",
					Coments:    "En curso",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreLoan(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Loans{
				store: mockStore,
			}
			got, err := uc.GetLoans(tt.args.parameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("Loans.CreateLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantLoan) {
				t.Errorf("Loans.CreateLoan() = %v, want %v", got, tt.wantLoan)
			}
		})
	}
}

func TestLoans_CreateLoan(t *testing.T) {
	type args struct {
		data entity.Loan
	}
	type fields struct {
		store func(m *mocks.MockStoreLoan)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Loan
		wantErr bool
	}{

		{
			name: "Succesful CreateLoan",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {
					m.EXPECT().CreateLoan(entity.Loan{
						Uuid: "",
						Loan_Book: []string{
							"Inmensamente pequeño",
							"Como soportar a un hombre",
						},
						Loan_User:  "alma@gmail.com",
						Date_Begin: "",
						Date_End:   "",
						State:      "",
						Coments:    "En curso",
					}).Return(&entity.Loan{
						Uuid: "23fc2204-13b5-4f78-a211-20214477b456",
						Loan_Book: []string{
							"Inmensamente pequeño",
							"Como soportar a un hombre",
						},
						Loan_User:  "alma@gmail.com",
						Date_Begin: "2021-02-10 11:37:46.6019703-06:00",
						Date_End:   "00/00/0000",
						State:      "On loan",
						Coments:    "En curso"}, nil)
				},
			},
			args: args{
				data: entity.Loan{
					Uuid: "",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
					},
					Loan_User:  "alma@gmail.com",
					Date_Begin: "",
					Date_End:   "",
					State:      "",
					Coments:    "En curso",
				},
			},
			want: &entity.Loan{
				Uuid: "23fc2204-13b5-4f78-a211-20214477b456",
				Loan_Book: []string{
					"Inmensamente pequeño",
					"Como soportar a un hombre",
				},
				Loan_User:  "alma@gmail.com",
				Date_Begin: "2021-02-10 11:37:46.6019703-06:00",
				Date_End:   "00/00/0000",
				State:      "On loan",
				Coments:    "En curso"},
			wantErr: false,
		}, {
			name: "Succesful CreateLoan error Invalid User",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {},
			},
			args: args{
				data: entity.Loan{
					Uuid: "",
					Loan_Book: []string{
						"Inmensamente pequeño",
						"Como soportar a un hombre",
					},
					Loan_User:  "",
					Date_Begin: "",
					Date_End:   "",
					State:      "",
					Coments:    "En curso",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful CreateLoan error Invalid Loans Books",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {},
			},
			args: args{
				data: entity.Loan{
					Uuid:       "",
					Loan_Book:  []string{},
					Loan_User:  "alma@gmail.com",
					Date_Begin: "",
					Date_End:   "",
					State:      "",
					Coments:    "En curso",
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

			mockStore := mocks.NewMockStoreLoan(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Loans{
				store: mockStore,
			}
			got, err := uc.CreateLoan(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Loans.CreateLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Loans.CreateLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoans_UpdateLoan(t *testing.T) {
	type args struct {
		data entity.Loan
	}
	type fields struct {
		store func(m *mocks.MockStoreLoan)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Loan
		wantErr bool
	}{
		{
			name: "Succesful GetLoans",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {
					m.EXPECT().UpdateLoan(entity.Loan{
						Uuid:       "23fc2204-13b5-4f78-a211-20214477b456",
						Loan_Book:  []string{},
						Loan_User:  "",
						Date_Begin: "",
						Date_End:   "",
						State:      "Devuelto",
						Coments:    "Sin comentarios",
					}).
						Return(&entity.Loan{
							Uuid: "23fc2204-13b5-4f78-a211-20214477b456",
							Loan_Book: []string{
								"Inmensamente pequeño",
								"Como soportar a un hombre",
							},
							Loan_User:  "alma@gmail.com",
							Date_Begin: "2021-02-10 11:37:46.6019703-06:00",
							Date_End:   "2021-03-24 18:37:46.6019703-06:00",
							State:      "Devuelto",
							Coments:    "Sin comentarios"}, nil)
				},
			},
			args: args{
				data: entity.Loan{
					Uuid:       "23fc2204-13b5-4f78-a211-20214477b456",
					Loan_Book:  []string{},
					Loan_User:  "",
					Date_Begin: "",
					Date_End:   "",
					State:      "Devuelto",
					Coments:    "Sin comentarios",
				},
			},
			want: &entity.Loan{
				Uuid: "23fc2204-13b5-4f78-a211-20214477b456",
				Loan_Book: []string{
					"Inmensamente pequeño",
					"Como soportar a un hombre",
				},
				Loan_User:  "alma@gmail.com",
				Date_Begin: "2021-02-10 11:37:46.6019703-06:00",
				Date_End:   "2021-03-24 18:37:46.6019703-06:00",
				State:      "Devuelto",
				Coments:    "Sin comentarios"},
			wantErr: false,
		}, {
			name: "Succesful GetLoans error Id",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {},
			},
			args: args{
				data: entity.Loan{
					Uuid:       "",
					Loan_Book:  []string{},
					Loan_User:  "",
					Date_Begin: "",
					Date_End:   "",
					State:      "Devuelto",
					Coments:    "Sin comentarios",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful GetLoans error Comments",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {},
			},
			args: args{
				data: entity.Loan{
					Uuid:       "23fc2204-13b5-4f78-a211-20214477b456",
					Loan_Book:  []string{},
					Loan_User:  "",
					Date_Begin: "",
					Date_End:   "",
					State:      "Devuelto",
					Coments:    "",
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful GetLoans error State",
			fields: fields{
				store: func(m *mocks.MockStoreLoan) {},
			},
			args: args{
				data: entity.Loan{
					Uuid:       "23fc2204-13b5-4f78-a211-20214477b456",
					Loan_Book:  []string{},
					Loan_User:  "",
					Date_Begin: "",
					Date_End:   "",
					State:      "",
					Coments:    "Sin comentarios",
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

			mockStore := mocks.NewMockStoreLoan(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Loans{
				store: mockStore,
			}
			got, err := uc.UpdateLoan(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Loans.UpdateLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Loans.UpdateLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}
