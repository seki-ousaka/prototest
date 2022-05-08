package main

import(
	"log"
	"prototest/exec"
	"prototest/serve"
)

func init() {
	log.SetPrefix("[main] ")
}

func main() {
	exec.Exec()
	serve.Serve()
}