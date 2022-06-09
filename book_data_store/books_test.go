package bookdatastore

import (
	"reflect"
	"testing"
)

func TestPublishBook(t *testing.T) {
	type args struct {
		ba BooksWithAuthor
	}
	tests := []struct {
		name    string
		args    args
		notWant int
		wantErr bool
	}{
		{
			"publish book for jane - author exist",
			args{BooksWithAuthor{Book{"God of small Things", 0, 0, "A book won Booker Price", "", 19.99}, Author{0, "", "jane"}}},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublishBook(tt.args.ba)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublishBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Book.BookID == tt.notWant {
				t.Errorf("PublishBook() = %v, want %v", got, tt.notWant)
			}
		})
	}
}

func TestFindBook(t *testing.T) {
	type args struct {
		queryParam BooksWithAuthor
		argType    string
	}
	tests := []struct {
		name    string
		args    args
		want    BooksWithAuthor
		wantErr bool
	}{

		{
			"Find  book for jane - God of small Things",
			args{BooksWithAuthor{Book{"God of small Things", 0, 0, "", "", 0}, Author{0, "", "jane"}}, "BooksWithAuthor"},
			BooksWithAuthor{Book{"God of small Things", 0, 0, "A book won Booker Price", "", 19.99}, Author{0, "", "jane"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindBook(tt.args.queryParam, tt.args.argType)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindBooks(t *testing.T) {
	type args struct {
		queryParam BooksWithAuthor
	}
	tests := []struct {
		name    string
		args    args
		want    []BooksWithAuthor
		wantErr bool
	}{
		{
			"finds everything",
			args{BooksWithAuthor{}},
			[]BooksWithAuthor{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindBooks(tt.args.queryParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
