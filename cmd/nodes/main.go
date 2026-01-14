package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Handles up to 1 million fairly well during configure

const NODE_COUNT = 10

var reporter = Reporter{
	nodes: make([]int, 0),
	edges: make([]Edge, 0),
}

func main() {

	// TODO: Handle deleting a node and how a node will handle that error
	// TODO: Handle communication between nodes (heartbeat kind of thing)

	var nodes []*Node

	go func() {
		for {
			println("TICK")
			for _, n := range nodes {
				n.RunChan <- Exec{fn: func(n *Node) {
					n.Tick()
				}}
			}

			time.Sleep(time.Second) // Tick per second
		}
	}()

	go StartServer()

	for i := 0; i <= NODE_COUNT; i++ {
		node := NewNode(i)

		nodes = append(nodes, node)

		reporter.NewNode(node.Id)

		// TODO: can add random sleeps here to make it so  that nodes get different ages

		if i > 0 {
			// Create a random grid at first. Then afterward, each node will randomly re-configure
			node.ConfigureV2(nodes[rand.IntN(i)])
		}

		time.Sleep(time.Second)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	for _, n := range nodes {
		n.Stop()
	}
}

func StartServer() {
	r := gin.Default()

	r.Use(CORSMiddleware()) // Allows all origins

	// Define a simple GET endpoint
	r.GET("/graph", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, reporter.GetDoc())
	})

	r.GET("/", func(c *gin.Context) {
		// read from file
		data, err := os.ReadFile("/Users/davidlamar/Projects/Mycelium/cmd/nodes/index.html")
		if err != nil {
			// error handler
		}

		c.Header("Content-Type", "text/html")

		_, _ = c.Writer.Write(data)
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
