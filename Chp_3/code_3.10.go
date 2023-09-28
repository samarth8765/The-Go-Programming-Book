package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(commaToInt("1234"))
	fmt.Println(commaToInt("12345"))
	fmt.Println(commaToInt("123456"))
}

func commaToInt(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	remainder := n % 3
	if remainder == 0 {
		remainder = 3
	}
	var buff bytes.Buffer
	buff.WriteString(s[:remainder])

	for i := remainder; i < n; i += 3 {
		buff.WriteString(",")
		buff.WriteString(s[i : i+3])
	}
	return buff.String()
}
