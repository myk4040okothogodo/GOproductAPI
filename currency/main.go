package main


import (
    "net"
    "os"
    "github.com/hashicorp/go-hclog"
    "google.golang.org/grpc"
    "github.com/myk4040okothogodo/Grpc2/currency/server"
     protos "github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos"
    "google.golang.org/grpc/reflection"
  )

func main() {
   log := hclog.Default()
   
   gs := grpc.NewServer()
   cs :=  server.NewCurrency(log)


   protos.RegisterCurrencyServer(gs, cs)
   
   reflection.Register(gs)

   l, err := net.Listen("tcp", ":9092")
  if err != nil {
     log.Error("Unable to listen", "error", err) 
     os.Exit(1)
  }

   gs.Serve(l)
}
