package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var ErrNotTheEnd = errors.New("found data after the object, possibly this is NDJSON which this doesn't handle")

func produceKeys(r io.Reader, found func(string)) error {
	dec := json.NewDecoder(r)

	tok, err := dec.Token()
	if err != nil {
		return fmt.Errorf("could not read initial token: %w", err)
	}
	// just verify that was what we expected:
	if !checkTokenIsObjectStart(tok) {
		return fmt.Errorf("I expected a top level object, but first token didn't look good: %v", tok)
	}

	return slurpObject(dec, found, "")
}

func slurpObject(dec *json.Decoder, found func(string), prefix string) error {
	expectKey := true
	lastKey := ""

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			// possibly a future error?
			break
		}
		if err != nil {
			return err
		}

		s, isString := tok.(string)
		if isString && expectKey {
			expectKey = false
			lastKey = s
			fullPath := pathJoin(prefix, s)
			found(fullPath)
			continue
		}

		expectKey = true

		if checkTokenIsObjectStart(tok) {
			slurpObject(dec, found, pathJoin(prefix, lastKey))
		} else if checkTokenIsArrayStart(tok) {
			err := skipArray(dec)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return fmt.Errorf("in array: %w", err)
			}
		} else if checkTokenIsObjectEnd(tok) {
			return checkForEnd(dec)
		}
	}

	return nil
}

func checkForEnd(dec *json.Decoder) error {
	if dec.More() {
		return ErrNotTheEnd
	}
	return nil
}

func skipArray(dec *json.Decoder) error {

	arraysOpen := 1
	//	objectsOpen := 1

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		delim, isDelim := tok.(json.Delim)

		// don't care, some value I guess
		if !isDelim {
			continue
		}

		switch delim.String() {
		case "[":
			arraysOpen++
		case "]":
			arraysOpen--
		}

		if arraysOpen == 0 {
			break
		}
	}

	return nil
}

func checkTokenIsObjectStart(t json.Token) bool {
	delim, isDelim := t.(json.Delim)
	return isDelim && delim.String() == "{"
}

func checkTokenIsObjectEnd(t json.Token) bool {
	delim, isDelim := t.(json.Delim)
	return isDelim && delim.String() == "}"
}

func checkTokenIsArrayStart(t json.Token) bool {
	delim, isDelim := t.(json.Delim)
	return isDelim && delim.String() == "["
}

func pathJoin(prefix, key string) string {
	if prefix == "" {
		return key
	}
	return fmt.Sprintf("%s.%s", prefix, key)
}
