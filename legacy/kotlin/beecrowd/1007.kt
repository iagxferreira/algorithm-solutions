import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val a: String = sc.next()
    val b: String = sc.next()
    val c: String = sc.next()
    val d: String = sc.next()
    println("DIFERENCA = " + BigDecimal((a.toDouble() * b.toDouble() - c.toDouble() * d.toDouble())).setScale(0, RoundingMode.HALF_EVEN))
}