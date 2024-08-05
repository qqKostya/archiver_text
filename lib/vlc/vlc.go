package vlc

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

type encodingTable map[rune]string

const chunkSize = 8

func Encode(str string) string {
	str = prepareText(str)
	bStr := encodeBin(str)
	chunk := splitByChanks(bStr, chunkSize)
	fmt.Println(chunk)

	return ""
}

// splitByChunks splits binary string by chunks with given size,
// i.g.: 100101011001010110010101' -> 10010101 10010101 10010101'
func splitByChanks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunkCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunkCount++
	}

	res := make(BinaryChunks, 0, chunkCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}

// encodeBin encodes str into binary codes string without spaces
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()
	res, ok := table[ch]

	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'L': "001001",
		'U': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "",
		'x': "00000000001",
		'z': "000000000000",
	}
}

// prepareText prepares text to be fit in encode:
// changes upper case latters to: ! + lower case latter
// i.g.: My name is Ted -> !my name is !ted
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
