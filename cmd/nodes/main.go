package main

import (
	"Mycelium/pkg/graph"
	reporter2 "Mycelium/pkg/reporter"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//go:embed index.html
var indexHTML []byte

// Handles up to 1 million fairly well during configure

const NODE_COUNT = 10

var reporter = reporter2.NewReporter()

var graphTest = graph.NewGraph()

func main() {

	// TODO: Handle deleting a node and how a node will handle that error
	// TODO: Handle communication between nodes (heartbeat kind of thing)

	var nodes []*Node

	go func() {
		// Initial 3-second delay to allow the initial configuration to stabilize
		time.Sleep(3000 * time.Millisecond)

		for {
			println("TICK")
			for _, n := range nodes {
				n.RunChan <- Exec{fn: func(n *Node) {
					n.Tick()
				}}
			}

			time.Sleep(1 * time.Second) // Tick per second
		}
	}()

	go StartServer()

	for i := 0; i <= NODE_COUNT; i++ {
		//time.Sleep(1000 * time.Millisecond)

		node := NewNode(i)

		nodes = append(nodes, node)

		reporter.NewNode(node.Id)

		// TODO: can add random sleeps here to make it so  that nodes get different ages

		if i > 0 {
			// Create a random grid at first. Then afterward, each node will randomly re-configure
			//node.ConfigureV2(nodes[rand.IntN(i)])
		}

	}

	// Graph Test:
	go func() {
		next := 6
		n := []int{1, 2, 3, 4, 5}

		graphTest.FullyConnected(distance, n...)

		time.Sleep(5 * time.Second)

		go func() {
			for {
				graphTest.Disconnect()
				time.Sleep(1 * time.Second)
			}
		}()

		go func() {
			for {
				time.Sleep(5 * time.Second)
				graphTest.PartiallyConnected(distance, next, rand.IntN(next-1)+1)
				next++
			}
		}()

	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	for _, n := range nodes {
		n.Stop()
	}
}

func distance(a int, b int) int {
	return a + b + 1
}

func StartServer() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/graph", func(c *gin.Context) {
		//c.JSON(http.StatusOK, reporter.GetDoc())
		c.JSON(http.StatusOK, graphTest.Export())
	})

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")

		_, _ = c.Writer.Write(indexHTML)
	})

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
