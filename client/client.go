package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"playground/golang-protobuf-example/protox"
)

type ClientResponse struct {
	Greeting string `json:"greeting"`
	Response string `json:"response"`
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := protox.NewRequestHandlerClient(conn)

	g := gin.Default()

	g.GET("/one", func(ctx *gin.Context) {
		req := &protox.ExampleRequest{
			Message: "Call me maybe",
			Number:  24,
		}
		if resp, err := client.HandleOne(ctx, req); err == nil {
			ctx.JSON(200, ClientResponse{
				Greeting: resp.GetGreeting(),
				Response: resp.GetResponse(),
			})
		} else {
			ctx.String(500, "Something went wrong")
		}
	})

	g.GET("/two", func(ctx *gin.Context) {
		req := &protox.ExampleRequest{
			Message: "You are annoying",
			Number:  42,
		}
		if resp, err := client.HandleTwo(ctx, req); err == nil {
			ctx.JSON(200, ClientResponse{
				Greeting: resp.GetGreeting(),
				Response: resp.GetResponse(),
			})
		} else {
			ctx.String(500, "Something went wrong")
		}
	})

	err = g.Run(":8080")
	if err != nil {
		log.Fatal("Could not start gin server at port 8080")
	}
}
