package main

import "5g-v2x-data-management-service/internal/container"

func main() {
	c := container.NewContainer()
	if err := c.Run().Error; err != nil {
		panic(err)
	}
}
