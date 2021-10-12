package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

// refs: https://github.com/itchyny/maze
// refs: https://github.com/hajimehoshi/go-inovation/blob/main/ino/internal/field/field.go

type Point struct {
	X, Y int
}

type Game struct {
	mapData [][]int
	Height  int
	Width   int
	Start   *Point
	Goal    *Point
	Cursor  *Point
}

func readMapDataFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to readLine: %s\n", err)
		}

		strLine := string(line)

		fmt.Print(strLine)
		if !isPrefix {
			fmt.Println()
		}
	}
}

func main() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	x, y := 10, 10

	cat := '🙀'
	s.SetContent(x, y, cat, nil, defStyle)

	// Event loop
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		// Update screen
		s.Show()
		s.Clear()
		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			} else if ev.Rune() == 'h' {
				x -= 2
				s.SetContent(x, y, cat, nil, defStyle)
			} else if ev.Rune() == 'l' {
				x += 2
				s.SetContent(x, y, cat, nil, defStyle)
			} else if ev.Rune() == 'k' {
				y -= 1
				s.SetContent(x, y, cat, nil, defStyle)
			} else if ev.Rune() == 'j' {
				y += 1
				s.SetContent(x, y, cat, nil, defStyle)
			}
		}
	}
}
