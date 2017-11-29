package manage

import (
	"fmt"
	"testing"
)

func TestGetRandomNames(t *testing.T) {
	name := GetRandomName()
	fmt.Printf("Random name is %s\n", name)
	// adjNoun := strings.Split(name, "_")
	// if index := sort.(adjectives, adjNoun[0]); index < 0 {
	// 	t.Fail()
	// }
	// if index := sort.SearchStrings(nouns, adjNoun[1]); index < 0 {
	// 	t.Fail()
	// }
}
