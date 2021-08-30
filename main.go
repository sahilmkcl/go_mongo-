package main

func main() {
	var saveToken = make(map[string]string)
	r := registerRoutes(saveToken)
	r.Run(":8080")
}
