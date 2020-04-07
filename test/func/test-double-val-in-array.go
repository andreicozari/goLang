package main

import "fmt"

func doubleAt(arr []int, i int) {
	arr[i] *= 2
}

func double(n int) {
	n *= 2
}

/*
The * and & operators§
In Go a pointer is represented using the * (asterisk) character followed by the type of the stored value. In the zero function xPtr is a pointer to an int.
* is also used to “dereference” pointer variables. Dereferencing a pointer gives us access to the value the pointer points to. When we write *xPtr = 0 we are saying “store the int 0 in the memory location xPtr refers to”. If we try xPtr = 0 instead we will get a compiler error because xPtr is not an int it's a *int, which can only be given another *int.
Finally we use the & operator to find the address of a variable. &x returns a *int (pointer to an int) because x is an int. This is what allows us to modify the original variable. &x in main and xPtr in zero refer to the same memory location.

*/
func doublePtr(pointerToNumber *int) {
	//  dereference the pointer like this *n
	//  dereference --> means: give access to the value the pointer points to
	*pointerToNumber *= 2
}

func main() {
	// define arr slice:
	arr := []int{1, 2, 3, 4}

	// pass the slice as a reference to the doubleAt func
	doubleAt(arr, 2)
	fmt.Println(arr) // expect the result  1 2 6 4 which is correct

	n := 2
	double(n)
	fmt.Println(n) // one would believe the value = 4 here, but it's not
	// because GO sends a copy of that number

	// here is the correct doublePtr
	n = 2
	doublePtr(&n)
	fmt.Println(n) // expect here n = 4 , which now is correct

}
