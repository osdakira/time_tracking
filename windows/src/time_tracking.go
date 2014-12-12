package main

import (
	"bytes"
	"encoding/csv"
	// "fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func frontAppName() string {
	directory := filepath.Dir(os.Args[0])
	cmd := exec.Command(filepath.Join(directory, "getForegroundWindowName.exe"))

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
		writer.Write([]string{strconv.FormatInt(time.Now().Unix(), 10), appName})
		writer.Flush()
		time.Sleep(1 * 1000 * 1000 * 1000)
		i++
		if i%5 == 0 {
			writer.Flush()
			i = 0
		}
	}
}
