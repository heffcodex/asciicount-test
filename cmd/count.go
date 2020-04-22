package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/sync/errgroup"

	"asciicount-test/consts"
	"asciicount-test/helper"
)

type CommandCount struct{}

func (c *CommandCount) Execute() error {
	m, err := c.count(consts.CounterWorkers)
	if err != nil {
		return err
	}

	if m.IsEmpty() {
		return nil
	}

	pl := m.ToPairList()
	sort.Sort(sort.Reverse(pl))

	barPartCapacity := pl[0].V / consts.CounterBarLength
	barFormatAligned := fmt.Sprintf("[%%c] %%%dd [%%s%%s]\n", helper.CountDigits(pl[0].V))

	for _, pair := range pl {
		barLenFilled := pair.V / barPartCapacity
		barLenUnfilled := consts.CounterBarLength - barLenFilled

		fmt.Printf(
			barFormatAligned,
			pair.K,
			pair.V,
			strings.Repeat("|", barLenFilled),
			strings.Repeat(" ", barLenUnfilled),
		)
	}

	return nil
}

func (c *CommandCount) count(workers int) (*helper.CountMap, error) {
	ch := make(chan string)
	m := helper.NewCountMap(consts.ASCIIRange)
	g, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < workers; i++ {
		g.Go(func() error {
			for path := range ch {
				select {
				case <-ctx.Done():
					return nil
				default:
				}

				if err := c.counterFunc(m, path); err != nil {
					return err
				}
			}

			return nil
		})
	}

	errs := make([]string, 0, 2)

	if err := filepath.Walk(consts.Directory, c.walkFilesFunc(ctx, ch)); err != nil {
		errs = append(errs, fmt.Sprintf("files walking aborted: %s", err.Error()))
	}

	close(ch)

	if err := g.Wait(); err != nil {
		errs = append(errs, fmt.Sprintf("symbols counting aborted: %s", err.Error()))
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf(strings.Join(errs, "; "))
	}

	return m, nil
}

func (c *CommandCount) counterFunc(m *helper.CountMap, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			m.AddBytes(line...)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (c *CommandCount) walkFilesFunc(ctx context.Context, ch chan string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ch <- path

		return nil
	}
}
