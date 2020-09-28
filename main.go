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
			fmt.Println("capricornio")
		} else {
			fmt.Println("acuario")
		}
		break
	case 2: //Febrero
		if d < 20 {
			fmt.Println("acuario")
		} else {
			fmt.Println("piscis")
		}
		break
	case 3: //Marzo
		if d < 21 {
			fmt.Println("piscis")
		} else {
			fmt.Println("aries")
		}
		break
	case 4: //Abril
		if d < 21 {
			fmt.Println("aries")
		} else {
			fmt.Println("tauro")
		}
		break
	case 5: //Mayo
		if d < 22 {
			fmt.Println("tauro")
		} else {
			fmt.Println("geminis")
		}
		break
	case 6: //Junio
		if d < 22 {
			fmt.Println("geminis")
		} else {
			fmt.Println("cancer")
		}
		break
	case 7: //Julio
		if d < 24 {
			fmt.Println("cancer")
		} else {
			fmt.Println("leo")
		}
		break
	case 8: //Agosto
		if d < 24 {
			fmt.Println("leo")
		} else {
			fmt.Println("v")
		}
		break
	case 9: //Septiembre
		if d < 24 {
			fmt.Println("virgo")
		} else {
			fmt.Println("libra")
		}
		break
	case 10: //Octubre
		if d < 24 {
			fmt.Println("libra")
		} else {
			fmt.Println("escorpion")
		}
		break
	case 11: //Noviembre
		if d < 23 {
			fmt.Println("escorpion")
		} else {
			fmt.Println("sagitario")
		}
		break
	case 12: //Diciembre
		if d < 22 {
			fmt.Println("sagitario")
		} else {
			fmt.Println("capricornio")
		}
		break

	}

}
