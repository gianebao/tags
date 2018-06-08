# tags
--
    import "github.com/gianebao/tags"


## Usage

#### type Tag

```go
type Tag struct {
	Name string
}
```


#### func (Tag) GetValueKey

```go
func (tg Tag) GetValueKey(v interface{}, fields ...string) (string, interface{}, error)
```

#### func (Tag) StuctToValueList

```go
func (tg Tag) StuctToValueList(v interface{}, ns string, m map[string]interface{})
```
StuctToValueList maps all tagged property with `tg` to its value. Nested structs
will have a "." separated key
