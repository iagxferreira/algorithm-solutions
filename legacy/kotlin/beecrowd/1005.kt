import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val first: String = sc.next()
    val second: String = sc.next()
    println("MEDIA = " + BigDecimal((first.toDouble()* 3.5 + second.toDouble() * 7.5) / 11).setScale(5, RoundingMode.HALF_EVEN))
}