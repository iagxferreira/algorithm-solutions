package hackerrank

fun diagonalDifference(array: Array<Array<Int>>): Int {
    var sumLeftToRight = 0;
    var sumRightToLeft = 0;
    for(i in 0..array.size){
        for(j in 0..array.size-1){
            // finding the sum of primary diagonal
            if(i == j) {
                sumLeftToRight += array[i][j];
            }
            // finding the sum of secondary diagonal
            if(i + j == array.size - 1){
                sumRightToLeft += array[i][j];
            }
        }
    }
    return Math.abs(sumLeftToRight - sumRightToLeft);
}
