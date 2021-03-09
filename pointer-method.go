package main

import (
	"fmt"
	"os"
	init_package "quii-tutorial/init-package"
)

type Biodata struct {
	Nama, Hobi string
	Umur       int
}

func (b *Biodata) Introduce() {
	fmt.Printf("Halo nama saya %s, saya umur %d Tahun, dan Hobi saya %s\n", b.Nama, b.Umur, b.Hobi)
}

func (b *Biodata) TambahPrefix() {
	b.Nama = "Mr. " + b.Nama
}

func main() {
	biodataKanaya := Biodata{
		Nama: "Kanaya",
		Hobi: "Main sama Mommy",
		Umur: 0,
	}

	biodataKanaya.Introduce()

	biodataRyan := Biodata{
		Nama: "Ryan",
		Hobi: "Baca",
		Umur: 30,
	}

	biodataRyan.TambahPrefix()
	biodataRyan.Introduce()

	fmt.Println("INIT RETURN : ", init_package.GetInit())
	fmt.Println("JAVA_PACKAGE", os.Getenv("JAVA_HOME"))
	fmt.Println("JAVA_PACKAGE", os.Getenv("PATH"))

}
