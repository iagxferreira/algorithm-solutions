package hackerrank

/**
 *  crie uma função que receba dois arrays a e compare seus valores, atribuindo seus resultados a um terceiro array
 *  sendo a posição 0 do array contando pontos para a quantidade de numeros maiores no primeiro
 *  e a segunda posição para contar a quantidade de numeros maiores no segundo array
 *  caso os dois numeros sejam iguais, nenhum deles ganha ponto
*/
fun compareTriplets(aliceOptions: Array<Int>, bobOptions: Array<Int>): Array<Int> {
    val points = arrayOf(0,0)
    loop@ for(i in 0 ..2){
        when {
            aliceOptions[i] > bobOptions[i] -> points[0]++
            aliceOptions[i] < bobOptions[i] -> points[1]++
            else -> continue@loop
        }
    }
    return points
}