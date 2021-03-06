package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS != "darwin" {
		cmd := exec.Command("/usr/bin/lsb_release", "-a")
		var res bytes.Buffer
		cmd.Stdout = &res

		if err := cmd.Run(); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(res.String())
	} else {
		fmt.Println(runtime.GOOS)
	}
}
