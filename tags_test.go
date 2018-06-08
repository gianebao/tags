package tags_test

import (
	"fmt"

	"github.com/gianebao/tags"
)

type Family struct {
	Father string `tag:"father"`
	Mother string `tag:"mother"`
	Child  string `tag:"child"`
}

type Community struct {
	Family  Family `tag:"family"`
	Captain string `tag:"captain"`
	Mayor   string
}

func ExampleTag_StuctToValueList() {
	c := Community{
		Family: Family{
			Father: "Daddy",
			Mother: "Mommy",
			Child:  "Baby",
		},
		Captain: "Obvious",
		Mayor:   "Winner",
	}

	m := map[string]interface{}{}

	(tags.Tag{Name: "tag"}).StuctToValueList(c, "", m)

	fmt.Println("family.father:", m["family.father"])
	fmt.Println("family.mother:", m["family.mother"])
	fmt.Println("family.child:", m["family.child"])
	fmt.Println("captain:", m["captain"])
	// Output:
	// family.father: Daddy
	// family.mother: Mommy
	// family.child: Baby
	// captain: Obvious
}

func ExampleTag_GetValueKey() {
	c := Community{
		Family: Family{
			Father: "Daddy",
			Mother: "Mommy",
			Child:  "Baby",
		},
		Captain: "Obvious",
		Mayor:   "Winner",
	}

	f, v, err := (tags.Tag{Name: "tag"}).GetValueKey(c, "Family", "Mother")
	fmt.Println("tag:", f)
	fmt.Println("val:", v)
	fmt.Println("error:", err)
	// Output:
	// tag: family.mother
	// val: Mommy
	// error: <nil>
}
