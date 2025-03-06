import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val (_, firstItemCount, firstItemValue) = sc.nextLine().split(" ".toRegex())
    val (_, secondItemCount, secondItemValue) = sc.nextLine().split(" ".toRegex())

    println("VALOR A PAGAR: R\$ ${BigDecimal(firstItemCount.toInt() * firstItemValue.toDouble() + secondItemCount.toInt() * secondItemValue.toDouble()).setScale(2, RoundingMode.HALF_EVEN)}")
}