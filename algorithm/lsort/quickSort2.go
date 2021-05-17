package main 

func main() {

}

func partition1(list []int, low, high int) int {
    //取出基准值，导致当前位置为空  注意理解
    pivot := list[low]

    for low < high {
        for low < high && pivot <= list[high] {
            //碰到比基准大的继续左移，直至碰到一个小的
            high--
        }
        //填补上low位置   此时high位置为空了
        list[low] = list[high]

        for low < high && pivot >= list[low] {
            //碰到比基准小的继续右移，直至碰到一个大的
            low++
        }
        //填不上high的位置  此时low位置为空
        list[high] = list[low]
    }
    //填补上low位置
    list[low] = pivot
    //此时 以low为基准，左右两侧已经有序
    return low
}

func quickSort1(list []int, low, high int)  {
    if high > low {
        pivot := partition1(list, low, high)
        quickSort1(list, low, pivot - 1)
        quickSort1(list, pivot + 1, high)
    }
}