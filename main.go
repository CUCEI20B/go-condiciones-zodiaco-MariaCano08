package main

import "fmt"

//Signos zodiacales
//Preguntar si solo es sobre casos reales (no va aponer cosas chukys)
//Preguntar hasta donde es que signo
func main() {
	var d, m int64
	fmt.Scan(&d) //Dia
	fmt.Scan(&m) //Mes
	switch m {
	case 1: //Enero
		if d < 21 {
			fmt.Println("Capricornio")
		} else {
			fmt.Println("Acuario")
		}
		break
	case 2: //Febrero
		if d < 20 {
			fmt.Println("Acuario")
		} else {
			fmt.Println("Piscis")
		}
		break
	case 3: //Marzo
		if d < 21 {
			fmt.Println("Piscis")
		} else {
			fmt.Println("Aries")
		}
		break
	case 4: //Abril
		if d < 21 {
			fmt.Println("Aries")
		} else {
			fmt.Println("Tauro")
		}
		break
	case 5: //Mayo
		if d < 22 {
			fmt.Println("Tauro")
		} else {
			fmt.Println("Geminis")
		}
		break
	case 6: //Junio
		if d < 22 {
			fmt.Println("Geminis")
		} else {
			fmt.Println("Cancer")
		}
		break
	case 7: //Julio
		if d < 24 {
			fmt.Println("Cancer")
		} else {
			fmt.Println("Leo")
		}
		break
	case 8: //Agosto
		if d < 24 {
			fmt.Println("Leo")
		} else {
			fmt.Println("Virgo")
		}
		break
	case 9: //Septiembre
		if d < 24 {
			fmt.Println("Virgo")
		} else {
			fmt.Println("Libra")
		}
		break
	case 10: //Octubre
		if d < 24 {
			fmt.Println("Libra")
		} else {
			fmt.Println("Escorpion")
		}
		break
	case 11: //Noviembre
		if d < 23 {
			fmt.Println("Escorpion")
		} else {
			fmt.Println("Sagitario")
		}
		break
	case 12: //Diciembre
		if d < 22 {
			fmt.Println("Sagitario")
		} else {
			fmt.Println("Capricornio")
		}
		break

	}

}
