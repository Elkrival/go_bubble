/*
for each pair (first, second) in consecutive pairs
if first > second;
swap(first, second)
else
 # Do nothing!
	step 1. We iterate and take the pairs as they form.
	step 2. They are the function arguments
	steap 3. Return swap function of (first, second) or nothing happens
*/
package main
import "fmt"
// How would i write a for loop that will give the collection of pairs
// we add two indexes to the for loop and when we iterate we use the second Index to
// Mark our condition to stop the loop
// N is the length of our iterator
// firstIndex and secondINdex mark the position we are in
// To test our algorithms properly we need to test consistently 
// To do so we will add some fmt printLn functions to our loop

/*
	Optimizations
Sorting if To optimize these functions it will be by adding a second parameter called preveious passes
and to my suspicion we grab the difference between the length and the prevpasses
those become the new indexes. the prevpass argument value is i 

Stop sorting if our list is sorted: 
 we will return a boolean and a return type is outside the arguments parens
 the boolean will either be true or false
*/

func main (){
	numbers := []int{5, 4, 2, 3, 1, 0}
	fmt.Println("UnSorted: ", numbers)
	bubbleSort(numbers)
	fmt.Println("Sorted: ", numbers)
	reverse_main()
}
func sweep(numbers []int, prevPasses int) bool {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false 

	for secondIndex < (N - prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			// If the numbers match this criteria we will run the swap function which will swap the positions of the two
			// second moves left and first movies right
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
func bubbleSort(numbers []int) {
	var N int = len(numbers)
	var i int
	for i = 0; i < N; i++{
		if !sweep(numbers, i){
			return 
		}
	}
}
/*
*/
func reverse_main (){
	numbers := []int{5, 4, 2, 3, 1, 0}
	fmt.Println("UnSorted: ", numbers)
	reverseBubbleSort(numbers)
	fmt.Println("Sorted: ", numbers)
}
func reverseSweep(numbers []int, prevPasses int) bool {
	var N int = len(numbers)
	var firstIndex int = N
	var secondIndex int = N - 1
	var didSwap bool = false 

	for firstIndex > (N + prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			// If the numbers match this criteria we will run the swap function which will swap the positions of the two
			// second moves left and first movies right
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
func reverseBubbleSort(numbers []int) {
	var N int = len(numbers)
	var i int
	for i = N; i > 0; i--{
		if !reverseSweep(numbers, i){
			return 
		}
	}
}
