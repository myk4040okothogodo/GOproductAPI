package data

import (
    "fmt"
    protos "github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos"
)

//ErrProductNotFound is an error raised when a product cannot be found in th database
var ErrProductNotFound = fmt.Errorf("Product not found")

//Product defines the structure for an API product
type Product struct {
    ID int `json:"id"`   //Unique identifier for the product
    Name string `json:"name" validate:"required"`
    Description string `json: "description"`
    Price float64 `json:"price" validate:"required, gt=0"`
    //The sku for the product
    //required: true
    //pattern: [a-z]+-[a-z]+-[a-z]+
    SKU string `json: "sku" valiate:"sku"`
}

//Product defined a slice of Product
type Products []*Product


type ProductsDB struct {
    currency protos.CurrencyClient
    log hclog.Logger
}

func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
    return &ProductsDB{c, l}

}

//GetProducts returns all products from the database
func (p *ProductsDB) GetProducts( currency string ) (Products,error) {
    if currency == ""{
        return productList, nil
      }
    rate, err := p.getRate(currency)
    if err != nil {
        p.log.Error("unable to get rate for currency :", currency,"error :",err)
        return nil, err
    }
    pr := &Products{}
    for _,p := range ProductList {
        np := *p
        np.Price = np.Price * resp.Rate
        pr = append(pr, &np)
    }
    return pr, nil
}

// GetProductByID returns a single product which matches the id from the database
// if a product is not found this function return a ProductNotFound error

func(p *ProductsDB) GetProductByID(id int) (*Product, error) {
    i := findIndexByProductID(id)
    if id == -1 {
        return nil, ErrProductNotFound
    }

    if currency == ""{
        return productList[i], nil
    }
    rate, err := p.getRate(currency)
    if err != nil {
        p.log.Error("Unable to get rate","currency", currency,"error",err)
        return nil, err

    }

    np := *productList[i]
    np.Price = np.Price * rate
    return &np, nil
}


//updateProduct replaces a product in the database with the given item
//If a product with the given id doesnt exist in the database
//this function returns a ProductNotFound error

func(p *ProductsDB) UpdateProduct(pr Product) error {
    i := findIndexByProductID(pr.ID)
    if i == -1 {
        return ErrProductNotFound
    }

    //update the product in the DB
    productList[i] = &pr
    return nil

}

//Addproduct adds a new product to the database
func (p *ProductsDB) AddProduct(p Product) {
    // get the next id in the sequence
    maxID := productList[len(productList)-1].ID
    pr.ID = maxID + 1
    productList = append(productList,   &pr)

}


//Deleteproduct deletes a product from the database
func (p *productsDB) DeleteProduct(id int) error {
    i := findIndexByProductID(id)
    if i == -1 {
        return ErrProductNotFound
    }

    productList = append(productList[:i], productList[i+1])

    return nil
}


//findIndex finds the index of a product in the databas
//returns -1 when no product can be found
func (p *productsDB)findIndexByProductID(id int) int {
     for i, p := range productList {
          if p.ID == id {
              return i
          }
     }
     return -1

}


func(p *ProductsDB) getRate(destination string)(float64, error){
    rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	resp, err := p.currency.GetRate(context.Background(), rr)
	return resp.Rate, err

}


var productList = []*Product{
    &Product{
        ID:   1,
        Name:   "Latte",
        Description: "Frothy milky coffee",
        Price: 2.45,
        SKU:    "abc323",
    },
    &Product{
        ID:    2,
        Name:   "Esspresso",
        Description: "Short and Strong coffee with milk",
        Price:    1.99,
        SKU: "fjd34",


    },
}

