package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestChunkDir(t *testing.T) {
	err := ReadLine("/Users/terrill/chunk_dir.dat", func(s string) {
		s = strings.Trim(s, ".")
		s = strings.Trim(s, "{")
		s = strings.Trim(s, "}")
		s = strings.Trim(s, "\"")
		fmt.Println(s)
		fmt.Println(GetAllFile(s))
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestChunks(t *testing.T) {
	fmt.Println(chunks())
}
