// comments API.
package main

import (
	"github.com/mihnealun/commentix/infrastructure"
	"github.com/mihnealun/commentix/infrastructure/container"
)

func main() {
	containerInstance, err := container.GetInstance()
	if err != nil {
		panic(err.Error())
	}

	err = infrastructure.Start(containerInstance)
	if err != nil {
		panic(err.Error())
	}
}
