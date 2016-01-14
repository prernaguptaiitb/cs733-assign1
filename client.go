package main

import "net"
import "fmt"
import "bufio"
import "os"

func clientMain(){

	//connect to server
	connect,err := net.Dial("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for{
		//read action from user
		reader := bufio.NewReader(os.Stdin)
		msg_send,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return 
		}
		//write on socket
		connect.Write([]byte(msg_send))
		//Read from socket
		msg,err := bufio.NewReader(connect).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(msg)
	}

}


func main(){

	clientMain()

}
