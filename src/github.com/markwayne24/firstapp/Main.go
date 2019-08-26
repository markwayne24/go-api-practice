package main

// func main() {
// 	var i int = 42
// 	k := 32
// 	fmt.Println(i + k)
// 	fmt.Printf("%v, %T", k, k)

// 	var j float32
// 	j = float32(i)
// 	fmt.Printf("%v, %T\n", j, j)

// 	var n bool = true
// 	fmt.Printf("%v, %T\n", n, n)

// 	var c complex64 = 1 + 2i
// 	fmt.Printf("%v, %T\n", imag(c), imag(c))
// }

// func main() {
// 	s := "this is a string"
// 	//s[2] = "u"
// 	fmt.Printf("%v, %T\n", s, s)

// }

// func main() {
// 	//working with bytes instead of strings
// 	s := "this is a string"
// 	b := []byte(s)
// 	fmt.Printf("%v, %T\n", b, b)

// }

// func main() {
// 	const a int = 14
// 	const b string = "foo"
// 	const c float32 = 3.14
// 	const d bool = true
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%v\n", d)
// }

//_____________________________

// const (
// 	a = iota
// 	b
// 	c
// )

// func main() {
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// }

// _____________________________

// const (
// 	catSpecialist = iota
// 	dogSpecialist
// 	snakeSpecialist
// )

// func main() {
// 	var specialistType int = dogSpecialist
// 	fmt.Printf("%v\n", specialistType == dogSpecialist)
// }

//---------------------------------------

// const (
// 	isAdmin = 1 << iota
// 	isHeadquarters
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmera
// )

// func main() {
// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
// 	fmt.Printf("%b\n", roles)
// 	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
// 	fmt.Printf("Is HQ> %v", isHeadquarters&roles == isHeadquarters)
// }

//--------------------------------------

// ARRAY

// func main() {
// 	//students := [...]int{1,2,3} // or []int{1,2,3}
// 	var students [2]string
// 	students[0] = "Lisa"
// 	students[1] = "Lisa2"
// 	fmt.Printf("Students: %v\n", students)
// 	fmt.Printf("Number of Students: %v", len(students))
// }

//----------------------------------------
// MULTI-DIMENTIONAL ARRAY
// func main() {
// 	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
// 	fmt.Println(identityMatrix)
// }

//------------------------------------

// func main(){
// 	a := []int{1,2,3,4,5,6,7,8,9,10}
// 	b := a[:] //slice of a ll elements
// 	c := a[3:] // slice from 4th element to end
// 	d := a[3:6] // slice first 6 elements
// 	e := a[3:6] // sliece the 4th, 5th, and 6th elements
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// }

//----------------------------------

//FOR LOOP w/ range

// func main() {
//     // Define sharks variable as a slice of strings
//     sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

//     // For loop that iterates over sharks list and prints each string item
//     for _te, shark := range sharks {
//         fmt.Println(shark)
// 		   fmt.Println(_te)
//     }
// }

//-----------------------------------
// POP the first or last
// func main() {
//     a := []int{1,2,3,4,5}

//     //POP the first
//     b := a[1:]
//     //POP the last
//     c := a[:len(a)-1]
//     fmt.Println(b)
//     fmt.Println(c)
// }

//-----------------------------------
//Append
// POP the first or last
// func main() {
// func main() {
// 	a := []int{}

// 	a = append(a, []int{1,2,3}...)
//	a = append(a, 1,2,3)
// 		or
//  b := []int{1,2,3}
//  a = append(a, b...)
//
// 	fmt.Println(a)
// }
// }

//------------------------

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	//Print all array
// 	fmt.Println(a)
// 	b := append(a[:2], a[3:]...)

// 	//remove center
// 	fmt.Println(b)
// 	fmt.Println(a)
// }

//-------------------
//MAPPING
// func main() {
// 	statePopulations := make(map[string]int)
// 	statePopulations = map[string]int{
// 		"California":   1231231,
// 		"Texas":        2231231,
// 		"Florida":      412312,
// 		"New York":     513123,
// 		"Pennsylvania": 8424324,
// 		"Illinois":     91231083,
// 	}

// 	fmt.Println(statePopulations)

// 	//OR

// statePopulations := map[string]int{
// 	"California":   1231231,
// 	"Texas":        2231231,
// 	"Florida":      412312,
// 	"New York":     513123,
// 	"Pennsylvania": 8424324,
// 	"Illinois":     91231083,
// }

// 	fmt.Println(statePopulations)
//  fmt.Println(statePopulations['Illinois'])
// }

//--------------------------Comma ok synthax-----------------

// func main() {
// 	statePopulations := map[string]int{
// 		"California":   1231231,
// 		"Texas":        2231231,
// 		"Florida":      412312,
// 		"New York":     513123,
// 		"Pennsylvania": 8424324,
// 		"Illinois":     91231083,
// 	}

// 	// Check if it exist
// 	_, ok := statePopulations["ohio"]
// 	fmt.Println(ok)
// }

// func main() {
// 	statePopulations := map[string]int{
// 		"California":   1231231,
// 		"Texas":        2231231,
// 		"Florida":      412312,
// 		"New York":     513123,
// 		"Pennsylvania": 8424324,
// 		"Illinois":     91231083,
// 	}

// 	// Check if it exist
// 	_, ok := statePopulations["ohio"]
// 	fmt.Println(ok)
// 	delete(statePopulations, "illinois")
// }

//---------------------------STRUCT---------------

// type Doctor struct {
// 	number int
// 	actorName string
// 	companions []string
// }

// func main() {
//  	aDoctor := Doctor{
// 		number: 3,
// 		actorName: "Jon Pertwee",
// 		companions: []string {
// 			"Liz Shaw",
// 			"Jo Grant",
// 			"Sarah Jane Smith",
// 		},
// 	}

// 	fmt.Println(aDoctor.companions[1])
//  }

//OR

// func main() {
// 	aDoctor := struct{name string}{name: "John Pertwee"}
//    anotherDoctor := aDoctor
//    anotherDoctor.name = "Tom Baker"

//    fmt.Println(aDoctor)
//    fmt.Println(anotherDoctor)

// }

//----------------------------Refactor Struck----------------------

// type Animal struct {
// 	Name   string
// 	Origin string
// }

// type Bird struct {
// 	SpeedKPH float32
// 	CanFly   bool
// }

// func main() {
// 	b := Bird{}
// 	b.Name = "Emu"
// 	b.Origin = "Australia"
// 	b.SpeedKPH = 48
// 	b.CanFly = false
// 	fmt.Println(b.Name)

// 	// OR

// 	b := Bird{
// 		Animal:   Animal{Name: "Emu", Origin: "Australia"},
// 		SpeedKPH: 48,
// 		CanFly:   false,
// 	}
// 	fmt.Println(b.Name)
// }

//-------------------------------------------------------------

// import (
// 	"fmt"
// 	"reflect"
// )

// type Animal struct {
// 	Name   string `required max:"100"`
// 	Origin string
// }

// func main(){
// 	t := reflect.TypeOf(Animal{})
// 	field, _ := t.FieldByName("Name")
// 	fmt.Println(field.Tag)
// }

// ---------------------------------------------

//------------------------If statement---------------

// func main(){
// 	statePopulations := map[string]int{
// 		"California":   1231231,
// 		"Texas":        2231231,
// 		"Florida":      412312,
// 		"New York":     513123,
// 		"Pennsylvania": 8424324,
// 		"Illinois":     91231083,
// 	}

// 	if pop, ok :=statePopulations["Florida"]; ok {
// 		fmt.Println(pop)
// 	}
// }

//--------------------------------------------
