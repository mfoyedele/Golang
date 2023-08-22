package main

import "fmt"

/*

 */
func main() {

	const GrinningFace = "\U0001f600"
	const FaceBlowingAKiss = "\U0001f618"

	const name = "Khalifa"
	const age = 31

	//const namer, age = "Khalifa", 31

	//fmt.Printf("%s is %d years old.\n", name, age)

	//fmt.Printf("%v is %v years old. \t and the type is %T and %T", name, age, name, age)

	fmt.Println("Hello, world", GrinningFace)
	fmt.Println("I'm MS Khalfa", FaceBlowingAKiss)
	fmt.Printf("It's nice %v being here %v \n", FaceBlowingAKiss, GrinningFace)
	fmt.Println(`	Hello, world,
	I'm Oyedele Musa Funso,
	aka MS Khalifa`)
	fmt.Println("This😊 is looking❤ good✔")
}
