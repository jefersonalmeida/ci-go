package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Resultado da Sum é inválido: Resultado %d, Esperado %d", total, 30)
	}
}

func TestSumX(t *testing.T) {
	total := SumX(15, 15)

	if total != 45 {
		t.Errorf("Resultado da SumX é inválido: Resultado %d, Esperado %d", total, 45)
	}
}

func TestSub(t *testing.T) {
	total := Sub(15, 10)

	if total != 5 {
		t.Errorf("Resultado da Sub é inválido: Resultado %d, Esperado %d", total, 5)
	}
}

func TestTimes(t *testing.T) {
	total := Times(5, 10)

	if total != 50 {
		t.Errorf("Resultado da Times é inválido: Resultado %d, Esperado %d", total, 50)
	}
}

func TestDiv(t *testing.T) {
	total := Div(10, 2)

	if total != 5 {
		t.Errorf("Resultado da Div é inválido: Resultado %d, Esperado %d", total, 5)
	}
}
