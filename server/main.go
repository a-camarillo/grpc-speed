package main

func main() {
  var (
    address = ":4000"
    svc = &speedFetcherService{}
  )


  makeGRPCServerAndRun(address, svc)

}
