package handlers
import (
    "net/http"
    "github.com/myk4040okothogodo/Grpc2/product-api/data"

)


//Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")

    //fetch the product from the context
    prod := r.Context().Value(KeyProduct{}).(data.Product)
    p.l.Println("[DEBUG] updating record id", prod.ID)

    err := data.UpdateProduct(prod)
    if err == data.ErrProductNotFound {
        p.l.Println("[ERROR] product not found", err )

        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
        return
    }
    // write the no content sucess header
    rw.WriteHeader(http.StatusNoContent)

}
