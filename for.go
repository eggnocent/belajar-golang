package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke", counter)
	// 	counter++
	// }
	// fmt.Println("selesai")
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("selesa")

	names := []string{"Egi", "Wira", "Tama"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for _, name := range names {
		fmt.Println(name)
	}
}
