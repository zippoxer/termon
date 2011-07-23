package term

/*
#cgo LDFLAGS: -lncurses
#include <ncurses.h>

void AddAt(int x, int y, const char* str) {
	mvprintw(x, y, str);
}

void Add(const char* str) {
	printw(str);
}
*/
import "C"

import (
	"unsafe"
	"fmt"
)

var currentPair = 1

var Cols *int
var Rows *int

type void unsafe.Pointer

type Color struct {
	pair int
	colorPair C.int
	FG, BG int
}

func NewColor(fg, bg int) Color {
	c := Color{currentPair, 0, fg, bg}
	c.init()
	currentPair++
	return c
}

func (clr *Color) init() {
	C.init_pair(C.short(clr.pair), C.short(clr.FG), C.short(clr.BG))
	clr.colorPair = C.COLOR_PAIR(C.int(clr.pair))
}

// Set the current color in the terminal to this.
func (clr *Color) On() {
	C.attron(clr.colorPair)
}

// Set the current color in the terminal to this.
func (clr *Color) Off() {
	C.attroff(clr.colorPair)
}

// Must be executed before any other function.
func Init() {
	C.initscr()
	if Colorable() {
		C.start_color()
	}
	
	Cols = (*int)(void(&C.COLS))
	Rows = (*int)(void(&C.LINES))
}

// Waits for a string input (until enter is pressed) and returns it.
func GetText() string {
	s := C.CString("")
	C.getstr(s)
	return C.GoString(s)
}

// Waits for a char input and returns it.
func GetChar() int {
	return int(C.getch())
}

// Prints string or char at the current position.
func Add(v interface{}) {
	switch t := v.(type) {
	case string:
		C.Add(C.CString(t))
	case byte:
		C.addch(C.chtype(t))
	case int:
		C.addch(C.chtype(t))
	}
}

// Prints string or char at the given position.
func AddAt(x, y int, v interface{}) {
	cx, cy := C.int(x), C.int(y)
	switch t := v.(type) {
	case string:
		C.AddAt(cy, cx, C.CString(t))
	case byte:
		C.mvaddch(cy, cx, C.chtype(t))
	case int:
		C.mvaddch(cy, cx, C.chtype(t))
	}
}

// Enables the reading of function keys (like F1, F2, arrow keys etc).
func Keypad() {
	C.keypad(C.stdscr, true)
}

// Show user's input (default).
func Echo() {
	C.echo()
}

// Hide user's input.
func Noecho() {
	C.noecho()
}

// Checks if the terminal supports colors.
func Colorable() bool {
	return bool(C.has_colors())
}

// Adds spaces to all cells in the terminal.
func Clear() {
	for x := 0; x < *Cols; x++ {
		for y := 0; y < *Rows; y++ {
			AddAt(x, y, ' ')
		}
	}
}

func HalfDelay(tof int) {
	C.halfdelay(C.int(tof))
}

func Bold() {
	C.attron(C.A_BOLD)
}

func Unbold() {
	C.attroff(C.A_BOLD)
}

func Refresh() {
	C.refresh()
}

func End() {
	C.endwin()
}

//
// A test!
//
func Test() {
	Init()
	c := NewColor(C.COLOR_RED, 0)
	
	c.On()
	AddAt(0, 8, "asd")
	AddAt(8, 8, fmt.Sprintf("There are %d rows. Say something about that: ", *Rows))
	
	// Get his input with default color.
	c.Off()
	s := GetText()

	c.On()
	Add(fmt.Sprintf("\"%s\" is not an answer.", s))
	
	GetChar()
	End()
}
