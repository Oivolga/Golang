
package main

import ("fmt"
        "io/ioutil"
        "os"
        "bufio"
        "strings")

func main() {
    var name string
	fmt.Scan(&name)
	
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	
	for true {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
//    text = strings.Replace(text, ".", "", -1)
	    if text == "exit" {
	        fmt.Println("Bye!")
	        break
	    } else {
	        word := strings.Split(text, " ")
	        for x := 0; x < len(word); x++ {
	            contain := strings.Contains(string(data), strings.ToLower(word[x]))
	            if contain && word[x] != "a" {
	                fmt.Print(strings.Repeat("*", len(word[x])), " ")
                } else {
                    fmt.Print(word[x], " ")
                }
	        }
	    }
	}
}





