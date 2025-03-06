package hackerrank

// crie uma função que soma os valores de um array
fun arraySum(list: List<Int> ): Int {
   return list.reduce { acc, item -> acc + item }
}
