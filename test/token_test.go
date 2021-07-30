package main

import (
	"fmt"
	"testing"

	"github.com/hello-slide/token-manager/token"
)

func TestCheckToken(t *testing.T) {
	token.Key = []byte("YELLOW SUBMARINE, BLACK WIZARDRY")
	data := "hoge"

	generateToken, err := token.Create(data)
	if err != nil {
		t.Fatal(err)
	}

	newData, err := token.Verify(generateToken)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(generateToken)
	fmt.Println(newData)
	fmt.Println(data)

	if newData != data {
		t.Fatal("The compounded data is different.")
	}

}
