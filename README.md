# yson [![GoDoc](https://godoc.org/github.com/sheerun/yson?status.svg)](http://godoc.org/github.com/sheerun/yson)

Zero-allocation, human-friendly JSON library for Go :cake:

## License

MIT






















## <a name="EachKey">func</a> [EachKey](/src/target/yson.go?s=176:222#L1)
## <a name="EachPair">func</a> [EachPair](/src/target/yson.go?s=901:962#L22)
## <a name="EachValue">func</a> [EachValue](/src/target/yson.go?s=480:530#L10)
## <a name="KeysAndValues">func</a> [KeysAndValues](/src/target/yson.go?s=1315:1364#L34)
#### <a name="pkg-examples">Examples</a>
#### <a name="pkg-files">Package files</a>
* [EachKey](#example_EachKey)
* [EachPair](#example_EachPair)
* [EachValue](#example_EachValue)
* [func EachKey(json []byte, fn func(key []byte))](#EachKey)
* [func EachPair(json []byte, fn func(key []byte, value []byte))](#EachPair)
* [func EachValue(json []byte, fn func(value []byte))](#EachValue)
* [func KeysAndValues(json []byte) map[string][]byte](#KeysAndValues)
Iterates though each key and value in json. Does nothing if json is not an object.
Iterates though each key in json. Does nothing if json is not an object.
Iterates though each value in json. Does nothing if json is not an object.
Memory usage: O(1), Time usage: O(n)
Memory usage: O(1), Time usage: O(n)
Memory usage: O(1), Time usage: O(n)
Memory usage: O(1), Time usage: O(n)
Returns an iterator over keys and values of json. Returns empty map if json is not an object.
Use it for iterating over huge input, as it uses GC better than yson.EachStringKeyAndValue(json)
Warning: The order is not preserved while iterating on result.
[yson.go](/src/target/yson.go) 
```
```
```
```
``` go
``` go
``` go
``` go
func EachKey(json []byte, fn func(key []byte))
func EachPair(json []byte, fn func(key []byte, value []byte))
func EachValue(json []byte, fn func(value []byte))
func KeysAndValues(json []byte) map[string][]byte

## License

MIT
