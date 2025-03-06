/*
* Dado um array de inteiros e um target, retorne os indices de dois numeros que adicionados resultem no target.
* */

fun twoSum(nums: IntArray, target: Int): IntArray {
    for (i in 0..nums.size - 1){
        for(j in i + 1.. nums.size -1){
            if(nums[i] + nums[j] == target){
                return intArrayOf(i, j)
            }
        }
    }
    return intArrayOf()
}