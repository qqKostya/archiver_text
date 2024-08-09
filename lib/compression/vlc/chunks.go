package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

type encodingTable map[rune]string

const chunkSize = 8

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}

	return res
}

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}

	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)

	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}

	return byte(num)
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
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

// Join joins chunks into one line and returns as string.
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, ch := range bcs {
		buf.WriteString(string(ch))
	}

	return buf.String()
}
