package main

import "fmt"

// How to create a slice.

func main() {
	// COMPOSITE LITERAL
	//         0, 1, 2, 3, 4, 5, 6, 7, 8
	x := []int{22, 33, 44, 55, 66, 34, 23, 53, 45} // Slice, allows you to group together values of the same time
	fmt.Println(x)
	fmt.Println(x[4])   // get index position of slice
	fmt.Println(len(x)) // get length of slice
	fmt.Println(cap(x))
	fmt.Println(x[1:])  // Display the slice starting with index 1 till the end of the slice
	fmt.Println(x[2:6]) // Display the slice starting at index 2, up to index 6 but not including index 6

	fmt.Println("For Loop with Range:")
	for i, v := range x { // loop over a slice using the range keyword
		fmt.Printf("Index: %v | Value: %v\n", i, v)
	}

	fmt.Println("For Loop without Range:")
	for i := 0; i < len(x); i++ { // loop over a slice without range keyword
		fmt.Printf("Index: %v | Value: %v\n", i, x[i])
	}

	// Append to slices
	x = append(x, 1, 2, 3, 4, 5) // append additional values to slice x
	fmt.Println("This is slice x after appending move values.")
	fmt.Println(x)
	y := []int{195, 198, 145, 450, 600} // create slice y
	x = append(x, y...)                 // Appending slice y into slice x.
	/*
		Note: After slice y use the ... to add all of the values.
		Otherwise it would try to place the slice in the slice which won't
		work cause the slice is not of Type slice.
	*/
	fmt.Println(x)

	// Deleting from slices
	x = append(x[:9], x[14:]...) // will delete everything between index 8 and 14
	fmt.Println(x)

	fmt.Println("This is data from slice m:")
	// using make to set slice min and max length
	m := make([]int, 10, 15) // creates a slice of m with a minimum capacity of 10 and a maximum capacity of 100
	fmt.Println(m)
	fmt.Println(len(m)) // current length of m
	fmt.Println(cap(m)) // maximum capacity of m

	m[0] = 10  // set 0 index to 10
	m[9] = 100 // set 9 index to 100, 10th position
	fmt.Println(m)
	fmt.Println(len(m)) // current length of m
	fmt.Println(cap(m)) // maximum capacity of m

	m = append(m, 110, 120, 130, 140, 150) // fill slice to max capacity of 15
	fmt.Println(m)
	fmt.Println(len(m)) // current length of m
	fmt.Println(cap(m)) // maximum capacity of m

	m = append(m, 160) // once the slice is past max capacity, it doubles capacity
	fmt.Println(m)
	fmt.Println(len(m)) // current length of m
	fmt.Println(cap(m)) // maximum capacity of m

	m = append(m[:1], m[9:]...) // keep everything before index position 1 and after 8. | capacity does not dynamically reduce once items remove.
	fmt.Println(m)
	fmt.Println(len(m)) // current length of m
	fmt.Println(cap(m)) // maximum capacity of m

	// Multi-dimensional slices
	jb := []string{"James", "Bond", "chocolate", "martini"} // normal slice
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "strawberry", "hazelnut"} // normal slice
	fmt.Println(mp)

	xp := [][]string{jb, mp} // multi-dimensional slice with a Type of []string
	fmt.Println(xp)
}
