package main

import (
	"fmt"

	"github.com/StefanJakjoski/ps_dn05mod"
)

func main() {
	var studenti = make(map[string]ps_dn05mod.Student)
	ocene1 := []int{10, 9, 10, 9}
	studenti["12345678"] = ps_dn05mod.CreateStudent("asd", "wasd", ocene1)

	fmt.Println("### Testing creation ###")
	ps_dn05mod.IzpisRedovalnice(studenti)
	ps_dn05mod.IzpisiKoncniUspeh(studenti)
	fmt.Println()

	fmt.Println("### Testing dodajOceno ###")
	ps_dn05mod.DodajOceno(studenti, "12345678", 5)
	ps_dn05mod.IzpisRedovalnice(studenti)
	ps_dn05mod.IzpisiKoncniUspeh(studenti)
	fmt.Println()

	fmt.Println("### Testing final ###")
	studenti["23456789"] = ps_dn05mod.CreateStudent("lorem", "ipsum", []int{3, 3, 3, 3})
	studenti["12312334"] = ps_dn05mod.CreateStudent("random", "name", []int{10, 10, 10, 10})
	studenti["98765432"] = ps_dn05mod.CreateStudent("Joze", "Priimek", []int{5, 6, 7, 9})
	ps_dn05mod.IzpisRedovalnice(studenti)
	ps_dn05mod.IzpisiKoncniUspeh(studenti)

	return
}
