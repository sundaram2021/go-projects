package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("file not provide")
	}
	fileName := args[1]

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close() 

	buf := make([]byte, 1024)
	var letters, lines, chars int 

	inWord := false
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		if n == 0 {
			break
		}

		for i:=0;i<n;i++ {
			if buf[i] == '\n'{
				lines++
			}

			if buf[i] == ' ' || buf[i] == '\n' || buf[i] == '\t'{
				inWord = false
			} else{
				if inWord == true{
					letters++
					inWord = false
				}
			}

			if inWord == true {
				letters++
			}

			chars++
		}
	}

	fmt.Println(lines, letters, chars, " ",file.Name())
}

// $ time go run main.go wc.txt 
// 26 0 646   wc.txt

// real    0m0.385s
// user    0m0.046s
// sys     0m0.045s


// $ time wc wc.txt
//  26 135 646 wc.txt

// real    0m0.109s
// user    0m0.075s
// sys     0m0.030s