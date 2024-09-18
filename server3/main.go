package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Number struct {
	Num int
	Mu  sync.RWMutex
}

func main() {
	n := Number{
		Num: 0,
		Mu:  sync.RWMutex{},
	}

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
			return
		}

		num1 := r.URL.Query().Get("num1")
		convertedNum, err := strconv.Atoi(num1)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}


		n.Mu.Lock()
		defer n.Mu.Unlock() 
		fmt.Println("Before update, n.Num: ", n.Num)
		n.Num += convertedNum 
		fmt.Println("After update, n.Num: ", n.Num)
		fmt.Fprintf(w, "Updated Sum: %d\n", n.Num)
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
