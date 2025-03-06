import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val name: String = sc.next()
    val wage: String = sc.next()
    val sales: String = sc.next()
    println("TOTAL = R\$ " + BigDecimal(wage.toDouble() + 0.15 * sales.toDouble()).setScale(2, RoundingMode.HALF_EVEN))
}