package args

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var (
	Limit    int = 1_000_000
	PoolSize int
)

func Parse() error {
	flag.IntVar(&PoolSize, "p", runtime.NumCPU(), "Maximum number of parallel workers")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		return nil
	}

	limit, err := parseLimit(args[0])
	if err != nil {
		return fmt.Errorf("invalid limit: %w", err)
	}

	Limit = limit
	return nil
}

func parseLimit(value string) (int, error) {
	value = strings.Replace(value, ",", "", -1)
	value = strings.Replace(value, "_", "", -1)
	return strconv.Atoi(value)
}
