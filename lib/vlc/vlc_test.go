package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test prepareText 1",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "base test prepareText 2",
			str:  "MY nAmE IS ted",
			want: "!m!y n!am!e !i!s ted",
		},
		{
			name: "base test prepareText 3",
			str:  "My n4me i$ Ted!?>*",
			want: "!my n4me i$ !ted!?>*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %#v, want: %#v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test encodeBin 1",
			str:  "!ted",
			want: "001000100110100101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %#v, want: %#v", got, tt.want)
			}
		})
	}
}

func Test_splitByChanks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "base test splitByChanks 1",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChanks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChanks() = %v, want: %v", got, tt.want)
			}
		})
	}
}
