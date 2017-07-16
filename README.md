# yson [![CircleCI](https://circleci.com/gh/sheerun/yson.svg?style=svg)](https://circleci.com/gh/sheerun/yson) [![Go Report Card](https://goreportcard.com/badge/github.com/sheerun/yson)](https://goreportcard.com/report/github.com/sheerun/yson) [![GoDoc](https://godoc.org/github.com/sheerun/yson?status.svg)](http://godoc.org/github.com/sheerun/yson)

Zero-allocation, human-friendly JSON library for Go :cake:

## Usage

Here's more advanced example:

```go
import "github.com/sheerun/yson"

json := byte[](`{
  "humans": {
    "Adam": {
      "happy": true,
      "age": 9
    },
    "John": {
      "happy": false,
      "age": 12
    }
  }
}`)

yson.EachValue(yson.Get(json, "humans"), func(value []byte) {
  fmt.Printf("%s ", yson.Get(value, "age"))
})

// Output: 9 12
```


## API

Yson functions accept JSON is raw `byte[]` form. Most of them don't allocate memory but just return slices of it.

### yson.Get

> yson.Get(json byte[], keys ...string) []byte

Gets a value from JSON object. Can accept multiple keys. Returns `nil` if any key doesn't exist.

```go
json := byte[](`{
  "Adam": { "age": 9 },
  "John": { "age": 12 }
}`)

fmt.Printf("%s ", yson.Get(json, "Adam", "age"))
// Output: 9
```

### yson.EachKey

```go
json := byte[](`{ "Adam": 9, "John": 12 }`)

yson.EachKey(json, func(key []byte) {
  fmt.Printf("%s ", string(key))
})

// Output: Adam John
```

### yson.EachValue

```go
json := byte[](`{ "Adam": 9, "John": 12 }`)

yson.EachValue(json, func(value []byte) {
  fmt.Printf("%s ", value)
})

// Output: 9 12
```

### yson.EachPair

```go
json := byte[](`{ "Adam": 9, "John": 12 }`)

yson.EachPair(json, func(key []byte, value []byte) {
  fmt.Printf("%s=%s ", key, value)
})

// Output: Adam=9 John=12
```

## License

MIT
