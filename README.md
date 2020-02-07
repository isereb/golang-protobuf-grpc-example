# Golang protobuf and gRPC example

This project demonstrates client-server architecture, where client serves gRPC and client serves HTTP. When request is made to the client, client calls server via gRPC using protobuf and then translates it to HTTP and returns JSON.

Server port: 4040

Client port: 8080

Quick demo:
1) Start the server
2) Start the client
3) Open your browswer and type in http://localhost:8080/one
4) Viola , you get the response
