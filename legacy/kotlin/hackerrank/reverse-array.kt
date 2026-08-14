package hackerrank

fun reverseArray(a: Array<Int>): Array<Int> {
    var buffer = mutableListOf<Int>()
    for( i in a.size -1 downTo 0 ){
       buffer.add(a[i])
    }
    return buffer.toTypedArray()
}