package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ShellHistory struct {
	History []string
	Count   int
}

func main() {
	shellLoop("yash> ", bufio.NewReader(os.Stdin))
}

func readLine(reader *bufio.Reader) string {
	var str string
	var in []byte
	var more bool

	in, more, _ = reader.ReadLine()
	for more {
		str += string(in)
		in, more, _ = reader.ReadLine()
	}
	str += string(in)
	return str
}

func shellLoop(prompt string, stdin *bufio.Reader) {
	run := true
	var inputString string
	var inputTokens []string
	var command string
	var bgJob bool
	var cmdErr error
	history := ShellHistory{make([]string, 1024, 1024), 0}
	for run {
		fmt.Printf(prompt)
		inputString = readLine(stdin)
		if len(inputString) <= 0 {
			continue
		}
		history.History[(history.Count)%cap(history.History)] = inputString
		history.Count++

		inputTokens = strings.Split(inputString, " ")
		command = inputTokens[0]
		if command == "quit" || command == "exit" {
			run = false
			break
		}
		if command == "history" {
			if len(inputTokens) == 2 && inputTokens[1] == "-c" {
				history.clear()
			} else if (len(inputTokens) == 2) && (inputTokens[1] != "-c") {
				fmt.Printf("history: Invalid flag %s\n", inputTokens[1])
			} else {
				history.printHistory()
			}
			continue
		}
		cmd := exec.Command(command, "")
		if inputTokens[len(inputTokens)-1] == "&" {
			cmd.Args = inputTokens[0 : len(inputTokens)-1]
			bgJob = true
		} else {
			cmd.Args = inputTokens[0:len(inputTokens)]
			bgJob = false
		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if bgJob {
			cmdErr = cmd.Start()
		} else {
			cmdErr = cmd.Run()
		}
		if cmdErr != nil {
			fmt.Printf("yash: %s\n", cmdErr.Error())
		}
	}
}

func (sh *ShellHistory) printHistory() {
	count := 0
	for count < sh.Count {
		fmt.Printf("%d\t%s\n", count+1, sh.History[count])
		count++
	}
}

func (sh *ShellHistory) clear() {
	sh.Count = 0
}
