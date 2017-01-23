package main

import "log"

type People struct {
	Name string
	Age  int
}

func PONE() {
	g1 := Init()
	g1.Connect("localhost", "golang")
	g1.Definition("people", People{})

	log.Println(g1.DB.Name)
}
