package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ClearScreen(){
		fmt.Print("\033[H\033[2J")
}



func InputNumber(reader *bufio.Reader) (value int, err error ){
		length,_ := reader.ReadString('\n')
		input := strings.TrimSpace(length)		
		newInput,err := strconv.Atoi(input)
		
		if err != nil  {
			fmt.Println("please input number corrently, instead of character");
			return 0,err
		}
		return newInput,err
}
