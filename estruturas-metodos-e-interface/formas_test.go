package estruturas_metodos_e_interface

import "testing"

func TestArea(t *testing.T) {
	testesArea := []struct {
		nome    string
		forma   Forma
		temArea float64
	}{
		{nome: "Retangulo", forma: Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
		{nome: "Circulo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triangulo", forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.temArea)
			}
		})
	}
}
