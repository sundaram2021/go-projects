// a very simple version

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jsonContent := `
	{
  "id": 101,
  "title": "Introduction to Golang",
  "author": {
    "name": "Raman",
    "email": "ram@example.com"
  },
  "published": true,
  "tags": ["golang", "programming", "backend"],
  "content": "Go (or Golang) is an open source programming language designed at Google...",
  "comments": [
    {
      "user": "dev_mike",
      "comment": "Great introduction!",
      "timestamp": ""
    },
    {
      "user": "code_queen",
      "comment": "Looking forward to more posts.",
      "timestamp": ""
    }
  ]
}	
	`
	var i any
	err := json.Unmarshal([]byte(jsonContent), &i)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(i)
}
