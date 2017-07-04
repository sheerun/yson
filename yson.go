package yson

import (
	"github.com/buger/jsonparser"
)

// Iterates though each key in json. Does nothing if json is not an object.
//
// Memory usage: O(1), Time usage: O(n)
func EachKey(json []byte, fn func(key []byte)) {
	jsonparser.ObjectEach(json, func(key []byte, value []byte, t jsonparser.ValueType, offset int) error {
		fn(key)
		return nil
	})
}

// Iterates though each value in json. Does nothing if json is not an object.
//
// Memory usage: O(1), Time usage: O(n)
func EachValue(json []byte, fn func(value []byte)) {
	jsonparser.ObjectEach(json, func(key []byte, value []byte, t jsonparser.ValueType, offset int) error {
		fn(value)
		return nil
	})
}

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
