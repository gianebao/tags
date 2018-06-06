# tags
--
    import "github.com/gianebao/tags"


## Usage

#### func  ValueList

```go
func ValueList(tg string, v interface{}, ns string, m map[string]interface{})
```
ValueList maps all tagged property with `tg` to its value. Nested structs will
have a "." separated key
