package main


import (
    "net"
    "os"
    "github.com/hashicorp/go-hclog"
    "google.golang.org/grpc"
    "github.com/myk4040okothogodo/Grpc2/currency/server"
     protos "github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos"
    "google.golang.org/grpc/reflection"
    "github.com/myk4040okothogodo/Grpc2/currency/data"
  )

func main() {
   log := hclog.Default()
   
   rates, err := data.NewRates(log)
   if err != nil {
       log.Error("unable to generate rates,","error", err)
       os.Exit(1)
   }

   gs := grpc.NewServer()
   cs :=  server.NewCurrency(rates, log)


   protos.RegisterCurrencyServer(gs, cs)
   
   reflection.Register(gs)

   l, err := net.Listen("tcp", ":9092")
  if err != nil {
     log.Error("Unable to listen", "error", err) 
     os.Exit(1)
  }

   gs.Serve(l)
}
