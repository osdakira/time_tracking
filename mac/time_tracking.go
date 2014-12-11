package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func osascript() (path string) {
	path, err := exec.LookPath("osascript")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func frontAppName() string {
	cmd_str := `Tell Application "System Events"
set frontApp to name of first application process whose frontmost is true
end tell`

	cmd := exec.Command(osascript(), "-e", cmd_str)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func main() {
	for {
		fmt.Print(frontAppName())
		time.Sleep(1 * 1000)
	}
}
