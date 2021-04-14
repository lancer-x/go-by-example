package main 

func main() {

}

func splitTest(str string, delimeter []byte) []string {
    ret := make([]string, 0)
    data := []byte(str)

    left := 0
    for i := 0; i < len(data); i++ {
        if checkEqual(data[i:i+len(delimeter)], delimeter)  && left > 0{
            ret = append(ret, data[left:i])
            i += len(delimeter)
            left = i + len(delimeter) - 1
         }
    }
    return ret
}

func checkEqual(target []byte, delimeter []byte) bool {
    bool ret = True
    for i := 0; i < len(delimeter); i++ {
        if target[i] != delimeter[i] {
            ret = false
            break
        }
    }
    return ret
}
