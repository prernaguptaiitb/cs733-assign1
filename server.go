package main

import (
	"net"
	"fmt"
	"log"
	"bufio"
	"strings"
//	"strconv"
	"sync"
)

// File Object
type File struct {
	filename string
	version int64
	expTime int
	numBytes int
	contentBuf []byte
	sync.RWMutex
}

// map to retrieve file from filename
var filemap map[string]File

func ErrorInvalidCmd(cmd string, Connect net.Conn){
	var msg_send string
	switch cmd {

		case "write":
			msg_send = "ERR_CMD_ERR\r\n Correct format : write <filename> <numbytes> [<exptime>]\r\n <content bytes>\r\n"
		case "read":
			msg_send = "ERR_CMD_ERR\r\n Correct format : read <filename>\r\n"
		case "cas":
			msg_send = "ERR_CMD_ERR\r\n Correct format : cas <filename> <version> <numbytes> [<exptime>]\r\n <content bytes>\r\n"
		case "delete":
			msg_send = "ERR_CMD_ERR\r\n Correct format : delete <filename> \r\n"
		case "default":
			msg_send = "ERR_CMD_ERR\r\n Invalid Command\r\n"

	}
	Connect.Write([]byte(msg_send))
}

func WriteFile(filename string, numBytes string, expTime string) {
	fmt.Println("In write")
	
}

func readFile(filename string){
	fmt.Println("In read")
}

func compareAndSwap( filename string, version string, numBytes string, expTime string){
	fmt.Println("In cas")
}

func deleteFile(filename string){
	fmt.Println("In Delete")
}

func IsValidCmd( InpCommand string, Connect net.Conn) {
	tokenizedCmd := strings.Fields(InpCommand)
	l := len(tokenizedCmd)
	if l==0 {
		ErrorInvalidCmd("default", Connect)
	}
	if tokenizedCmd[0] == "write" {
		if l==3 {
			go WriteFile( tokenizedCmd[1], tokenizedCmd[2], "0")
		} else if l==4 {
			go WriteFile( tokenizedCmd[1], tokenizedCmd[2], tokenizedCmd[3])
		} else {
			ErrorInvalidCmd("write", Connect)
		}
	} else if tokenizedCmd[0]=="read" {
		if l==2 {
			go readFile(tokenizedCmd[1]) 
		} else{
			ErrorInvalidCmd("read", Connect)
		}

	} else if tokenizedCmd[0] == "cas" {
		if l==4 {
			go compareAndSwap(tokenizedCmd[1], tokenizedCmd[2], tokenizedCmd[3],"0")
		}else if l==5{
			go compareAndSwap(tokenizedCmd[1], tokenizedCmd[2], tokenizedCmd[3], tokenizedCmd[4])
		}else{
			ErrorInvalidCmd("cas", Connect)
		}

	}else if tokenizedCmd[0] == "delete"{
		if l==2 {
			go deleteFile(tokenizedCmd[1])
		} else{
			ErrorInvalidCmd("delete", Connect)
		}
	} else {
		ErrorInvalidCmd("default", Connect)
	}
}


func handleThisClient(Connect net.Conn){
	for{
		msg_rec,err := bufio.NewReader(Connect).ReadString('\n')
		if err != nil {
			fmt.Println("Error in recieving messages")
			Connect.Close()
			break
		}
//		fmt.Print(string(msg_rec))
		IsValidCmd(string(msg_rec),Connect)
		msg_send := "client connected to server\n"
		Connect.Write([]byte(msg_send))
	}
}

func serverMain(){
	fmt.Println("Server started")
	listen,err := net.Listen("tcp",":8080")
	if err != nil {
		fmt.Println("Error in server listening to port 8080")
		log.Fatal(err)
	}
	for{
		connect,err := listen.Accept()
		if err != nil {
			fmt.Println("Error in connecting to client")
			log.Fatal(err)
		}
		go handleThisClient(connect)
	}
}

func main(){

	serverMain()

}
