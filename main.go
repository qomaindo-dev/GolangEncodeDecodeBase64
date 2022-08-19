package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	textString := "Nama:Qomaindo"
	encode := base64Encode(textString)

	fmt.Println("encode: ", encode)

	decode, valid := base64Decode(encode)
	if !valid {
		fmt.Println("Tidak valid euy")
	}

	fmt.Println("decode: ", decode)
}

func base64Encode(inputText string) string {
	return base64.StdEncoding.EncodeToString([]byte(inputText))
}

func base64Decode(plainText string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(plainText)
	if err != nil {
		return "", false
	}

	return string(data), true
}
