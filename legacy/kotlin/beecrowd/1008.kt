import java.math.BigDecimal
import java.math.RoundingMode
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(System.`in`)
    val employeeNumber: String = sc.next()
    val hours: String = sc.next()
    val hourPrice: String = sc.next()
    println("NUMBER = $employeeNumber")
    println("SALARY = U\$ " + BigDecimal(hours.toDouble() * hourPrice.toDouble())
            .setScale(2, RoundingMode.HALF_EVEN))
}