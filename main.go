package main

import (
	"fmt"
	"read-write-bq-api/bqutils"
	"read-write-bq-api/csvutils"
)

func main(){

  fmt.Println(" golang-poc ")
  data := bqutils.RetrieveData()
  fmt.Println(data)
  csvutils.WriteToCsv(data)
  
}