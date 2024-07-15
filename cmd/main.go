package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	for {
		initCommands()

		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %s\n", err)
			os.Exit(1)
		}

		input = strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")
		input = strings.TrimSpace(input)
		inputs := strings.Split(input, "\x20")

		command := inputs[0]
		arguments := inputs[1:]

		builtInCommand, isCommandPresent := builtInCommands[command]

		if isCommandPresent {
			builtInCommand(arguments)
		} else {
			execFn(command, arguments)
		}

	}
}

type builtInFunc func([]string)

var builtInCommands = make(map[string]builtInFunc)

func initCommands() {
	builtInCommands["exit"] = exitFn
	builtInCommands["echo"] = echoFn
	builtInCommands["type"] = typeFn
	builtInCommands["pwd"] = pwdFn
	builtInCommands["cd"] = cdFn
}

// Built in commands

func exitFn(arguments []string) {
	if len(arguments) < 1 {
		os.Exit(0)
	}
	stringifiedExitCode := arguments[0]
	exitCode, err := strconv.Atoi(stringifiedExitCode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input exit code should be a number: %s\n", err)
	}
	os.Exit(exitCode)
}

func echoFn(arguments []string) {
	if len(arguments) < 1 {
		fmt.Fprintf(os.Stderr, "Not enough arguments usage: echo <message>\n")
	}
	statement := strings.Join(arguments[0:], " ")
	fmt.Println(statement)
}

func typeFn(arguments []string) {
	if len(arguments) < 1 {
		fmt.Fprintf(os.Stderr, "Not enough arguments usage: type <command>\n")
	}

	_, isBuiltIn := builtInCommands[arguments[0]]
	if isBuiltIn {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", arguments[0])
		return
	}

	path, isPresent := getProgramPath(arguments[0])
	if isPresent {
		fmt.Fprintf(os.Stderr, "%s is %s\n", arguments[0], path)
		return
	}

	fmt.Fprintf(os.Stdout, "%s: not found\n", arguments[0])
}

func pwdFn(arguments []string) {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(dir)
}

func cdFn(arguments []string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if arguments[0] == "~" {
		err = os.Chdir(home)
	} else {
		err = os.Chdir(arguments[0])
	}

	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", arguments[0])
	}
}

// Executes external commands

func execFn(command string, arguments []string) {

	commandPath, isPresent := getProgramPath(command)
	if !isPresent {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		return
	}

	cmd := exec.Command(commandPath, arguments...)
	output, err := cmd.Output()
	fmt.Print(string(output))

	if err != nil {
		fmt.Println(err)
	}
}

// Utils

func getProgramPath(program string) (filePath string, isPresent bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		filePath = filepath.Join(path, program)
		_, err := os.Stat(filePath)
		if err == nil {
			return filePath, true
		}
	}

	return "", false
}
