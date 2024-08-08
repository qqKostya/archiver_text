package vlc

import (
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



func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test TestEncode 1",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("BinaryChunks_ToHex() = %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		encodedText  string
		want string
	}{
		{
			name: "base test",
			encodedText: "20 30 3C 18 77 4A E4 4D 28", 
			want: "My name is Ted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.encodedText); got != tt.want {
				t.Errorf("BinaryChunks_ToHex() = %v, want: %v", got, tt.want)
			}
		})
	}
}
