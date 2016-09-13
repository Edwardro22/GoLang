package functi

import (
	"log"
	"strconv"
	"testing"
)

var X int

func TestSuma(t *testing.T) {
	x := 2
	y := 3
	X = Suma(x, y)
	if X != x+y {
		log.Fatalln("Suma nu-i ce trebe %d + %d =%d", x, y, X)
	}

}

func TestWrite(t *testing.T) {
	x := Write(X, "./Test_functi.txt")

	if x != strconv.Itoa(X) {
		log.Fatalln("Nu s-o scris ce trebe")
	}
}

func TestRead(t *testing.T) {
	x := Read("./Test_functi.txt")
	if x != strconv.Itoa(X) {
		log.Fatalln("Nu s-o citit ce trebe")
	}
}
