package main

import (
	"strconv"
	"strings"
)

type Codec struct{}

func (c *Codec) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteString("#")
		sb.WriteString(s)
	}
	return sb.String()
}

func (c *Codec) Decode(s string) []string {
	var res []string
	i := 0
	for i < len(s) {
		j := i

		for s[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(s[i:j])
		i = j + 1

		res = append(res, s[i:i+length])
		i += length
	}
	return res
}
