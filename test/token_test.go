package main

import (
	"fmt"
	"testing"

	"github.com/hello-slide/token-manager/token"
)

func TestCheckToken(t *testing.T) {
	key := []byte("YELLOW SUBMARINE, BLACK WIZARDRY")
	data := "hoge"

	generateToken, err := token.Create(data, key)
	if err != nil {
		t.Fatal(err)
	}

	newData, err := token.Verify(generateToken, key)
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

func TestDifferentKey(t *testing.T) {
	key1 := []byte("YELLOW SUBMARINE, BLACK WIZARDRY")
	key2 := []byte("YELLOW SUBMARINE, BLACK WIZARDRA")
	data := "hoge"

	generateToken, err := token.Create(data, key1)
	if err != nil {
		t.Fatal(err)
	}

	_, err = token.Verify(generateToken, key2)
	if err == nil {
		t.Fatal(err)
	}
}
