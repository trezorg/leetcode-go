package multiplystrings

import (
	"strconv"
	"strings"
)

func multiplyOne(s string, c int) string {
	res := make([]string, len(s)+1)
	rest := 0
	ln := len(s)
	for i := ln - 1; i >= 0; i-- {
		c1, _ := strconv.Atoi(string(s[i]))
		v := c1*c + rest
		rest = v / 10
		res[i+1] = strconv.Itoa(v % 10)
	}
	if rest > 0 {
		res[0] = strconv.Itoa(rest)
	}
	return strings.Join(res, "")
}

func multiplyOneFunc(s string) func(c int) string {
	mp := map[int]string{}

	return func(c int) string {
		v, ok := mp[c]
		if ok {
			return v
		}
		v = multiplyOne(s, c)
		mp[c] = v
		return v
	}

}

func addStrings(s1 string, s2 string) string {

	if len(s1) == 0 {
		return s2
	} else if len(s2) == 0 {
		return s1
	}

	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}

	rest := 0

	i := len(s1) - 1
	j := len(s2) - 1
	res := make([]string, len(s1)+1)

	for i >= 0 && j >= 0 {
		c1, _ := strconv.Atoi(string(s1[i]))
		c2, _ := strconv.Atoi(string(s2[j]))
		v := c1 + c2 + rest
		rest = v / 10
		res[i+1] = strconv.Itoa(v % 10)
		i--
		j--
	}

	for i >= 0 {
		s := string(s1[i])
		if rest > 0 {
			c1, _ := strconv.Atoi(s)
			v := c1 + rest
			rest = v / 10
			res[i+1] = strconv.Itoa(v % 10)
		} else {
			res[i+1] = s
		}
		i--
	}

	if rest > 0 {
		res[i+1] = strconv.Itoa(rest)
	}

	return strings.Join(res, "")

}

func multiply(num1 string, num2 string) string {
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	s := ""
	count := 0
	fc := multiplyOneFunc(num1)
	for i := len(num2) - 1; i >= 0; i-- {
		c, _ := strconv.Atoi(string(num2[i]))
		sr := fc(c) + strings.Repeat("0", count)
		s = addStrings(s, sr)
		count++
	}

	i := 0

	for i < len(s)-1 && string(s[i]) == "0" {
		i++
	}

	return s[i:]

}
