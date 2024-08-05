package vlc

import "testing"

func Test_prepareText(t *testing.T) {
	tests := []struct{
		name string
		str string
		want string
	}{
		{
			name: "base test",
			str: "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "2 Upper case",
			str: "MY nAmE IS ted",
			want: "!m!y n!am!e !i!s ted",
		},
		{
			name: "number and symbols",
			str: "My n4me i$ Ted!?>*",
			want: "!my n4me i$ !ted!?>*",
		},
		
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %#v, want: %#v", got, tt.want)
			}
		})
	}
}