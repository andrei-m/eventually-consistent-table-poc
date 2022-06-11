package main

import (
	"github.com/andrei-m/eventually-consistent-table-poc/server"
)

func main() {
	r := server.GetRouter()
	r.Run()
}
