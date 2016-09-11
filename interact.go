package ask

import (
	"fmt"
	"time"
)

/* Example
package main

import "fmt"
import interact "github.com/planetrobbie/interact"

func main() {
	humeur := interact.Ask("Bonjour ... \n Bienvenue dans l'aventure The_Game \n a:x b:z c:j\nd:x e:z f:j\n", "a", "b", "c", "d", "e", "f")
	temps := interact.Ask("Quel temps fait-il", "beau", "moche")
	jim := interact.Ask("Jim est t-il interesse par Golang ?", "yes", "no")
	fmt.Println(humeur, temps, jim)
}*/

// ask a question (asked) and verify input is one of the provided constraint
func Ask(asked string, constraint ...string) string {
	output := ""

	for !(contains(output, constraint...)) {
		fmt.Println(asked)
		fmt.Scanf("%s", &output)
	}
	return output
}

func Dialog(personnage string, text string, delay time.Duration) {
	time.Sleep(time.Microsecond * delay * 100000)
	fmt.Println(personnage, text)
}

// verify if item is within ...string
func contains(item string, s ...string) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}
