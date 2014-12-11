package main

import (
	"bytes"
	"encoding/csv"
	// "fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
	return strings.TrimSpace(out.String())
}

func main() {
	var appName string
	var i int
	file, _ := os.OpenFile("tracking.csv", os.O_WRONLY|os.O_CREATE, 0600)
	writer := csv.NewWriter(file)
	for {
		appName = frontAppName()
		// fmt.Print(appName)
		writer.Write([]string{appName, time.Now().Format(time.RFC850)})
		time.Sleep(1 * 1000 * 1000 * 1000)
		i++
		if i%5 == 0 {
			writer.Flush()
			i = 0
		}
	}
}
