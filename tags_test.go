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

func ExampleValueList() {
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

	tags.ValueList("tag", c, "", m)

	fmt.Println(m["family.father"])
	fmt.Println(m["family.mother"])
	fmt.Println(m["family.child"])
	fmt.Println(m["captain"])
	// Output:
	// Daddy
	// Mommy
	// Baby
	// Obvious
}
