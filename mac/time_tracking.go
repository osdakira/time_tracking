package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd_str := `Tell Application "System Events"
set frontApp to name of first application process whose frontmost is true
end tell`

	path, err := exec.LookPath("osascript")
	if err != nil {
		log.Fatal(err)
	}
	osascript := path
	cmd := exec.Command(osascript, "-e", cmd_str)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out.String())
}
