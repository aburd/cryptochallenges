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