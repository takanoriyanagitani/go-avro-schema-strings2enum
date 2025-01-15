package source

import (
	"bufio"
	"context"
	"io"
	"iter"
	"os"

	. "github.com/takanoriyanagitani/go-avro-schema-strings2enum/util"
)

type Source IO[[]string]

func IterSource(i iter.Seq[string]) IO[[]string] {
	return func(ctx context.Context) ([]string, error) {
		var ret []string
		for s := range i {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			ret = append(ret, s)
		}
		return ret, nil
	}
}

func ScannerToIter(s *bufio.Scanner) iter.Seq[string] {
	return func(yield func(s string) bool) {
		for s.Scan() {
			var line string = s.Text()
			if !yield(line) {
				return
			}
		}
	}
}

var ScannerSource func(*bufio.Scanner) IO[[]string] = Compose(
	ScannerToIter,
	IterSource,
)

var ReaderSource func(io.Reader) IO[[]string] = Compose(
	bufio.NewScanner,
	ScannerSource,
)

var StdinSource IO[[]string] = ReaderSource(os.Stdin)
