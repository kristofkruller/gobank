package main

func main() {
	server := NewAPIServer(":209")
	server.Run()
}
