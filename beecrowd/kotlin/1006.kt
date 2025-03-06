import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val a: String = sc.next()
    val b: String = sc.next()
    val c: String = sc.next()
    println("MEDIA = " + BigDecimal((a.toDouble()* 2 + b.toDouble() * 3 + c.toDouble() * 5)/10 ).setScale(1, RoundingMode.HALF_EVEN))
}