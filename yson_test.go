package yson_test

import (
	"fmt"
	"testing"

	"github.com/sheerun/yson"
)

func jsonWithThousandKeys() []byte {
	result := []byte(`{`)

	for i := 0; i < 1000; i++ {
		result = append(result, []byte(fmt.Sprintf(`"key-%d":"value-%d",`, i, i))...)
	}

	return append(result, []byte(`"key-final":"value-final"}`)...)
}

func ExampleEachKey() {
	json := []byte(`{
		"foo": "bar",
		"fiz": "fuz",
		"1": "2"
	}`)

	yson.EachKey(json, func(key []byte) {
		fmt.Printf("%s ", string(key))
	})
	// Output: foo fiz 1
}

func BenchmarkEachKey(b *testing.B) {
	json := jsonWithThousandKeys()

	for n := 0; n < b.N; n++ {
		yson.EachKey(json, func(key []byte) {
			// nothing
		})
	}
}

func ExampleEachValue() {
	json := []byte(`{
		"foo": "bar",
		"fiz": "fuz",
		"1": "2"
	}`)

	yson.EachValue(json, func(key []byte) {
		fmt.Printf("%s ", string(key))
	})
	// Output: bar fuz 2
}

func BenchmarkEachValue(b *testing.B) {
	json := jsonWithThousandKeys()

	for n := 0; n < b.N; n++ {
		yson.EachValue(json, func(key []byte) {
			// nothing
		})
	}
}

func ExampleEachPair() {
	json := []byte(`{
		"foo": "bar",
		"fiz": "fuz",
		"1": "2"
	}`)

	yson.EachPair(json, func(key []byte, value []byte) {
		fmt.Printf("%s:%s ", key, value)
	})
	// Output: foo:bar fiz:fuz 1:2
}

func BenchmarkEachPair(b *testing.B) {
	json := jsonWithThousandKeys()

	for n := 0; n < b.N; n++ {
		yson.EachPair(json, func(key []byte, value []byte) {
			// nothing
		})
	}
}
