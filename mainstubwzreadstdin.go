package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "strconv"
)

func StrsliceFromStdin() []string {
        var dat []string = nil
        nscan := bufio.NewScanner(os.Stdin)
        for nscan.Scan() {
                dat = append(dat, nscan.Text())
        }
        return dat
}

func ToIntsliceFromStrslice( ss []string ) []int {
        var dat []int = nil
        for _, str := range ss {
                v, err := strconv.Atoi(str)
                fmt.Println(v)
                dat = append(dat, v)
                if err != nil {
                        log.Fatal(err)
                }
        }
        return dat
}
func main() {
        ssitem := StrsliceFromStdin()
        fmt.Println(ssitem)
        isitem := ToIntsliceFromStrslice(ssitem)
        fmt.Println(isitem)
}
