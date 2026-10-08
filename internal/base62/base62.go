package base62

import (
	"errors"
	"slices"
	"strings"
)

const codecAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var (
	ErrEmpty            = errors.New("base62: empty code")
	ErrInvalidCharacter = errors.New("base62: invalid character")
	ErrOverflow         = errors.New("base62: code is too long")
)

// Encode renders a number with no padding.
func Encode(v uint64) string {
	// result
	var res []byte

	// div value by the "base" of the count system, so we have first (least significannt) symbol
	// and keep dividing till we have something in value
	// each time symbol will became more significant
	for {
		res = append(res, codecAlphabet[v%62])
		v /= 62
		if v == 0 {
			break
		}
	}

	// we syarted with least significant symbol. so, we have to reverse to get string
	// in right order
	slices.Reverse(res)

	return string(res)
}

// EncodeWidth does the same but left-pads with zeroes to at least width characters. The value itself is never truncated.
func EncodeWidth(v uint64, width int) string {

	// encode
	s := Encode(v)

	// will be reused
	len := len(s)

	// no need to pad
	if width <= len {
		return s
	}

	// calk len to pad
	var pad = width - len

	// do left paddign
	s = strings.Repeat("0", pad) + s

	return s
}

// Decode parses a code back into a number.
func Decode(s string) (uint64, error) {
	if s == "" {
		return 0, ErrEmpty
	}

	// check symbols
	var efflen uint64 = 0
	for _, r := range s {
		if !strings.ContainsRune(codecAlphabet, r) {
			return 0, ErrInvalidCharacter
		}
		efflen++
	}

	// check acceptable len
	if efflen > 11 {
		return 0, ErrOverflow
	}

	var v uint64 = 0
	for _, r := range s {
		idx := strings.IndexRune(codecAlphabet, r)
		v = v*62 + uint64(idx)
	}

	return v, nil
}
