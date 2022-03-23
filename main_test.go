package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"read-write-bq-api/bqutils"
	"read-write-bq-api/csvutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveData(t *testing.T) {

	data := bqutils.RetrieveData()
	expec := [][]string{{"rahul", "22", "hisar"},
		{"kamalkant", "26", "nawada"},
		{"nikhil", "29", "nohar"},
		{"shivraj", "35", "gwalior"},
		{"rahul", "22", "hisar"},
		{"kamalkant", "26", "nawada"},
		{"nikhil", "29", "nohar"},
		{"shivraj", "35", "gwalior"},
		{"rahul", "22", "hisar"},
		{"kamalkant", "26", "nawada"},
	}
	//fmt.Println(output)
	assert.Equal(t, data, expec)
}

func TestWriteToCsv(t *testing.T) {

	data := [][]string{{"qwert", "12345"},
		{"asdfg", "6789"},
		{"zxcvb", "12345"},
		{"qwert", "6789"},
		{"asdfg", "12345"},
	}
	csvutils.WriteToCsv(data)
	file, err := os.Open("write.csv")
	if err != nil {
		fmt.Println(err)
	}
	csvdata, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, data, csvdata)
}
