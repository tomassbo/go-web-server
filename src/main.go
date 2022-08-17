package main

func main() {

	server := NewServer(":5000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/user", CreateUser)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()

}
