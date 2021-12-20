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
                log.Println("ToIntsliceFromStrslice ",v)
                dat = append(dat, v)
                if err != nil {
                        log.Fatal(err)
                }
        }
        return dat
}
/** for day1 part I */
func IncDetect(initem []int) int {
        cnt :=0
        prevValue  := 0
        currValue  := 0
        for ind,n := range initem  {
                currValue=n
                if currValue > prevValue {
                        if ind!=0 {
                                cnt++
                        }
                }
                prevValue=n
        }
	log.Println("IncDetect cnt",cnt)
	return cnt
}

/** for day1 part II */
func WindowThreeLabels (initem []int) []int {

        alen := len(initem)
        var wnitem []int = nil
        for idx,_ :=range initem  {
                var val2,val3 int =0,0
                if idx>=alen-2 {
                        val2=0
                } else {
                        val2=initem[idx+1]
                }
                if idx>=alen-3{
                        val3=0
                }else{
                        val3=initem[idx+2]
                }
                tmpval := initem[idx]+val2+val3
                wnitem= append(wnitem,tmpval)
        }
	return wnitem
}

func main() {
        ssitem := StrsliceFromStdin()
	initem := ToIntsliceFromStrslice(ssitem)
	fmt.Println("Day1 PartII answer: ",IncDetect(WindowThreeLabels(initem)))
}
