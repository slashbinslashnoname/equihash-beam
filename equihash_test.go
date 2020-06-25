package equihash

import (
	"encoding/hex"
	"testing"
)

func TestVerifySolution(t *testing.T) {
	input := "34b6e7b63e12372f7df4d6dac9935b753b7b652f02f24c958d20c57c33678ae1"
	nonce := "0ffffff3d6b80532"
	output := "014d5b42581fe065c8b82d219804d7190763273114cc4c49ad9506357047a8fb7139761874d211067417037f5d2118f8346185fc04d21487fa1eb062fec07463010aef5046bfe3a1113c1054f88c078e0d84c93c80bed3b44e736e089c10823230d14472e07bc844"
	inputBytes, _ := hex.DecodeString(input)
	nonceBytes, _ := hex.DecodeString(nonce)
	outputBytes, _ := hex.DecodeString(output)

	valid := VerifySolution(inputBytes, nonceBytes, outputBytes)
	if !valid {
		t.Fatal("Invalid")
	}

	nonce = "94ac0343d8a411" // edit the last character to 8. others remain unchange.
	nonceBytes, _ = hex.DecodeString(nonce)

	valid = VerifySolution(inputBytes, nonceBytes, outputBytes)
	if valid {
		t.Fatal("Invalid")
	}
}
