package main

func main() {
	a := app{}
	a.Initialize(nil)
	a.initializeRoutes()
	a.Run(":8080")
}
