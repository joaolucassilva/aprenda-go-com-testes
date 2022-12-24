package arrays_e_slices

import "testing"

func TestSoma(t *testing.T) {
	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("coleção de qualquer número", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})
}