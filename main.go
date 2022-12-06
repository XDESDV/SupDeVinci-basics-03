package main

import "fmt"

func main() {
	m := createSpiral(3)
	fmt.Println(m)
}

func createSpiral(n int) [][]int {

	var matrice [][]int
	// on test si
	if n > 0 {
		// Par le calcul
		posLeft := 0
		posRight := n-1
		posTop := 0
		posBottom := n-1

		line := make([]int, n*n)

		values := 1
		// todo explain
		for posLeft < posRight {
			
			for i := posLeft; i <= posRight; i++ {
				line[posTop*n+i] = values
				values++
			}

			posTop++
			for y := posTop; y <= posBottom; y++ {
				line[y*n+posRight] = values
				values++

			} 

			posRight--
			// todo explain
			if posTop == posBottom {
				break
			}

			for x := posRight; x >= posLeft; x-- {
				line[posBottom*n+x] = values
				values++
			}
			
			posBottom--
			for y := posBottom; y >= posTop; y-- {
				line[y*n+posLeft] = values
				values++
			}
		
			posleft++
		}

		line[posTop*n+posLeft] = values


		for i := 0; i <= n-1; i++ {
			low := n * i
			high := n*i + n
			matrice = append(matrice, line[low:high])
		}
	}

	return matrice
}



func createSpiral2(n int) [][]int {

    var matrice [][]int

    // votre code ici
    for i := 1; i <= n; i++ {
        matrice = append(matrice, make([]int, n))
    }

    inc_ligne := 0
    inc_colonne := 1
    pos_ligne := 0
    pos_colonne := 0
    valeur := 1
	
    for valeur <= (n * n) {
        matrice[pos_ligne][pos_colonne] = valeur
        valeur++
        if (pos_ligne + inc_ligne) > (n - 1) {
            inc_colonne = -1
            inc_ligne = 0
        } else {
            if (pos_ligne + inc_ligne) < 0 {
                inc_colonne = 1
                inc_ligne = 0
            } else {
                if (pos_colonne + inc_colonne) > (n - 1) {
                    inc_ligne = 1
                    inc_colonne = 0
                } else {
                    if (pos_colonne + inc_colonne) < 0 {
                        inc_ligne = -1
                        inc_colonne = 0
                    }
                }
            }
        }

        if valeur <= (n * n) {
            if matrice[pos_ligne+inc_ligne][pos_colonne+inc_colonne] > 0 {
                if (inc_colonne) != 0 {
                    inc_ligne = inc_colonne
                    inc_colonne = 0
                } else {
                    inc_colonne = inc_ligne * -1
                    inc_ligne = 0
                }
            }
        }
        pos_ligne += inc_ligne
        pos_colonne += inc_colonne

    }

    return matrice
}

