websocket-binarydata-example
============================

A sample dartlang app that displays binary data sent through websocket. The server is in golang which is capable of serving multiple clients from a single source of data.


Requirements
============

* Golang
* go.net : https://code.google.com/p/go/source/checkout?repo=net
* Dart SDK with Dartium browser : https://www.dartlang.org

How to run ?
============

1) Compile server.go :

$ go build server.go

2) Run golang server
$ ./server

3) Open Dartium browser and go to http://localhost:1234/web/websocket_binary_data_example.html


More info about the code @ http://dennisfrancis.wordpress.com


