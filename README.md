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

Gets a value from JSON object. Can accept multiple keys. Returns `nil` on any incorrect input.

```go
json := byte[](`{
  "Adam": { "age": 9 },
  "John": { "age": 12 }
}`)

if age := yson.Get(json, "Adam", "age"); age != nil {
  fmt.Printf("%s", age)
}
// Output: 9
```

### yson.EachKey

Iterates over JSON keys. Does nothing on any incorrect input (including `nil`).

```go
json := byte[](`{ "Adam": 9, "John": 12 }`)

yson.EachKey(json, func(key []byte) {
  fmt.Printf("%s ", key)
})

// Output: Adam John
```

### yson.EachValue

Iterates over JSON values. Does nothing on any incorrect input (including `nil`).

```go
json := byte[](`{ "Adam": 9, "John": 12 }`)

yson.EachValue(json, func(value []byte) {
  fmt.Printf("%s ", value)
})

// Output: 9 12
```

### yson.EachPair

Iterates over JSON keys and values. Does nothing on any incorrect input (including `nil`).

```go
json := byte[](`{ "Adam": 9, "John": 12 }`)

yson.EachPair(json, func(key []byte, value []byte) {
  fmt.Printf("%s=%s ", key, value)
})

// Output: Adam=9 John=12
```

## License

MIT
