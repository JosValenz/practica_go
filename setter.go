package main


import "fmt"


type Mascota struct{
	edad int
	nombre, raza string
	sana bool
}

func (m *Mascota) ladrar() {
	fmt.Printf("¡guau! me llamo %s y tengo %d años\n", m.nombre, m.edad)
}

func (m *Mascota) Edad() int{
	return m.edad
}

func (m *Mascota) SetEdad(nuevaEdad int) {
	m.edad = nuevaEdad
}

func main() {
	mascota := Mascota{
		edad: 10,
		nombre: "Guayaba",
		raza: "Tampoco la conozco",
		sana: true,
	}
	mascota.ladrar()

	//modificar edad
	mascota.SetEdad(2)

	mascota.ladrar()
}