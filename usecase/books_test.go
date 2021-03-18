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

func TestNewBooks(t *testing.T) {
	type args struct {
		db StoreBook
	}
	tests := []struct {
		name string
		args args
		want *Books
	}{
		{
			name: "Succesful NewBook usecase",
			args: args{
				db: &store.Store{},
			},
			want: &Books{
				store: &store.Store{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBooks(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooks_GetBook(t *testing.T) {
	type args struct {
		key string
	}

	type fields struct {
		store func(m *mocks.MockStoreBook)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Book
		wantErr bool
	}{
		{
			name: "Succesful GetBook",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {
					m.EXPECT().GetBook("Inmensamente pequeño").
						Return(&entity.Book{
							Id:        3,
							Tittle:    "Inmensamente pequeño",
							Author:    "Charles Perrolt",
							Category:  "Infantil",
							Pages:     20,
							Copies:    14,
							Available: true,
						}, nil)
				},
			},
			args: args{
				key: "Inmensamente pequeño",
			},
			want: &entity.Book{
				Id:        3,
				Tittle:    "Inmensamente pequeño",
				Author:    "Charles Perrolt",
				Category:  "Infantil",
				Pages:     20,
				Copies:    14,
				Available: true},
			wantErr: false,
		}, {
			name: "Succesful GetBook error invalid book",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
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

			mockStore := mocks.NewMockStoreBook(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Books{
				store: mockStore,
			}

			got, err := uc.GetBook(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Books.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Books.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooks_GetBooks(t *testing.T) {
	type fields struct {
		store func(m *mocks.MockStoreBook)
	}
	tests := []struct {
		name    string
		fields  fields
		want    []entity.Book
		wantErr bool
	}{
		{
			name: "Succesful GetBooks",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {
					m.EXPECT().GetBooks().
						Return([]entity.Book{
							{
								Id:        2,
								Tittle:    "Como soportar a un hombre",
								Author:    "Lidia Escoto",
								Category:  "Suspenso",
								Pages:     567,
								Copies:    1,
								Available: true,
							},
							{
								Id:        3,
								Tittle:    "Inmensamente pequeño",
								Author:    "Charles Perrolt",
								Category:  "Infantil",
								Pages:     20,
								Copies:    14,
								Available: true,
							},
						}, nil)
				},
			},
			want: []entity.Book{
				{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "Lidia Escoto",
					Category:  "Suspenso",
					Pages:     567,
					Copies:    1,
					Available: true,
				},
				{
					Id:        3,
					Tittle:    "Inmensamente pequeño",
					Author:    "Charles Perrolt",
					Category:  "Infantil",
					Pages:     20,
					Copies:    14,
					Available: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreBook(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Books{
				store: mockStore,
			}

			got, err := uc.GetBooks()
			if (err != nil) != tt.wantErr {
				t.Errorf("Books.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Books.GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooks_CreateBook(t *testing.T) {
	type args struct {
		data entity.Book
	}
	type fields struct {
		store func(m *mocks.MockStoreBook)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Book
		wantErr bool
	}{
		{
			name: "Succesful CreateBook",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {
					m.EXPECT().CreateBook(entity.Book{
						Id:        2,
						Tittle:    "Como soportar a un hombre",
						Author:    "Lidia Escoto",
						Category:  "Suspenso",
						Pages:     567,
						Copies:    1,
						Available: true,
					}).Return(&entity.Book{
						Id:        2,
						Tittle:    "Como soportar a un hombre",
						Author:    "Lidia Escoto",
						Category:  "Suspenso",
						Pages:     567,
						Copies:    1,
						Available: true,
					}, nil)
				},
			},
			args: args{
				data: entity.Book{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "Lidia Escoto",
					Category:  "Suspenso",
					Pages:     567,
					Copies:    1,
					Available: true,
				},
			},
			want: &entity.Book{
				Id:        2,
				Tittle:    "Como soportar a un hombre",
				Author:    "Lidia Escoto",
				Category:  "Suspenso",
				Pages:     567,
				Copies:    1,
				Available: true,
			},
			wantErr: false,
		}, {
			name: "Succesful CreateBook error invalid tittle",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
			},
			args: args{
				data: entity.Book{
					Id:        2,
					Tittle:    "",
					Author:    "Lidia Escoto",
					Category:  "Suspenso",
					Pages:     567,
					Copies:    1,
					Available: true,
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful CreateBook error invalid number of pages",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
			},
			args: args{
				data: entity.Book{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "Lidia Escoto",
					Category:  "Suspenso",
					Pages:     0,
					Copies:    1,
					Available: true,
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful CreateBook error invalid cathegory",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
			},
			args: args{
				data: entity.Book{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "Lidia Escoto",
					Category:  "",
					Pages:     567,
					Copies:    1,
					Available: true,
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful CreateBook error invalid Author",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
			},
			args: args{
				data: entity.Book{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "",
					Category:  "Suspenso",
					Pages:     567,
					Copies:    1,
					Available: true,
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Succesful CreateBook error invalid number of copies",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
			},
			args: args{
				data: entity.Book{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "Lidia Escoto",
					Category:  "Suspenso",
					Pages:     567,
					Copies:    0,
					Available: true,
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

			mockStore := mocks.NewMockStoreBook(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Books{
				store: mockStore,
			}

			got, err := uc.CreateBook(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Books.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Books.CreateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooks_UpdateBook(t *testing.T) {
	type args struct {
		key  string
		data entity.Book
	}
	type fields struct {
		store func(m *mocks.MockStoreBook)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Book
		wantErr bool
	}{
		{
			name: "Succesful UpdateBook",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {
					m.EXPECT().UpdateBook("Como soportar a un hombre", entity.Book{
						Id:        2,
						Tittle:    "Como soportar a un hombre",
						Author:    "Lidia Escoto",
						Category:  "Terror",
						Pages:     507,
						Copies:    1,
						Available: true,
					}).Return(&entity.Book{
						Id:        2,
						Tittle:    "Como soportar a un hombre",
						Author:    "Lidia Escoto",
						Category:  "Terror",
						Pages:     507,
						Copies:    1,
						Available: true,
					}, nil)
				},
			},
			args: args{
				key: "Como soportar a un hombre",
				data: entity.Book{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "Lidia Escoto",
					Category:  "Terror",
					Pages:     507,
					Copies:    1,
					Available: true,
				},
			},
			want: &entity.Book{
				Id:        2,
				Tittle:    "Como soportar a un hombre",
				Author:    "Lidia Escoto",
				Category:  "Terror",
				Pages:     507,
				Copies:    1,
				Available: true,
			},
			wantErr: false,
		}, {
			name: "Succesful UpdateBook error invalid book",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
			},
			args: args{
				key: "",
				data: entity.Book{
					Id:        2,
					Tittle:    "Como soportar a un hombre",
					Author:    "Lidia Escoto",
					Category:  "Suspenso",
					Pages:     567,
					Copies:    1,
					Available: true,
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

			mockStore := mocks.NewMockStoreBook(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Books{
				store: mockStore,
			}

			got, err := uc.UpdateBook(tt.args.key, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Books.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Books.UpdateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooks_DeleteBook(t *testing.T) {
	type args struct {
		key string
	}
	type fields struct {
		store func(m *mocks.MockStoreBook)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Succesful DeleteBook",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {
					m.EXPECT().DeleteBook("Como soportar a un hombre").
						Return(fmt.Errorf("the book doesnt exist"))
				},
			},
			args: args{
				key: "Como soportar a un hombre",
			},
			want:    "the book was erased",
			wantErr: false,
		}, {
			name: "Succesful DeleteBook error invalid book",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {},
			},
			args: args{
				key: "",
			},
			want:    "",
			wantErr: true,
		}, {
			name: "Succesful DeleteBook",
			fields: fields{
				store: func(m *mocks.MockStoreBook) {
					m.EXPECT().DeleteBook("Como soportar a un hombre").
						Return(nil)
				},
			},
			args: args{
				key: "Como soportar a un hombre",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockStore := mocks.NewMockStoreBook(mockCtrl)
			if tt.fields.store != nil {
				tt.fields.store(mockStore)
			}

			uc := &Books{
				store: mockStore,
			}

			got, err := uc.DeleteBook(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Books.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Books.DeleteBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
