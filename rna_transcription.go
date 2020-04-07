package strand

import (
	"strings"
)

func zip(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}
	return r
}

// ToRNA transcribes the DNA to RNA
func ToRNA(dna string) string {
	array1 := []string{"G", "C", "T", "A"}
	array2 := []string{"C", "G", "A", "U"}
	return strings.NewReplacer(zip(array1, array2)...).Replace(dna)
}
