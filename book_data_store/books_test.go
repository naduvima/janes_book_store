package bookdatastore

import "testing"

func TestPublishBook(t *testing.T) {
	type args struct {
		ba BooksWithAuthor
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"publish book for jane - author exist",
			args{BooksWithAuthor{Book{"God of small Things", 0, 0, "A book won Booker Price", "", 19.99}, Author{0, "", "jane"}}},
			"",
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
			if got != tt.want {
				t.Errorf("PublishBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
