package ps_dn05mod

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	temp := studenti[vpisnaStevilka]
	delete(studenti, vpisnaStevilka)
	temp.ocene = append(temp.ocene, ocena)
	studenti[vpisnaStevilka] = temp

	return
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}

	var sum float64 = 0.0
	arrLength := len(student.ocene)
	for i := 0; i < arrLength; i++ {
		sum += float64(student.ocene[i])
	}
	var res float64 = sum / float64(arrLength)
	return res
}

func IzpisRedovalnice(studenti map[string]Student) {
	for key, val := range studenti {
		fmt.Printf("%s - %s %s: ", key, val.ime, val.priimek)
		fmt.Println(studenti[key].ocene)
	}

	return
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for key, val := range studenti {
		var avg float64 = povprecje(studenti, key)

		if avg >= 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.2f -> Odličen študent!\n", val.ime, val.priimek, avg)
		} else if 6.0 <= avg && avg < 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.2f -> Povprečen študent!\n", val.ime, val.priimek, avg)
		} else if avg < 6.0 {
			fmt.Printf("%s %s: povprečna ocena 0.0 -> Neuspešen študent!\n", val.ime, val.priimek)
		}
	}

	return
}

func CreateStudent(name string, lastname string, grades []int) Student {
	t := Student{ime: name, priimek: lastname, ocene: grades}
	return t
}
