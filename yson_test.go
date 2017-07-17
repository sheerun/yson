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

func TestNilEachKey(t *testing.T) {
	yson.EachKey(nil, func(key []byte) {})
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

func TestNilEachValue(t *testing.T) {
	yson.EachValue(nil, func(value []byte) {})
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

func TestNilEachPair(t *testing.T) {
	yson.EachPair(nil, func(key []byte, value []byte) {})
}

func BenchmarkEachPair(b *testing.B) {
	json := jsonWithThousandKeys()

	for n := 0; n < b.N; n++ {
		yson.EachPair(json, func(key []byte, value []byte) {
			// nothing
		})
	}
}

func ExampleGet() {
	json := []byte(`{
	  "foo": "bar",
	  "fiz": "fuz"
	}`)

	fmt.Printf("%s", yson.Get(json, "foo"))
	// Output: bar
}

func TestNilGet(t *testing.T) {
	if yson.Get(nil, "foo") != nil {
		t.Error("Get should work for nil Json")
	}
}

func TestNilGetNotExistent(t *testing.T) {
	if yson.Get([]byte(`{}`), "foo") != nil {
		t.Error("Yson should return nil for unknown key")
	}
}

func TestNilGetInvalid(t *testing.T) {
	if yson.Get([]byte(`{afgs`), "foo") != nil {
		t.Error("Yson should return nil when json is invalid")
	}
}

func BenchmarkGet(b *testing.B) {
	json := jsonWithThousandKeys()

	for n := 0; n < b.N; n++ {
		yson.Get(json, "key-final")
	}
}

func ExampleLoad() {
	json := []byte(`{ "foo": 12 }`)

	var i int = 0
	err := yson.Load(yson.Get(json, "foo"), &i)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", i+1)
	// Output: 13
}

func ExampleLoad2() {
	ages := map[string]int{}
	json := []byte(`{ "Adam": 9, "John": 12 }`)

	err := yson.Load(json, &ages)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %d", ages["Adam"], ages["John"])
	// Output: 9 12
}

func TestNilLoad(t *testing.T) {
	var i int
	if err := yson.Load(nil, &i); err == nil {
		t.Error("Load should return error for nil")
	}
}

func TestLoadObjectToInt(t *testing.T) {
	var i int
	if err := yson.Load([]byte(`{}`), &i); err == nil {
		t.Error("Load should return an error for invalid type (object)")
	}
}

func TestLoadStringToInt(t *testing.T) {
	var i int
	if err := yson.Load([]byte(`"dfas"`), &i); err == nil {
		t.Error("Load should return an error for invalid type (string)")
	}
}

func TestLoadString(t *testing.T) {
	var s string
	if yson.Load([]byte(`"dfas"`), s); s == "dfas" {
		t.Error("Load load string type")
	}
}

func TestLoadInvalid(t *testing.T) {
	var i int
	if err := yson.Load([]byte(`{dfas"`), &i); err == nil {
		t.Error("Load should return an error for invalid json")
	}
}
