package main  



func main()  {
   // "123454321"
}

func demo(str string) bool {
    data := []byte(str)
    return check(data)
}

func check(data []byte) bool {
    if len(data) == 0 || len(data) == 1 {
        return true
    }
    return data[0] == data[len(data) - 1] && check(data[1:len(data) - 1])
}

func check2(data []byte) bool {
    if len(data) == 0 || len(data) == 1 {
        return true
    }
    left,right := 0, len(data) - 1
    for left <= right {
        if data[left] == data[right] {
            left++
            right--
        } else {
            return false
        }
    }
    return true
}