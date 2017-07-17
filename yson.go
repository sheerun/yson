package yson

import (
	j "encoding/json"

	"github.com/buger/jsonparser"
)

// Iterates though each key and value in json. Does nothing if json is not an object.
//
// Use it for iterating over huge input, as it uses GC better than yson.EachStringKeyAndValue(json)
//
// Memory usage: O(1), Time usage: O(n)
func EachPair(json []byte, fn func(key []byte, value []byte)) {
	jsonparser.ObjectEach(json, func(key []byte, value []byte, t jsonparser.ValueType, offset int) error {
		fn(key, value)
		return nil
	})
}

// Iterates though each key in json. Does nothing if json is not an object.
//
// Memory usage: O(1), Time usage: O(n)
func EachKey(json []byte, fn func(key []byte)) {
	EachPair(json, func(key []byte, value []byte) {
		fn(key)
	})
}

// Iterates though each value in json. Does nothing if json is not an object.
//
// Memory usage: O(1), Time usage: O(n)
func EachValue(json []byte, fn func(value []byte)) {
	EachPair(json, func(key []byte, value []byte) {
		fn(value)
	})
}

// Gets single element from json.
//
// Warning: the behavior for fetching items by numeric key is for now undefined.
//
// Memory usage: O(1), Time usage: O(n)
func Get(json []byte, keys ...string) []byte {
	value, _, _, err := jsonparser.Get(json, keys...)

	if err != nil {
		return nil
	}

	return value
}

// Parses json to go-lang structure or value
//
// Memory usage: O(1), Time usage: O(1)
func Load(json []byte, value interface{}) (err error) {
	return j.Unmarshal(json, value)
}
