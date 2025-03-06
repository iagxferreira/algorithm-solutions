import java.util.*
import kotlin.math.abs

fun main(args: Array<String>) {

    val sc = Scanner(System.`in`)
    val (a, b, c) = sc.nextLine().split(" ".toRegex())

    val greaterAB = (a.toInt() + b.toInt() +Math.abs(a.toInt()-b.toInt())) / 2
    val greaterOverall = (greaterAB + c.toInt() + abs(greaterAB-c.toInt())) / 2

    println("${greaterOverall} eh o maior")
}