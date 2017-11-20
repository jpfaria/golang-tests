package main

import (
	"github.com/xrash/smetrics"
	"log"
)

func main() {

	a1 := "iPhone 7 32GB Preto Matte Apple - 4G, Tela 4.7, Câmera 12MP + 7MP frontal com Flash Retina, Gravação de vídeos em 4K , iOS 10, MN8X2BR/A"
	b1 := "iPhone 7 32GB Preto Matte Desbloqueado IOS 10 Wi-fi + 4G Câmera 12MP - Apple"

	r1 := smetrics.JaroWinkler(a1, b1, 0.7, 500)

	log.Print(r1)

	a2 := "iPhone 7 32GB Preto Matte Apple - 4G, Tela 4.7, Câmera 12MP + 7MP frontal com Flash Retina, Gravação de vídeos em 4K , iOS 10, MN8X2BR/A"
	b2 := "iPhone 7 128GB Preto Matte Desbloqueado IOS 10 Wi-fi + 4G Câmera 12MP - Apple"

	r2 := smetrics.JaroWinkler(a2, b2, 0.7, 4)

	log.Print(r2)

	a3 := "iPhone Preto 7 32GB Matte"
	b3 := "Preto Matte iPhone 7 32GB"

	r3 := smetrics.JaroWinkler(a3, b3, 0.7, 4)

	log.Print(r3)

	a4 := "TV Samsung"
	b4 := "Preto Matte iPhone 7 128GB"

	r4 := smetrics.JaroWinkler(a4, b4, 0.7, 4)

	log.Print(r4)

	a5 := "Preto Matte iPhone 7 32GB"
	b5 := "Preto Matte iPhone 7 128GB"

	r5 := smetrics.JaroWinkler(a5, b5, 0.7, 4)

	log.Print(r5)

	a6 := "32GB Preto Matte iPhone 7"
	b6 := "Preto Matte iPhone 7 32GB"

	r6 := smetrics.JaroWinkler(a6, b6, 0.7, 500)

	log.Print(r6)

}