package handlers
import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/myk4040okothogodo/Grpc2/product-api/data"
    protos "github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos"
)

//KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Products handler for getting and updating products
type Products struct {
    l *log.Logger
    v *data.Validation
    cc protos.CurrencyClient

}

//NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation, cc protos.CurrencyClient) *Products {
    return &Products{l, v, cc}

}

//ErrInavlidProductPath is  an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid path, path should be /products/[id]")


//GenericError is a generic error message returned by a server
type GenericError struct {
    Message string `json:"messsage"`
}

//ValidationError is a collection of vallidation error messages
type ValidationError struct {
    Messages []string `json:"messages"`
}

//getProductID returns the product ID from the URL
//Panics if cannot convert the id into an integer
//this should never happen as the router ensure that this is a valid number

func getProductID(r *http.Request) int {
    // Parse the product id from the url
    vars := mux.Vars(r)

    // convert the id into an integer and return
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        //should never happen
        panic(err)

    }
    return id
}

