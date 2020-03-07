package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type IScanner interface {
	ReadLine() string
}

type Scanner struct {
	scanner *bufio.Scanner
}

const (
	InputPath  = "./input/"
	OutputPath = "./output/"
	InputFile  = "INPUT"
	OutputFile = "OUTPUT"
)

func NewScanner() IScanner {
	var reader io.Reader
	var err error
	input := os.Getenv(InputFile)
	if input == "" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(InputPath + input)
		if err != nil {
			panic(err)
		}
	}
	return &Scanner{scanner: bufio.NewScanner(reader)}
}

func (sc *Scanner) ReadLine() string {
	sc.scanner.Scan()
	return sc.scanner.Text()
}

func main() {
	sc := NewScanner()

	// TODO : Write Answer Code
	n, _ := strconv.Atoi(sc.ReadLine())
	sum := 0
	factorial := 1
	for i := 0; i < n; i++ {
		m, _ := strconv.Atoi(sc.ReadLine())
		sum += m
		factorial = factorial * m
	}

	// TODO : Output Answer
	fmt.Println(sum)
	fmt.Println(factorial)
}
