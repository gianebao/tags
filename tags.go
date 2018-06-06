package tags

import (
	"reflect"

	model "github.com/jeevatkm/go-model"
)

// ValueList maps all tagged property with `tg` to its value. Nested structs will have a "." separated key
func ValueList(tg string, v interface{}, ns string, m map[string]interface{}) {
	var (
		val  interface{}
		g    reflect.StructTag
		t    reflect.Kind
		f, _ = model.Fields(v)
	)

	for i := 0; i < len(f); i++ {
		val, _ = model.Get(v, f[i].Name)
		g, _ = model.Tag(v, f[i].Name)

		if g.Get(tg) == "" || g.Get(tg) == "-" {
			continue
		}

		if t, _ = model.Kind(v, f[i].Name); t.String() == "struct" {
			ValueList(tg, val, g.Get(tg)+".", m)
			continue
		}

		m[ns+g.Get(tg)] = val
	}
}
