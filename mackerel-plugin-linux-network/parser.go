package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func splitKeyValues(s, kvSep string) (key string, vals []string, err error) {
	if kv := strings.SplitN(s, kvSep, 2); len(kv) == 2 {
		key = kv[0]
		vals = strings.Fields(kv[1])
	} else {
		err = fmt.Errorf("No key-value separator found in input string.")
	}

	return
}

func splitKeyValue(s string) (key string, val string, err error) {
	if kv := strings.Fields(s); len(kv) == 2 {
		key = kv[0]
		val = kv[1]
	} else {
		err = fmt.Errorf("Input string should have exactly two space-separated fields.")
	}

	return
}

// Parse a file in the following format (eg. /proc/net/netstat)
//
//   Key: Foo Bar Baz
//   Key: 0 1 42
//   OtherKey: ...
//   OtherKey: ...
//   ...
//
// into a map {"KeyFoo": 0, "KeyBar": 1, "KeyBaz": 42, ...}.
func parseCompact(file io.Reader) map[string]int64 {
	result := make(map[string]int64)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		key, names, err := splitKeyValues(scanner.Text(), ": ")

		if err != nil || !scanner.Scan() {
			break
		}

		key2, values, err := splitKeyValues(scanner.Text(), ": ")

		if err != nil || key != key2 || len(names) != len(values) {
			break
		}

		for i := 0; i < len(names); i++ {
			val, err := strconv.ParseInt(values[i], 10, 64)
			if err == nil {
				result[key + names[i]] = val
			}
		}
	}

	return result
}

func parseTable(file io.Reader) map[string]int64 {
	result := make(map[string]int64)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		key, value, err := splitKeyValue(scanner.Text())

		if err != nil {
			continue
		}

		valueInt, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			result[key] = valueInt
		}
	}

	return result
}
