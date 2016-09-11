package main

import(
    "os"
    "os/exec"
    "net"
	"net/rpc/jsonrpc"
	 /*"os/signal" */  
    
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Client_Run(){
		
cmd := exec.Command("./server")
cmd.Stderr = os.Stderr

/*cmd := exec.Command("./server")*/

err := cmd.Start()
if err != nil {

    panic(err)
}
_,err =cmd.StderrPipe()
if err != nil {

    panic(err)
}
 err = cmd.Process.Kill();
    if err != nil {
        panic(err)
 	}
}



func Client_close(){
	
	/*cmd := exec.Command("go","run","/home/eduardt/Desktop/gothing/src/server/server.go")*/
/*	cmd := exec.Command("./server")*/
/*	 err := cmd.Process.Kill();
    if err != nil {
        panic(err)
 	}*/
}

func main(){
go Client_Run()

    client, err := net.Dial("tcp", ":9000")
	checkError(err)

	defer client.Close()

	c := jsonrpc.NewClient(client)
    
    var reply Reply
	var args *ArgsSum
   /* var write *ArgsWrite
    var read *ArgsRead*/
    
    args = &ArgsSum{2, 3}
	err = c.Call("MyServer.Sum", args, &reply)
	/*if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}
	if reply.C != args.Item1+args.Item2 {
		t.Errorf("Add: expected %d got %d", reply.C, args.Item1+args.Item2)
	} 
    rez1:=reply.C
	write = &ArgsWrite{reply.C, "./Test_String.txt"}
    	err = c.Call("MyServer.Write", write, &reply)
        	if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}
	if reply.X != strconv.Itoa(rez1) {
		t.Errorf("Nu s-a scris string")
	}
    rez2 :=reply.X
    read = &ArgsRead{"./Test_String.txt"}
    err = c.Call("MyServer.Read", read, &reply)
	if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}
    rez,err:= strconv.Atoi(rez2)
    checkError(err)
	if  rez != rez1 {
		t.Errorf("Nu s-a scris ce trebuie")
	}*/
go Client_close()
}