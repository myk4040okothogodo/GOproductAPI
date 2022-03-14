package handlers

import (
    "net/http"
    "github.com/myk4040okothogodo/Grpc2/product-api/data"
)

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Add("Content-Type", "application/json")
    id := getProductID(r)

    p.l.Debug("deleting record", "id", id)

    err := p.productDB.DeleteProduct(id)
    if err == data.ErrProductNotFound {
        p.l.Error("deleting record id doesnt exist")
        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    }
    if err != nil {
        p.l.Error("deleting record","error", err)
        rw.WriteHeader(http.StatusInternalServerError)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    }


    rw.WriteHeader(http.StatusNoContent)
}
