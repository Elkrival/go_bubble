package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	array := []int{5, 4, 3, 2, 1, 0}
	fmt.Println(array, "A--------------R R A --------------Y")
	sum(array)
	whileMultiplier(array)
	slices()
	var users = slices()
	bubbleSortUsersByLastName(users)
}

func sum(numbers []int) {
	var N = len(numbers)
	sum := 0
	for i := 1; i < N; i++ {
		sum += i
	}
	fmt.Println(sum, "--------Array Sum--------")
}
func whileMultiplier(numbers []int) {
	var N = len(numbers)
	var multiplier = 1
	var n = 0
	fmt.Println(N)
	for n < N {
		if numbers[n] == 0 {

		} else {
			multiplier *= numbers[n]
		}
		n++
	}
	fmt.Println(multiplier, "====================+++^^^^^^^^^^^^")
}

type User struct {
	EMAIL       string `json:"email"`
	F_NAME      string `json:"f_name"`
	L_NAME      string `json:"l_name"`
	CONTACT_NUM string `json:"contact_num"`
}

func slices() []User {
	data, err := os.Open("../_user.json")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()
	var users []User // this is a slice
	dataByteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal([]byte(dataByteValue), &users)
	fmt.Println(users)
	return users
}
func bubbleSortUsersByLastName(users []User) []User {
	var N = len(users)
	var i int
	for i = 0; i < N; i++ {
		ascSweep(users, i)
	}
	fmt.Println("sorted:", users)
	return users
}
func ascSweep(users []User, prevPasses int)bool {
	var N int = len(users)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool =false
	for secondIndex < N {
		var firstUser string = users[firstIndex].L_NAME
		var nextUser string = users[secondIndex].L_NAME
		fmt.Println(firstUser, nextUser)
		if firstUser > nextUser {
			users[firstIndex].L_NAME = nextUser
			users[secondIndex].L_NAME = firstUser
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
