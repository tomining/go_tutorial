package main

import (
	"text/scanner"
	"os"
	"fmt"
	"sync"
	"bytes"
	"flag"
	"path/filepath"
)

const BUF_SIZE = 1000

func parseArgs() (string, string) {
	flag.Parse()
	return *path, *pattern
}

func find(path string) <-chan string {
	out := make(chan string, BUF_SIZE)

	done := make(chan struct{}, workers)
	for i := 0; i < workers; i++ {
		go func() {
			filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
				out <- file
				return nil
			})
			done <- struct{}{}
		}()
	}

	go func() {	//모두 find한 후 종료하기 위함
		for i := 0; i < cap(done); i++ {
			<-done
		}
		close(out)
	}()

	return out
}

type partial struct {
	token string
	scanner.Position
}

func mapper(path string, out chan<- partial) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	//정상적인 파일이 아닌 경우 바로 반환
	if info, err := file.Stat(); err != nil || info.Mode().IsDir() {
		return
	}

	var s scanner.Scanner
	s.Filename = path
	s.Init(file)

	//파일의 모든 토큰을 스캔하여 out 채널로 전송
	tok := s.Scan()
	for tok != scanner.EOF {
		fmt.Print(s.Pos())
		out <- partial{s.TokenText(), s.Pos()}
		tok = s.Scan()
	}
}

func runMap(paths <-chan string) <-chan partial {
	out := make(chan partial, BUF_SIZE)
	go func() {
		for path := range paths {
			mapper(path, out)
		}
		close(out)
	}()
	return out
}

// key: tokin
// value: token position
type intermediate map[string][]scanner.Position

func (m intermediate) addPartial(p partial) {
	positions, ok := m[p.token]
	if !ok {
		positions = make([]scanner.Position, 1)
	}
	positions = append(positions, p.Position)
	m[p.token] = positions
}

func collect(in <-chan partial) intermediate {
	tokenPositions := make(intermediate, 10)
	for t := range in {
		tokenPositions.addPartial(t)
	}
	return tokenPositions
}

func reducer(token string, positions []scanner.Position) map[string]int {
	result := make(map[string]int)
	for _, p := range positions {
		result[p.Filename] += 1
	}
	return result
}

type summary struct {
	//key: token
	//value: map[string]int
	//		key: file path
	//		value: token count
	m map[string]map[string]int

	// m을 보호하기 위한 뮤텍스
	mu sync.Mutex
}

func (s summary) String() string {
	var buffer bytes.Buffer

	for token, value := range s.m {
		buffer.WriteString(fmt.Sprintf("Token: %s\n", token))
		total := 0
		for path, cnt := range value {
			if path == "" {
				continue
			}
			total += cnt
			buffer.WriteString(fmt.Sprintf("%8d %s ", cnt, path))
			buffer.WriteString("\n")
		}
		buffer.WriteString(fmt.Sprintf("Total: %d\n\n", total))
	}
	return buffer.String()
}

func runReduce(tokenPositions intermediate) summary {
	s := summary{m: make(map[string]map[string]int)}
	for token, positions := range tokenPositions {
		s.mu.Lock()
		s.m[token] = reducer(token, positions)
		s.mu.Unlock()
	}
	return s
}

func main() {
	path, _ := parseArgs()
	paths := find(path)
	fmt.Println(runReduce(collect(runMap(paths))))
}