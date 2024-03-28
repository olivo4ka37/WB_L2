package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// execCmd исполняет команды
func execCmd(argv string, in io.Reader, out io.Writer) error {
	args := strings.Fields(argv)
	if len(args) == 0 {
		return errors.New("parse error")
	}

	command := args[0]
	args = args[1:]

	switch command {
	case "cd":
		switch len(args) {
		case 0:
			dir, err := os.UserHomeDir()
			if err != nil {
				fmt.Fprintln(out, err.Error())
				break
			}
			if err = os.Chdir(dir); err != nil {
				fmt.Fprintln(out, err.Error())
			}
		case 1:
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Fprintln(out, err.Error())
			}
		default:
			fmt.Fprintln(out, "chdir: too many arguments")
		}
	case "pwd":
		if len(args) != 0 {
			fmt.Fprintln(out, "pwd: too many arguments")
			break
		}

		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(out, err.Error())
		} else {
			fmt.Fprintln(out, path)
		}
	case "echo":
		for _, s := range args {
			fmt.Fprint(out, s, " ")
		}
		fmt.Fprintln(out)
	case "kill":
		if len(args) != 1 {
			fmt.Fprintln(out, "kill: invalid arguments")
			break
		}

		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(out, err.Error())
			break
		}

		prc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Fprintln(out, err.Error())
			break
		}

		err = prc.Kill()
		if err != nil {
			fmt.Fprintln(out, err.Error())
		}
	case "ps":
		if len(args) != 0 {
			fmt.Fprintln(out, "ps: too many arguments")
			break
		}

		procs, err := ps.Processes()
		if err != nil {
			fmt.Fprintln(out, err.Error())
			break
		}

		fmt.Fprintln(out, "PID\tCOMMAND")

		for _, p := range procs {
			fmt.Fprintf(out, "%d\t%s\n", p.Pid(), p.Executable())
		}
	case "exit":
		os.Exit(0)
	default:
		cmd := exec.Command(command, args...)
		cmd.Stdin = in
		cmd.Stdout = out

		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(out, err.Error())
		}
	}

	return nil
}

// prompt печатает промпт
func prompt() {
	dir, _ := os.Getwd()
	fmt.Printf("shell %s %% ", filepath.Base(dir))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// в цикле считываем строку
	for prompt(); scanner.Scan(); prompt() {
		// создаем буферы для использования их как пайпов
		var in, out bytes.Buffer

		scan := bufio.NewScanner(os.Stdin)
		in.Write(scan.Bytes())

		// разделяем строку по пайпам на несколько команд
		cmd := strings.Split(scanner.Text(), "|")

		// выполняем команды
		for i := 0; i+1 < len(cmd); i++ {
			if err := execCmd(cmd[i], &in, &out); err != nil {
				fmt.Fprintf(os.Stderr, "shell: %s\n", err.Error())
				break
			}

			in = out
			out.Reset()
		}

		// выполняем последнюю команду, ее вывод должен быть направлен в os.Stdout
		if err := execCmd(cmd[len(cmd)-1], &in, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "shell: %s\n", err.Error())
		}
	}
}
