package multimedia

import "fmt"

type ContenidoWeb struct {
	Multimedia []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {
	for _, e := range cw.Multimedia {
		e.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

func (i *Imagen) Mostrar() {
	fmt.Println("[", i.Titulo, i.Formato, i.Canales, "]")
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int
}

func (a *Audio) Mostrar() {
	fmt.Println("[", a.Titulo, a.Formato, a.Duracion, "]")
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int
}

func (v *Video) Mostrar() {
	fmt.Println("[", v.Titulo, v.Formato, v.Frames, "]")
}
