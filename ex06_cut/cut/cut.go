package cut

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type cut struct {
	*flags
}

func NewCut() *cut {
	return &cut{ParseFlags()}
}

func (c *cut) GetFlags() flags {
	return *c.flags
}

func (c *cut) Run() (err error) {
	var input io.Reader

	input = os.Stdin
	if c.input != "" {
		file, err := os.Open(c.input)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err = file.Close()
		}(file)
		input = file
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
	}

	// TODO: Need main functional

	return nil
}
