package main

func main() {
	router := Server()
	router.Run(":8080")
}
