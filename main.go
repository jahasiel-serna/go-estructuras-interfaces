package main

import (
	"fmt"
	"./multimedia"
)

func agregaImagen() multimedia.Imagen {
	var i multimedia.Imagen
	fmt.Print("\nTítulo: ")
	fmt.Scan(&i.Titulo)
	fmt.Print("Formato: ")
	fmt.Scan(&i.Formato)
	fmt.Print("Canales: ")
	fmt.Scan(&i.Canales)
	return i
}
func agregaAudio() multimedia.Audio {
	var a multimedia.Audio
	fmt.Print("\nTítulo: ")
	fmt.Scan(&a.Titulo)
	fmt.Print("Formato: ")
	fmt.Scan(&a.Formato)
	fmt.Print("Duración: ")
	fmt.Scan(&a.Duracion)
	return a
}
func agregaVideo() multimedia.Video {
	var v multimedia.Video
	fmt.Print("\nTítulo: ")
	fmt.Scan(&v.Titulo)
	fmt.Print("Formato: ")
	fmt.Scan(&v.Formato)
	fmt.Print("Frames: ")
	fmt.Scan(&v.Frames)
	return v
}

func main() {
	op := 0
	content := multimedia.ContenidoWeb{}
	for op != 5 {
		fmt.Print("\n1. Agregar imagen\n2. Agregar audio\n3. Agregar video\n4. Mostrar multimedia\n5. Salir\n > ")
		fmt.Scan(&op)
		switch op {
		case 1:
			t := agregaImagen()
			content.Multimedia = append(content.Multimedia, &t)
		case 2:
			t := agregaAudio()
			content.Multimedia = append(content.Multimedia, &t)
		case 3:
			t := agregaVideo()
			content.Multimedia = append(content.Multimedia, &t)
		case 4:
			content.Mostrar()
		}
	}
}
