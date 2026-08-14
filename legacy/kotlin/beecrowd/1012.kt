import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*
import kotlin.math.pow


fun main(args: Array<String>) {

    val sc = Scanner(System.`in`)
    val (a, b, c) = sc.nextLine().split(" ".toRegex())
    val pi: Double = 3.14159

    println("TRIANGULO: ${BigDecimal(a.toDouble() * c.toDouble() / 2.0).setScale(3, RoundingMode.HALF_EVEN)}")
    println("CIRCULO: ${BigDecimal( pi * c.toDouble().pow(2)).setScale(3, RoundingMode.HALF_EVEN)}")
    println("TRAPEZIO: ${BigDecimal((a.toDouble() + b.toDouble())*c.toDouble() / 2.0).setScale(3, RoundingMode.HALF_EVEN)}")
    println("QUADRADO: ${BigDecimal(b.toDouble().pow(2.0)).setScale(3, RoundingMode.HALF_EVEN)}")
    println("RETANGULO: ${BigDecimal(a.toDouble() * b.toDouble()).setScale(3, RoundingMode.HALF_EVEN)}")
}