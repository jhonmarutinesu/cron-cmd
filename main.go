package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/claudiu/gocron"
)

func cmdCommand() {

	output, err := exec.Command("cmd", "/c", "echo Hello World").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	fmt.Println(string(output))
}

func main() {

	gocron.Start()
	s := gocron.NewScheduler()
	gocron.Every(5).Seconds().Do(cmdCommand)

	<-s.Start()
}
