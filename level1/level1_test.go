package main

import (
	"testing"

	"github.com/aburd/cryptochallenge/internal"
)

func TestHexToBase64(t *testing.T) {
	var (
		original string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
		expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	)
	res := util.HexToBase64(original)
	if res != expected {
		t.Errorf("HexToBase64 was incorrect, got: %s, want: %s.", res, expected)
	}
}

func TestHexToFixedXor(t *testing.T) {
	var (
		original string = "1c0111001f010100061a024b53535009181c"
		against  string = "686974207468652062756c6c277320657965"
		expected string = "746865206b696420646f6e277420706c6179"
	)
	res := util.HexToFixedXor(original, against)
	if res != expected {
		t.Errorf("HexToFixedXor was incorrect, got: %s, want: %s.", res, expected)
	}
}
