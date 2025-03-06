import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*
import kotlin.math.pow

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val value: String = sc.next()
    val pi: Double = 3.14159

    println("VOLUME = ${BigDecimal((4.0 / 3.0) * pi * value.toDouble().pow(3.0)).setScale(3, RoundingMode.HALF_EVEN)}")
}