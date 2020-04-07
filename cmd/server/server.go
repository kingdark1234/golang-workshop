package main

import "workshop-service/internals/container"

func main() {
	sv, err := container.NewContainer()
	if err != nil {
		panic(err)
	}
	if err := sv.StartDatabase(); err != nil {
		panic(err)
	}
	if err := sv.Start(); err != nil {
		panic(err)
	}
}
