package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

type HexChunks []HexChunk

type HexChunk string

type encodingTable map[rune]string

const chunkSize = 8


func (hcs HexChunks) ToString() string {
	const sep = " "

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))
	for _, hc := range hcs[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(hc))
	}

	return buf.String()
}

func (bsc BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bsc))

	for _, chunk := range bsc {
		hexChunk := chunk.ToHex()
		res = append(res, hexChunk)
	}

	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)

	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}

	// res := strings.ToUpper(fmt.Sprintf("%x", num)) // попробовать тут "X" заместо ToUpper
	res := fmt.Sprintf("%X", num)

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
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