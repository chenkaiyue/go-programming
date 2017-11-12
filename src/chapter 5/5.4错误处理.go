package main

import (
	"bufio"
	"os"
	"io"
	"fmt"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	for {
		r,_,err:=in.ReadRune()
		if err == io.EOF{  //结束读取
			break
		}
		if err != nil{
			return fmt.Errorf("read failed %v",err)
		}
	}
}
