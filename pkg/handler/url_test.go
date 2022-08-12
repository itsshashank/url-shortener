package handler

import (
	"errors"
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {
	type result struct {
		url string
		err error
	}
	tests := []struct {
		name string
		arg  string
		want result
	}{
		{
			name: "YouTube",
			arg:  "https://www.youtube.com/",
			want: result{url: "aHR0cHM6Ly93d3cueW91dHViZS5jb20v", err: nil},
		},
		{
			name: "NormalString",
			arg:  "youtube.com",
			want: result{url: "", err: errors.New("invalid URI for request")},
		},
		{
			name: "oldurl",
			arg:  "http://www.youtube.com/",
			want: result{url: "aHR0cHM6Ly93d3cueW91dHViZS5jb20v", err: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GenerateShortUrl(tt.arg); got != tt.want.url && err != tt.want.err {
				t.Errorf("GenerateShortUrl() = %v, %v, want %v %v", got, err, tt.want.url, tt.want.err)
			}
		})
	}
}
