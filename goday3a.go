package main
// DAY 3 Part I
import (
        "bufio"
        "fmt"
        "log"
        "os"
        "strconv"
	"math"
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

func ToIntFromStrPos(str string,pos int) int {
	v, err  := strconv.Atoi(str[pos:pos+1])
	if err != nil {
		log.Fatal(err)
	}
	return v
}
func getEpsylonbits(gr []int) []int {
	var epsbits []int
	for _,b :=range gr{
		if b==1 {
		epsbits = append (epsbits,0)
		} else {
		epsbits = append (epsbits,1)
		}
	}
	return epsbits
}
func BitToDec(bits []int) int {
	var retval int
	for d,val := range bits{
		retval += val*int(math.Pow(2,float64(11-d)))
	}
	return retval
}
func main() {
        ssitem := StrsliceFromStdin()
//	initem := ToIntsliceFromStrslice(ssitem)
	var gammabit []int 
	var epsylonbit []int
	for i:=0;i<=len(ssitem[0])-1;i++ {
		tbit:=0
		fbit:=0
		most:=0
		for _,val := range ssitem {
			s := ToIntFromStrPos(val,i)
			if s==1 {
				tbit++
			} else {
				fbit++
			}
		}
		if tbit>fbit {
			most=1
		}  else {
			most=0
		}
		gammabit = append (gammabit,most)
	}
/*
12	11 10   9   8  7  6  5 4 3 2 1
   0    0   0   0   0  0  0  0 0 0 0 0 
2048 1024 512 256 128 64 32 16 8 4 2 1
*/
	epsylonbit = getEpsylonbits(gammabit)
	log.Println(gammabit, BitToDec(gammabit))
	log.Println(epsylonbit, BitToDec(epsylonbit))
	fmt.Println("Day3 Part I answer: ", BitToDec(gammabit)*BitToDec(epsylonbit))
	
	
}

