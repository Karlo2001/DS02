package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	id   int32
	name string
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		//DO SOME ERROR HANDLING
		fmt.Println("WRONG INPUT")
		os.Exit(1)
	}
	tmp, _ := strconv.Atoi(args[0])
	id = int32(tmp)
	name = args[1]
	startNode()
}

//Should start the node
func startNode() {

}

/*func GetCriticalAccess(*pb.ClientInfo) {

}*/
