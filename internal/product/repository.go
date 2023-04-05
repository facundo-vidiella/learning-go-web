package product

import (
	"errors"

	"github.com/facundo-vidiella/learning-go-web/pkg/store"

	"github.com/facundo-vidiella/learning-go-web/internal/domain"
)



type Repository struct {
}

func (r *Repository) GetProducts() []domain.Product {

	return store.Products
}

func (r *Repository) GetProductById(id int) (domain.Product, error) {

	found := false
	var product domain.Product
	for _, n := range store.Products {
		if n.ID == id {
			found = true
			product = n
		}
	}
	if !found {
		return domain.Product{}, errors.New(`errorMsg: No product found with that id`)
	}

	return product, nil
}

func (r *Repository) GetProductsByPrice(priceGt float64) []domain.Product {

	var validProducts []domain.Product

	for _, n := range store.Products {
		if n.Price > priceGt {
			validProducts = append(validProducts, n)
		}
	}

	return validProducts
}

func (r *Repository) CreateProduct(productBody domain.Product) (string, error) {

	if productBody.ID == 0 {
		productBody.ID = len(store.Products) + 1
	}
	var err error
	for _, n := range store.Products {
		if n.CodeValue == productBody.CodeValue {
			err = errors.New("duplicated error")
			return "There is another product with the same code_value", err
		}
	}

	if !ValidateDate(productBody.Expiration) {
		err = errors.New("date error")
		return "Invalid date format or input", err
	}

	store.Products = append(store.Products, productBody)

	return "Product succesfully created", nil

}

func (r *Repository) UpdateProduct(id int, payload domain.Product) (string, error) {

	var foundProduct domain.Product
	var empty domain.Product
	for _, n := range store.Products {
		if id == n.ID {
			foundProduct = n
			n = payload
		}
	}
	if empty == foundProduct {
		err := errors.New("id error")
		return "No product was found with that id", err
	}

	return "Product succesfully updated", nil

}

func (r *Repository) UpdateProductField(id int, fields any) (string, error) {

	var foundProduct domain.Product
	var empty domain.Product
	for _, n := range store.Products {
		if id == n.ID {
			foundProduct = n

		}
	}
	if empty == foundProduct {
		err := errors.New("id error")
		return "No product was found with that id", err
	}
	return "", nil
}

func (r *Repository) DeleteProduct(id int) error {

	var foundProduct domain.Product
	var empty domain.Product
	for i, n := range store.Products {
		if id == n.ID {
			foundProduct = n
			store.Products[i] = store.Products[len(store.Products)-1]
			store.Products[len(store.Products)-1] = domain.Product{}
			store.Products = store.Products[:len(store.Products)-1]
		}
	}
	if empty == foundProduct {
		err := errors.New("id error")
		return err
	}
	return nil
}
