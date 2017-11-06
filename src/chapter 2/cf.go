package main
import(
	"../../pkg/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main(){
	for _,arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s=%s, %s=%s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Print("------------\n")
		fmt.Printf("%v", tempconv.AbsoluteZeroC)
	}
}