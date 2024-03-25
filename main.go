package main

import (
	"IoCExample/src"
	"fmt"
)

func main() {
	var container *src.Container

	container = src.NewContainer()

	container.AddService("IProductService", "ProductService")
	container.AddService("IUserService", "UserService")
	container.AddService("ILogService", "LogService")

	myConcreteService := container.GetService("ILogService")
	fmt.Println(myConcreteService)

	fmt.Println("Map:", container.GetServices())
}
