package client_pkg

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

type Operations interface {
	Sum(args *ArgsSum, reply *Reply) error
	Write(args *ArgsWrite, reply *Reply) error
	Read(args *ArgsRead, reply *Reply) error
}

type ArgsSum struct {
	Item1 int
	Item2 int
}
type ArgsWrite struct {
	Item     int
	FilePath string
}
type ArgsRead struct {
	FilePath string
}
type Reply struct {
	C int
	X string
}

func checkErrorfatal(err error) {
	if err != nil {
		log.Fatal("MyServer error:", err)
	}
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Client(p int, n int) {

	client, err := net.Dial("tcp", "localhost:9000")
	checkError(err)

	defer client.Close()

	c := jsonrpc.NewClient(client)

	var reply Reply
	var args *ArgsSum
	var write *ArgsWrite
	var read *ArgsRead

	args = &ArgsSum{p, n}

	err = c.Call("MyServer.Sum", args, &reply)
	checkErrorfatal(err)

	write = &ArgsWrite{reply.C, "./String.txt"}

	err = c.Call("MyServer.Write", write, &reply)
	checkErrorfatal(err)

	read = &ArgsRead{"./String.txt"}

	err = c.Call("MyServer.Read", read, &reply)
	checkErrorfatal(err)

	fmt.Println("Read from file:", reply.X)
}
