package store

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/facundo-vidiella/learning-go-web/internal/domain"
)

type StoreMethods interface {
	InitStore()
	Search()
	Update()
	Delete()
}

type Store struct {
}

var Products []domain.Product

func InitStore() {

	file, err := os.Open("./pkg/store/products.json")
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &Products)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
