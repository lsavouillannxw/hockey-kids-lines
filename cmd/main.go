package main

import (
	"flag"
	"HockeyLines/service"
	"fmt"
	"unicode/utf8"
)

func main() {
	numberOfPlayers:= flag.Int("numberOfPlayers", 9, "the number of players you have")
	numberOfLines:= flag.Int("numberOfLines", 10, "the number of lines you need during the game")
	lineSize:= flag.Int("lineSize", 4, "the size of a line")

	flag.Parse()

	result := service.NewProcessingHandler().Process(*numberOfPlayers, *numberOfLines, *lineSize)
	for _, m := range result.BestMatch {
		for _, l := range m {
			fmt.Println(reverse_rmuller(l))
		}
		fmt.Println()
	}
}

func reverse_rmuller(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}