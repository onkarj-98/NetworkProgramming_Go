//The package "net" defines many types, functions and methods of use in Go network programming. The type
//IP is defined as byte slices
package main
import (
	"fmt"
	"net"
	"os"
)
func main(){
	if len(os.Args) != 2{
		fmt.Fprintf(os.stderr,"Usage: %s ip-addr\n",os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil{
		fmt.Println("Invalid address")
	}else{
		fmt.Println("The address is :",addr.String())
	}
	os.Exit(0)
} 