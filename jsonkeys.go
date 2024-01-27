package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	if len(os.Args) <= 1 {
		err = produceKeys(os.Stdin, outputKey)
	} else {
		for _, fn := range os.Args[1:] {
			f, ferr := os.Open(fn)
			if ferr != nil {
				err = fmt.Errorf("opening %s: %w", fn, ferr)
				break
			}

			err = produceKeys(f, outputKey)
			if err != nil {
				err = fmt.Errorf("processing %s: %w", fn, err)
			}
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func outputKey(key string) {
	fmt.Println(key)
}
