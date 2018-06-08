package tags

import (
	"reflect"
	"strings"

	model "github.com/jeevatkm/go-model"
)

type Tag struct {
	Name string
}

// StuctToValueList maps all tagged property with `tg` to its value. Nested structs will have a "." separated key
func (tg Tag) StuctToValueList(v interface{}, ns string, m map[string]interface{}) {
	var (
		val  interface{}
		g    reflect.StructTag
		t    reflect.Kind
		tag  string
		f, _ = model.Fields(v)
	)

	for i := 0; i < len(f); i++ {
		val, _ = model.Get(v, f[i].Name)
		g, _ = model.Tag(v, f[i].Name)
		tag = g.Get(tg.Name)

		if tag == "" || tag == "-" {
			continue
		}

		if t, _ = model.Kind(v, f[i].Name); t.String() == "struct" {
			if ns == "" {
				tg.StuctToValueList(val, tag+".", m)
			} else {
				tg.StuctToValueList(val, ns+"."+tag+".", m)
			}
			continue
		}

		m[ns+tag] = val
	}
}

func (tg Tag) GetValueKey(v interface{}, fields ...string) (string, interface{}, error) {
	var s = []string{}

	for _, i := range fields {
		f, err := model.Get(v, i)

		if err != nil {
			return "", "", err
		}

		t, _ := model.Tag(v, i)

		if t.Get(tg.Name) == "" || t.Get(tg.Name) == "-" {
			return strings.Join(s, "."), f, err
		}

		s = append(s, t.Get(tg.Name))
		v = f
	}

	return strings.Join(s, "."), v, nil
}
