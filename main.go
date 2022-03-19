package main

import (
	"HOMEWORK-2-EZGIUSTUNEL/helper"
	"os"
)

var books, authors []string
var booksInfo []helper.Book

func init() {
	//book list
	books = []string{"Simyaci",
		"Bab-i Esrar",
		"Nar Ağaci",
		"Fareler ve İnsanlar",
		"Kürk Mantolu Madonna",
		"Hayvan Çiftliği",
		"Şeker Portakali",
		"Uçurtma Avcisi",
		"Suç ve Ceza",
		"Serenad",
		"Yeraltindan Notlar",
		"Toprak Ana",
		"Fatih Harbiye",
		"Saatleri Ayarlama Enstitüsü",
		"Acimak",
		"Ateşten Gömlek",
		"Çocukluğum",
		"Aşk",
		"Kuyucakli Yusuf",
		"Arkadaş",
		"Momo",
	}

	//author list
	authors = []string{"Paulo Coelho",
		"Ahmet Ümit",
		"Nazan Bekiroğlu",
		"John Steinback",
		"Sabahattin Ali",
		"George Orwell",
		"Mauro Vasgencelos",
		"Halid Hüseyni",
		"Fyodor Dostoyevski",
		"Zülfü Livaneli",
		"Fyodor Dostoyevski",
		"Cengiz Aytmatov",
		"Peyami Safa",
		"Ahmet Hamdi Tanpinar",
		"Reşat Nuri Güntekin",
		"Halide Edip Adivar",
		"Maksim Gorki",
		"Elif Şafak",
		"Sabahattin Ali",
		"Gorki",
		"Michael Ende",
	}

	setupData()
}

func main() {
	inputList := os.Args
	helper.PerformAction(inputList, booksInfo)
}

// Setup
func setupData() {
	for id, bookName := range books {
		stockNumber := helper.GenerateRandomInt(1000)
		price := helper.GenerateRandomFloat(100)
		pageNumber := helper.GenerateRandomInt(1000)
		stockCode, _ := helper.GenerateRandomCode(10)
		isbn, _ := helper.GenerateRandomCode(10)

		newBook := helper.InitBook(id, stockNumber, pageNumber, price, bookName, stockCode, isbn, authors[id])
		booksInfo = append(booksInfo, newBook)
	}
}
