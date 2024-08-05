package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

// TestSayHello testa a função SayHello.
// Neste teste, verificamos se a mensagem de saudação está correta para os nomes fornecidos.
func TestSayHello(t *testing.T) {
	// Testando a saudação com o nome "William"
	greeting := starter.SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)

	// Testando a saudação com um nome diferente
	another_greeting := starter.SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}

// TestOddOrEven testa a função OddOrEven.
// Este teste verifica se o número fornecido é corretamente identificado como ímpar ou par.
// Utilizamos subtestes para organizar os casos de teste para números não negativos e negativos.
func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		// Testando números não negativos
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})

	t.Run("Check Negative Numbers", func(t *testing.T) {
		// Testando números negativos
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}

// TestCheckhealth testa a função Checkhealth.
// Este teste verifica se a resposta da verificação de saúde está formatada corretamente e contém os valores esperados.
func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		// Criando uma requisição HTTP simulada
		req := httptest.NewRequest("GET", "http://google.com", nil)
		// Criando um gravador de resposta HTTP simulado
		writer := httptest.NewRecorder()
		// Chamando a função Checkhealth com a requisição e o gravador
		starter.Checkhealth(writer, req)
		// Obtendo a resposta gerada
		response := writer.Result()
		// Lendo o corpo da resposta
		body, err := io.ReadAll(response.Body)

		// Verificando se a mensagem de verificação de saúde está correta
		assert.Equal(t, "health check passed", string(body))
		// Verificando se o status da resposta é 200
		assert.Equal(t, 200, response.StatusCode)
		// Verificando se o cabeçalho Content-Type está correto
		assert.Equal(t, "text/plain; charset=utf-8", response.Header.Get("Content-Type"))
		// Verificando se não houve erros ao ler o corpo da resposta
		assert.Nil(t, err)
	})
}
