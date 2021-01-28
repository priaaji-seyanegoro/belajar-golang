package basic

import "fmt"

func VariableGo() {
	fmt.Println("hello world")

	/*
	 VALUES
	*/
	fmt.Println("go" + "lang")
	fmt.Println("Penjumlahan 1+1 = ", 1+1)
	fmt.Println("Float 7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)

	/*
		Variables
		- bisa diclare 1 atau lebih variable sekaligus
		- go langsung bs tau tipe variable dari nilai yang disikan
		- variabel yg ga diisi nilai nya akan disikan nilai default
			int, float = 0
			string = '',
			boolean = false
		- syntax := dipake buat initialize variable baru
		- const variabel sama kayak di js gbs di assign lagi constant

	*/

	var a, b int = 1, 2
	fmt.Println(a, b)

	var name = "aji"
	fmt.Println("nama saya : ", name)

	var isBoolean bool
	fmt.Println(isBoolean)

	myHobbies := "Playing Football"
	fmt.Println(myHobbies)

	const n = 500000
	fmt.Println(n)
}
