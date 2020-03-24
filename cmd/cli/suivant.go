package main

import (
	"bufio"
	"flag"
	"fmt"
	suivant "github.com/gypsydave5/au-suivant"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	var delay = flag.Int("delay", 10, "delay in seconds")
	flag.Parse()

	names := readNames(os.Stdin)

	s := suivant.New(names, time.Second*time.Duration(*delay))
	out := s.Start()

	for name := range out {
		notifyDriverChange(name)
	}
}

func readNames(stdin io.Reader) []string {
	var names []string
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	return names
}

func notifyDriverChange(name string) {
	title := "NEXT!"
	message := fmt.Sprintf("%s is Next", name)
	osa, _ := exec.LookPath("osascript")
	notify := exec.Command(osa, "-e", `display notification "`+message+`" with title "`+title+`"`)
	speak := exec.Command(osa, "-e", fmt.Sprintf("say %q", message))
	_ = notify.Run()
	_ = speak.Run()
}
