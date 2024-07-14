package main

func main() {
	// --- arrays ---

	// animals := [2]string{}

	// animals[0] = "dog"
	// animals[1] = "cat"

	// fmt.Println(animals)

	// ---- slices -----

	// animals := []string{ // we do not assign a value to the square brackets
	// 	"dog", "cat",
	// }

	// animals = append(animals, "moose") // has extendable size like C++ vectors

	// delete
	// animals = slices.Delete(animals, 1, 2)

	// --- looping ---
	// for i := 0; i < len(animals); i++ {
	// 	fmt.Println(animals[i])
	// }

	// using "range" introduced recently in 1.22
	// for index, value := range animals {
	// 	fmt.Println(index, " -> ", value)
	// }

	// can do it on an integer as well
	// for value := range 3 {
	// 	fmt.Println(value)
	// }

	// --- while loops ----
	// doesn't have a while keyword, basically a for keyword is used only with the conditional statement
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// fmt.Println(animals)

}
