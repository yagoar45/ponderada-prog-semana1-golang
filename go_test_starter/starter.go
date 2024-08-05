package starter

import (
	"fmt"
	"math"
	"net/http"
)

// SayHello retorna uma mensagem de saudação para o nome fornecido.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

// OddOrEven retorna se o número fornecido é ímpar ou par.
func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}

// Checkhealth escreve uma mensagem de verificação de saúde na resposta HTTP.
func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
}
