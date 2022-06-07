package bookdatastore

import (
	"reflect"
	"testing"
)

func TestFindAuthor(t *testing.T) {

	tests := []struct {
		name    string
		author  Author
		want    Author
		wantErr bool
	}{
		{
			"find existing user",
			Author{0,"","jane"},
			Author{1,"eyJkYXRhIjoidGVzdCJ9","jane"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindAuthor(tt.author)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAuthor() got = %v, want %v", got, tt.want)
			}
		})
	}
}