package main

import (
	"log"

	"github.com/manyids2/gophenslide-manyids2/openslide"
)

func main() {
	g, _ := openslide.Open("/data/slides/Aperio/CMU-1.svs")
	log.Printf("%+v\n", g)
	for _, name := range g.AssociatedImageNames() {
		if len(name) > 0 {
			w, h, _ := g.GetAssociatedImageSize(name)
			img, _ := g.GetAssociatedImage(name)
			log.Println(name, w, h, w*h*4, len(img))
		}
	}
}
