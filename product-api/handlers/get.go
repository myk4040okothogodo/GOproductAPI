package handlers

import (
    "net/http"
    "github.com/myk4040okothogodo/Grpc2/product-api/data"
    //protos "github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos"
)


//ListALl handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request){
    p.l.Debug("get all records")
    rw.Header().Add("Content-Type", "application/json")
    
    cur := r.URL.Query().Get("currency")

    prods,err := p.productDB.GetProducts(cur)
    if err != nil {
        rw.WriteHeader(http.StatusInternalServerError)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
      }
    err = data.ToJSON(prods, rw)
    if err != nil {
        //we should never be here but log the error just incase
        p.l.Error("serializing product","error", err)
    }
}

//ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request){
    rw.Header().Add("Content-Type", "application/json")
    id := getProductID(r)

    cur := r.URL.Query().Get("currency")

    p.l.Debug("get record id","id", id)
    prod, err := p.productDB.GetProductByID(id, cur)

    switch err {
    case nil:

    case data.ErrProductNotFound:
        p.l.Error(" fetching product","error", err)
        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    default:
        p.l.Error("fetching product","error", err)
        rw.WriteHeader(http.StatusInternalServerError)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return


    }
  

    err = data.ToJSON(prod, rw)
    if err != nil {
        //we should never be here but log the error just incase
        p.l.Error(" serializing product","error", err)

    }

}




