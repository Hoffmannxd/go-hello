package main

func main() {
	r := registerRoutes()

	r.Run(":3000") // listen and serve on localhost:8080
}