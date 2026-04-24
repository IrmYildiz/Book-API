package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Creating book struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Read   bool   `json:"read"`
}

func main() {

}
