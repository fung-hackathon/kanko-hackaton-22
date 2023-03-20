package main

import "kanko-hackaton-22/app/interfaces"

func main() {
	s := interfaces.NewServer()
	s.Serve()
}
