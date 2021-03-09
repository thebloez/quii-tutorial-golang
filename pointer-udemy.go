package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

type Citizen struct {
	Nama   string
	Negara string
}

func main() {
	casePertama()
}

func casePertama() {

	// person1 <- pointer person2
	person1 := Person{
		Id:   1,
		Name: "Ryan",
	}

	// person2 assign ke pointer person1
	person2 := &person1

	person2.Name = "Ryan"
	person2.Name = "Dewi"

	fmt.Println("person1 :", person1)
	fmt.Println("person2", person2)

	fmt.Println("--------------------")

	person3 := &person1
	*person3 = Person{
		Id:   4,
		Name: "Kanaya",
	}

	fmt.Println("person1 :", person1)
	fmt.Println("person2", person2)
	fmt.Println("person2", person3)

	fmt.Println("--------------------")

	person4 := rubahNama(&person1, "Hahaha")
	fmt.Println("person4", person4)
	fmt.Println("person1", person1)

	fmt.Println("--------------------")

	citizen1 := Citizen{
		Nama:   "Ryan",
		Negara: "Indonesia",
	}
	rubahNegara(&citizen1, "Ciamis")
	fmt.Println(citizen1)
}

func rubahNegara(citizen *Citizen, negara string) {
	citizen.Negara = negara
}

func rubahNama(person *Person, nama string) *Person {
	person.Name = nama
	return person
}
