package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, request string) *Response {
	return &Response{"OK","ECHO: " + request}
}

func (server *EchoServer) Name() string  {
	return "EchoServer"
}

func TestIpc(t *testing.T)  {
	echo := EchoServer{}
	fmt.Println("",echo)
	server := NewIpcServer(&echo)

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, err1 := client1.Call("hello","From Client1")
	resp2, err2 := client2.Call("hello","From Client2")

	if err1 != nil || err2 != nil {
		fmt.Println("err1 or err2 is not nil")
	}

	if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2: ", resp2)
	}

	client1.Close()
	client2.Close()


}
