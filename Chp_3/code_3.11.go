package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(commaToFloat("+1234567.8910")) // Outputs: 1,234,567.8910
	fmt.Println(commaToFloat("+123456.78"))    // Outputs: -123,456.78
	fmt.Println(commaToFloat("+12345.0"))      // Outputs: +12,345.0
	fmt.Println(commaToFloat("+0.12"))         // Outputs: 12

}

func commaToFloat(s string) string {
	var buff bytes.Buffer

	if s[0] == '+' || s[0] == '-' {
		buff.WriteString(s[0:1])
		s = s[1:]
	}
	// split the string into integral and fractional part
	parts := strings.Split(s, ".")
	integralPart := parts[0]

	n := len(integralPart)
	remainder := n % 3
	if remainder == 0 && n > 0 {
		remainder = 3
	}
	buff.WriteString(integralPart[:remainder])
	for i := remainder; i < n; i += 3 {
		buff.WriteString(",")
		buff.WriteString(integralPart[i : i+3])
	}

	if len(parts) == 2 {
		buff.WriteString(fmt.Sprint("." + parts[1]))
	}

	return buff.String()
}
