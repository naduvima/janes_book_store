package jwt_secure_access

import "testing"

func Test_decodeRequest(t *testing.T) {
	type args struct {
		tokenString string
		user        string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"successful token validation",
			args{
				tokenString: "eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjoidGVzdCIsInVzZXIiOiJqYW5lIiwicGFzc3dvcmQiOiJwYXNzd29yZCIsImlhdCI6MTY1NDM2MTExMH0.rLAolKGgJnBeag3LJUeAabGAuBmCDJtSLjjz4NLBsIE",
				user:        "jane",
			},
			true,
		},
		{
			"un successful token validation",
			args{
				tokenString: "eyJfdsfdsfdfd.eyJkYXRhIjoidGVzdCIsInVzZXIiOiJqYW5lIiwicGFzc3dvcmQiOiJwYXNzd29yZCJ9.bPrauWUbd_gErQnXfF0i6NjJVU89FnOU-OZ-57niuCo",
				user:        "jane",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodetRequest(tt.args.tokenString, tt.args.user); got != tt.want {
				t.Errorf("decodePostRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
