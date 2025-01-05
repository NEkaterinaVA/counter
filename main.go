package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var flagWords = []string{"счёт", "числ", "знач", "табл"}
var counter int

// Небольшое изменение для проверки git config
func main(){
	fmt.Println("Вы что-то хотели?")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil{
		return
	}

	if strings.Contains(text, "?"){
		for _, word := range flagWords{
			if strings.Contains(text, word){
				counter++
				break
			}
		}
	}

	fmt.Println("Количество любопытных человек -", counter)
}