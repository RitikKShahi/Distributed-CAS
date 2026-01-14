package main

import (
	"log"

	"github.com/RitikKShahi/DistributedCAS/p2p"
)

func main()  {
	tr := p2p.NewTCPTransport(":4000")
	err :=tr.ListenAndAccept()
	if err!=nil{
		log.Fatal(err)
	}
	select{}
}