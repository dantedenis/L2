package execute

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Shell struct {
	currentPath string
	intro       string
}

func NewShell() (*Shell, error) {
	path, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	return &Shell{
		currentPath: path,
		intro:       filepath.Base(path) + " > ",
	}, nil
}

func (s *Shell) Run() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(s.intro)
		scanner.Scan()
		cmd := strings.Split(scanner.Text(), " ")

		switch cmd[0] {
		case "cd":
			err := s.ChangeDir(cmd[1])
			if err != nil {
				return err
			}
		case "pwd":
			path, err := filepath.Abs(".")
			if err != nil {
				return err
			}
			fmt.Println(path)
		case "echo":
			s.Echo(cmd[1:])
		case "exit":
			fmt.Println("exit")
			return nil
		}
	}
}

func (s *Shell) Echo(arg []string) {
	space := ""
	for _, a := range arg {
		fmt.Print(space, a)
		space = " "
	}
	fmt.Println()
}

func (s *Shell) ChangeDir(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	s.currentPath, err = filepath.Abs(".")
	if err != nil {
		return err
	}
	s.intro = filepath.Base(s.currentPath) + " > "
	return nil
}
