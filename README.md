## Assignment #1 - File server 
Submitted by: Prerna Gupta, Roll: 143050021

## Introduction

This assignment is to build a simple file server, with a simple read/write interface. The file contents are stored in memory. The server listens on port 8080. Concurrent operation by multiple clients is taken into consideration

## Installation Instructions
<code>go get </code> github.com/prernaguptaiitb/cs733-assign1

Folder contains two files: server.go(containing server code) and concurrency_test.go (containing test cases for the server)

### How to use?
Run the server by "go run server.go" in one terminal from the assignment1 directory.
Run "telnet localhost 8080" on a different terminal to connect with the server as a telnet client.
The server supports multiple client connections by executing same command on different terminals and allows concurrent execution of these clients.
Once a connection is established one of the commands given in specification below can be run by a user through a client.

### Specification
* Write: create a file, or update the file’s contents if it already exists.
```
write <filename> <numbytes> [<exptime>]\r\n
 <content bytes>\r\n
```
exptime field is optional, it signifies the time in seconds after which the file on server 
is suppose to expire. If left unspecified or if its value is set to zero then file does not expires.

The server responds with the following:

```

OK <version>\r\n

``````
where version is a unique 64‐bit number (in decimal format) assosciated with the
filename.

* Read: Given a filename, retrieve the corresponding file:
```
read <filename>\r\n
```
The server responds with the following format if file is present at the server.
```
CONTENTS <version> <numbytes> <exptime> \r\n
 <content bytes>\r\n  
```
Here ```<exptime> ```is the remaining time in seconds left for the file after which it will expire. Zero v alue indicates that the file won't expire.

If the file is not present at the server or the file has expired then the server response is:
```
ERR_FILE_NOT_FOUND\r\n
```

* Compare and swap (cas): This replaces the old file contents with the new content
provided the version is still the same.
```
cas <filename> <version> <numbytes> [<exptime>]\r\n
 <content bytes>\r\n
```
exptime is optional and means the same thing as in "write" command.

The server responds with the new version if successful :-
```
OK <version>\r\n
```
If the file isn't found on the server or the file has expired then server response is:-
```
ERR_FILE_NOT_FOUND\r\n
```
If the version provided in the "cas" command does not match the version of the file, server
response is:-
```
ERR_VERSION <version>\r\n
```

* Delete file
```
delete <filename>\r\n
```
Server response (if successful)
```
OK\r\n
```

If the file isn't found on the server or the file has expired then server response is:-
```
ERR_FILE_NOT_FOUND\r\n
```

Apart from above mentioned errors if the above commands are not specified in proper format as mentioned, then server response is:-
```
ERR_CMD_ERR\r\n
```
After sending the above response to client, server terminates the client connection and thus forcing the client to connect again.



### Testing

For running the automated testing procedure, ensure that the server is not running and any other process is not busy at port 8080.
Run "go test -race" to run the tests. It would take some time.

The testing is divided into 2 parts:
* single client server communication
* concurrent execution of multiple clients

For a single client server communication all possible scenarios along with corner cases are tested. This program runs serially.

For each of the test cases the client sends a command to the server, reads the response and checks the result with the specified expected result mentioned in the test case.

For concurrent test cases, following test cases are considered:-
* 20 clients are spawned, out of these 10 clients execute normal routines commands on the file server. Each of these 10 clients operate on different files and write different contents. 
Each of the other 10 clients performs "write", "read", "cas" and "delete" operations (these operations also execute in parallel) on single file concurrently.

* 20 clients are spawned, out of these 10 clients perform write operation and other 10 performs cas operation on the same file in a loop. After their execution is complete, we check contents of the file and match it against the expected response ( which would the content written by any of the clients in their last iteration).

* 20 clients are spawned and each of them tries to perform cas operation on the same file. In this only one should succeed and others should get ```ERR_VERSION ``` error.


### Programming Details

* A structure is is used to store values corresponding to the file into a map. Key of map is filename and value is the structure.
* Channels are used to handle the race conditions created by the concurrent access of clients on the shared data structure.
* The server accepts connection on port 8080 and for every client handles its connection on a different thread.
* The command handling is done in a way that if the command format is not proper (like command length should not be greater than 500 bytes), then the connection is closed for the client issuing that command.

### Todo
1. To make the file server storage persistent suing leveldb
