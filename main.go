package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type User struct {
	id       int
	username string
	email    string
	age      int
}

var reader *bufio.Reader
var id int
var users map[int]User

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	cmd.Run()
}

func readLine() string {
	reader = bufio.NewReader(os.Stdin)

	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		// option = strings.TrimSuffix(option, "\r\n")
		return strings.TrimSpace(option)
	}
}

func crearUsuario() {
	fmt.Println("	== Crear usuario ==")

	fmt.Print("Ingresa tu username: \n >> ")
	username := readLine()

	fmt.Print("Ingresa tu email: \n >> ")
	email := readLine()

	fmt.Print("Ingresa tu age: \n >> ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No se puede convertir el string a un entero")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user

	fmt.Println("Usuario creado exitosamente!")
	fmt.Println(users)
}

func listarUsuarios() {
	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}
}

func eliminarUsuario() {
	clearConsole()

	fmt.Print("Ingresa el ID del usuario a eliminar: \n >> ")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No se puede convertir el string a un entero")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	}

	fmt.Println("Usuario eliminado exitosamente!")

}

func main() {

	// * Buscar diferencia entre maps y slices
	var option string //""
	users = make(map[int]User)

	for {
		fmt.Println("1. Crear")
		fmt.Println("2. Listar")
		fmt.Println("3. Actualizar")
		fmt.Println("4. Eliminar")

		fmt.Print("Ingresa una opcion: \n >> ")

		if option == "quit" || option == "q" {
			break
		}

		option = readLine()
		fmt.Println(option)

		switch option {
		case "1", "crear":
			crearUsuario()
		case "2", "listar":
			listarUsuarios()
		case "3", "actualizar":
			listarUsuarios()
		case "4", "eliminar":
			eliminarUsuario()

		}
		clearConsole()
	}

	fmt.Println("\n - Hasta luego, vuelva pronto!")

}

// for ok := true; ok; ok = variable != 10  DO - WHILE

// func miFuncion(valores ...string) VARIADIC Funcition
