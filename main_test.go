package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		inputFile  string
		outputFile string
	}{
		{
			"input1.txt",
			"output1.txt",
		},
		{
			"input2.txt",
			"output2.txt",
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			err := os.Setenv(InputFile, test.inputFile)
			if err != nil {
				t.Fail()
			}
			err = os.Setenv(OutputFile, test.outputFile)
			if err != nil {
				t.Fail()
			}

			expected := captureStdout(outputLines)
			actual := captureStdout(main)
			if expected != actual {
				t.Logf("expected: %s, but actual: %s", expected, actual)
				t.Fail()
			}
		})
	}
}

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	outputChannel := make(chan string)
	defer close(outputChannel)

	go func() {
		var buf strings.Builder
		_, _ = io.Copy(&buf, r)
		_ = r.Close()
		outputChannel <- buf.String()
	}()

	f()

	os.Stdout = stdout
	_ = w.Close()

	return <-outputChannel
}

func outputLines() {
	f, err := os.Open(OutputPath + os.Getenv(OutputFile))
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)

	for sc.Scan() && sc.Text() != "" {
		fmt.Println(sc.Text())
	}
}
