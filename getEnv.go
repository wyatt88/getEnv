// test
package main

import (
	"fmt"
	"os"
	"strings"
)

func GetENV() map[string]string {
	envDict := make(map[string]string)
	envs := os.Environ()
	for i := range envs {
		ss := strings.Split(envs[i], "=")
		envDict[ss[0]] = ss[1]
	}
	return envDict
}
func main() {
	fmt.Print(GetENV()["PATH"])

	// t, err := template.ParseFiles("test.tmpl")
	// if err != nil {
	// 	panic(err)
	// }
	// t.Execute(os.Stdout, d)
	// if err != nil {
	// 	fmt.Print("ERROR")
	// }
}
