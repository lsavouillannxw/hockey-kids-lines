package main

import (
	"flag"
	"fmt"
	"github.com/lsavouillannxw/hockey-kids-lines/service"
	"log"
	"time"
	"unicode/utf8"
)

func main() {
	defer DisplayTime(time.Now())
	numberOfPlayers := flag.Int("numberOfPlayers", 7, "the number of players you have")
	lineSize := flag.Int("lineSize", 3, "the size of a line")
	numberOfLines := flag.Int("numberOfLines", 10, "the number of lines you need during the game")

	flag.Parse()

	result := service.NewProcessingResult(*service.NewProcessingHandler(*numberOfPlayers, *numberOfLines, *lineSize).Process())
	for _, g := range result.BestMatch {
		for _, l := range g {
			fmt.Println(reverse_rmuller(l))
		}
		fmt.Println()
	}
}

func DisplayTime(start time.Time) {
	log.Printf("Program executed in %s", time.Since(start))
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
