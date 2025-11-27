package main

import (
	"bytes"
	"strings"
)

func Comma(s string) string {

	n := len(s)

	if n <= 3 {
		return s
	}

	return Comma(s[:n-3]) + "," + s[n-3:]
}

func CommaWithoutRecursion(s string) string {
	var buf bytes.Buffer
	sign := ""

	if s == "" {
		return s
	}

	if s[0] == '+' || s[0] == '-' {
		sign = s[:1]
		s = s[1:]
	}

	frac := ""

	if idx := strings.LastIndex(s, "."); idx >= 0 {
		frac = s[idx:]
		s = s[:idx]
	}

	n := len(s)

	if n <= 3 {
		return sign + s + frac
	}

	pre := n % 3
	if pre == 0 {
		pre = 3
	}

	buf.WriteString(s[:pre])
	for i := pre; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return sign + buf.String() + frac
}
