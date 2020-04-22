package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"asciicount-test/consts"
)

type CommandGenerate struct{}

func (c *CommandGenerate) Execute() error {
	if err := os.RemoveAll(consts.Directory); err != nil {
		return err
	}

	if err := os.MkdirAll(consts.Directory, os.ModePerm); err != nil {
		return err
	}

	for i := 1; i <= consts.FilesCount; i++ {
		if err := c.fillFile(fmt.Sprintf("%s/%d.txt", consts.Directory, i)); err != nil {
			return err
		}
	}

	return nil
}

func (c *CommandGenerate) fillFile(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	for b := range c.generateSymbols(consts.SymbolsPerFile) {
		if _, err = f.Write([]byte{b}); err != nil {
			return err
		}
	}

	return nil
}

func (c *CommandGenerate) generateSymbols(count int) chan byte {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	ch := make(chan byte)

	n := int32(consts.ASCIIRangeWithSpace)

	go func(ch chan byte) {
		for i := 0; i < count; i++ {
			if i%100 == 0 && i > 0 {
				ch <- '\n'
				continue
			}

			ch <- byte(consts.ASCIIMinWithSpace + rnd.Int31n(n) + 1)
		}

		close(ch)
	}(ch)

	return ch
}
