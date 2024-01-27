package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sample = `{
    "a": "one",
    "b": "two",
    "c": [
	    {
		"a2": [{"a3": 1}]
	    }
    ],
    "d" : {"e": [{"xxx":1}], "f": [{"xxx":2}]}
}`

func TestProduceKeys(t *testing.T) {
	r := strings.NewReader(sample)
	results := []string{}
	err := produceKeys(r, func(key string) {
		results = append(results, key)
	})
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c", "d", "d.e", "d.f"}, results)
}
