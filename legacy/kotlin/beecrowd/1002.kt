import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val radius: String = sc.next()
    val pi = BigDecimal(3.14159)
    println("A=" + (pi * radius.toBigDecimal().pow(2)).setScale(4, RoundingMode.HALF_EVEN))
}