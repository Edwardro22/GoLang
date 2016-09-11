package functi

import (
	"io/ioutil"
    "strconv"
    )
 

func  Suma(x int, y int) int {
    z:=x+y
    return z 
}

func Write(z int,s string)  string{
    x := strconv.Itoa(z)
	buf :=[]byte(x)
	err := ioutil.WriteFile(s, buf, 0644)
    if err != nil {
		panic(err)
	}
    return x 
}

func Read(s string)  string {
    filename := s
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r := string(content)
	return r
}