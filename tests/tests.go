package main

import (
	"log"
	"os/exec"
	"net"
	"os"
	"net/rpc/jsonrpc"
)

type ArgsSum struct {
	Item1 int
	Item2 int
}
type Reply struct {
	C int
	X string
}
var C *exec.Cmd

 func start(cmd *exec.Cmd) {	
	log.Printf("Waiting for Server to Start...")
	/*cmd := exec.Command("./server")*/
	log.Printf("Aici e pusa cmd.Start()")
	err := cmd.Start()
	log.Printf("Command finished with error: %v", err)
    
	
}

func main() {
	C := exec.Command("./server")
	go start(C)
	
	client, err := net.Dial("tcp", "localhost:9000")
	log.Printf("Command finished with error: %v", err)

	defer client.Close()

	c := jsonrpc.NewClient(client)
	var reply Reply
	var args *ArgsSum
	
	args = &ArgsSum{2, 3}
	err = c.Call("MyServer.Sum", args, &reply)
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}


	log.Printf("Waiting for command to be killed...")
	 err = C.Process.Signal(os.Interrupt);
	log.Printf("Command finished with error: %v", err)
}
