package client_pkg

import (
	"net"
	"net/rpc/jsonrpc"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

/*
Aici am incercat sa fac cum mi-ai spus da reuseam
sa fac doar procese zombie :/ si cu server build-uit pus in folder
:) deci nu e ca si ccum doar commanda aia e gresita :))
*/

func ClientRun() {

	cmd := exec.Command("go", "run", "server/server.go")
	err := cmd.Start()
	if err != nil {

		panic(err)
	}

	time.Sleep(300 * time.Millisecond)
	err = cmd.Process.Kill()
	if err != nil {
		panic(err)
	}
}

func TestClient(t *testing.T) {
	go ClientRun()
	time.Sleep(100 * time.Millisecond)
	client, err := net.Dial("tcp", "localhost:9000")
	checkError(err)

	defer client.Close()

	c := jsonrpc.NewClient(client)

	var reply Reply
	var args *ArgsSum
	var write *ArgsWrite
	var read *ArgsRead

	args = &ArgsSum{2, 3}
	err = c.Call("MyServer.Sum", args, &reply)

	if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}
	if reply.C != args.Item1+args.Item2 {
		t.Errorf("Add: expected %d got %d", reply.C, args.Item1+args.Item2)
	}
	rez1 := reply.C

	write = &ArgsWrite{reply.C, "./Test_String.txt"}
	err = c.Call("MyServer.Write", write, &reply)
	if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}

	if reply.X != strconv.Itoa(rez1) {
		t.Errorf("Nu s-a scris string")
	}

	rez2 := reply.X
	read = &ArgsRead{"./Test_String.txt"}

	err = c.Call("MyServer.Read", read, &reply)
	if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}
	rez, err := strconv.Atoi(rez2)
	checkError(err)
	if rez != rez1 {
		t.Errorf("Nu s-a scris ce trebuie")
	}

}
