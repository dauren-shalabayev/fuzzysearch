// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func isSubsequence(s string, t string) bool {
	var i int
	var j int

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			fmt.Println("oo", string(s[i]), string(t[j]), i, j)
			i++

		} else {
			fmt.Println("nonono", string(t[j]), j)
		}
		j++
		if i == len(s) {
			return true
		}

	}

	return false
}

func main() {
	fmt.Println(isSubsequence("cra", "cartwheel")) // true
	// fmt.Println(isSubsequence("cwhl", "cartwheel"))      // true
	// fmt.Println(isSubsequence("cwheel", "cartwheel"))    // true
	// fmt.Println(isSubsequence("cartwheel", "cartwheel")) // true
	// fmt.Println(isSubsequence("cwheeel", "cartwheel"))   // false
	// fmt.Println(isSubsequence("lw", "cartwheel"))        // false
}
