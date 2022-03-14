package handlers

import (
    "net/http"
    "github.com/myk4040okothogodo/Grpc2/product-api/data"
  )

//Create handles POST requests to add new products  
func (p *Products) Create (rw http.ResponseWriter, r *http.Request) {
    //fetch the product from the context
    prod := r.Context().Value(KeyProduct{}).(data.Product)
    p.l.Debug("Inserting product: %#v\n", prod)
    p.productDB.AddProduct(prod)
} 



