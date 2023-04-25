package bonus

import (
	"fmt"
	"sort"
)

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	sort.Ints(barLengths)
	fmt.Print(barLengths)
	var altura, torres int
	torres = 1
	altura = 1
	alturaAtual := 1
	if len(barLengths) == 1 {
		altura = 1
		torres = len(barLengths)
	}
	for i := 0; i < len(barLengths)-1; i++ {
		if barLengths[i] == barLengths[i+1] {
			alturaAtual++
		} else {
			torres++
			if alturaAtual > altura {
				altura = alturaAtual
			}
			alturaAtual = 1
		}
	}
	if alturaAtual > altura {
		altura = alturaAtual
	}
	return altura, torres
}
