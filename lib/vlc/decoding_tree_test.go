package vlc

import ("testing")

func Test_encodingTable_Decoding(Treet *testing. T) {
	tests := []struct{
		name string
		et encodingTable
		want DecodingTree
	} {
		{
			name: "base tree test",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: DecodingTree{
				Zero: &DecodingTree{
					One: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &DecodingTree{
					One: &DecodingTree{},
					Zero: &DecodingTree{},
				},
			},
		},
	}
}
