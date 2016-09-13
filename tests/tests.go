package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os/exec"
	"time"
)

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

//ClientRun goes
func ClientRun() {

	cmd := exec.Command("go", "run", "server/server.go")
	err := cmd.Start()
	if err != nil {

		panic(err)
	}

	time.Sleep(150 * time.Millisecond)
	err = cmd.Process.Kill()
	// err = cmd.Process.Signal(os.Interrupt)
	if err != nil {
		panic(err)
	}

}
func main() {
	go ClientRun()
	time.Sleep(100 * time.Millisecond)
	client, err := net.Dial("tcp", "localhost:9000")
	checkError(err)
	defer client.Close()
	c := jsonrpc.NewClient(client)

	var reply Reply
	var args *ArgsSum
	// var write *ArgsWrite
	// var read *ArgsRead

	args = &ArgsSum{2, 3}

	err = c.Call("MyServer.Sum", args, &reply)
	checkErrorfatal(err)
	fmt.Println("Sum is", reply.C)
	// write = &ArgsWrite{reply.C, "./String.txt"}
	//
	// err = c.Call("MyServer.Write", write, &reply)
	//
	// checkErrorfatal(err)
	//
	// read = &ArgsRead{"./String.txt"}
	//
	// err = c.Call("MyServer.Read", read, &reply)
	// checkErrorfatal(err)
	//
	// fmt.Println("Read from file:", reply.X)
}
