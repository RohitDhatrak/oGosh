# gosh

## Steps to Run Locally

### Run using golang
1. Install go in your system
2. Run go mod tidy
3. Run the file using go run ./cmd/...

### Run using executable
1. Download the executable file 'gosh'
2. Run the file using ./gosh (for mac/linux) or ./gosh.exe (for windows)

## Built in Commands
1. exit - to exit the shell
2. echo [text] - prints the input to the console
3. type [command] - tells if the command is a built-in or external command
4. pwd - prints the present working directory
5. cd [path] - changes the directory to path provided

## External Commands
You can also run external commands that you have added to your path.
1. ls - lists the files in the current directory
2. cat [file] - prints the content of the file to the console
3. mkdir [dir] - creates a new directory
4. clear - clears the console

## Improvements
1. Support for piping commands (echo "Hello" | wc)
2. I/O redirection (echo "Hello" > file.txt)
3. Environment variable expansion using $ (echo $ENV)
4. Export built-in command (export ENV=value)
5. Background processes using & (echo "Hello" &)