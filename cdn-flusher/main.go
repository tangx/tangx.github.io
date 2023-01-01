package main

import "github.com/tangx/tangx.github.io/cdn-flusher/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
