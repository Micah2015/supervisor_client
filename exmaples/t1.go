package main

import (
	"fmt"
	"os"

	"github.com/Micah2015/supervisor_client"
)

func main() {

	c, err := supervisor_client.NewSupervisorRpcClient("/opt/ws/run/supervisor.sock")
	if err != nil {
		fmt.Println("err: ", err.Error())
		os.Exit(1)
	}

	methods, _ := c.ListMethods()
	fmt.Println("methods:", methods)

	ver, _ := c.GetSupervisorVersion()
	fmt.Println("version:", ver)

	state, _ := c.GetSupervisorState()
	fmt.Println("state:", state)

	infos, _ := c.GetAllProcessInfo()
	fmt.Printf("%#v\n", infos)
}
