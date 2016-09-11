package server_pkg


import (

	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"functi"
)

type Operations interface{
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

type MyServer int

type MyServerAddResp struct{
	Id     interface{} `json:"id"`
    Result Reply       `json:"result"`
    Error  interface{} `json:"error"`
}


func (srv *MyServer) Sum(args *ArgsSum, reply *Reply) error {
    reply.C = functi.Suma(args.Item1,args.Item2)
    return nil
}


func (srv *MyServer) Write(args *ArgsWrite, reply *Reply) error {
	reply.X =functi.Write(args.Item, args.FilePath)
	return nil
}

func (srv *MyServer) Read(args *ArgsRead, reply *Reply) error {
	reply.X =functi.Read(args.FilePath)
	return nil
}
func checkErrorfatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func StartServer(){
myserver := new(MyServer)

	server := rpc.NewServer()
	server.Register(myserver)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
    l, err := net.Listen("tcp", "localhost:9000")
    if err != nil {
        log.Fatalln(err.Error())
    }
  defer l.Close()
 	for{
        conn, err := l.Accept()
        if err !=nil {
            log.Fatalln(err.Error())
			
        }
		defer conn.Close()
         go server.ServeCodec(jsonrpc.NewServerCodec(conn))
     }


}

