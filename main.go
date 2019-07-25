package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/AllenDowney/OlinPyShop/master/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	wordsArray := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		wordsArray[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range wordsArray {
		fmt.Println(k)
		if i == 10 {
			break
		}
		i++
	}

}
