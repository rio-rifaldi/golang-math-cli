package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func ClearScreen(){
		fmt.Print("\033[H\033[2J")
}

func InputNumber(reader *bufio.Reader) (value int, err error ){
		length,_ := reader.ReadString('\n')
		input := strings.TrimSpace(length)		
		newInput,err := strconv.Atoi(input)
		
		if err != nil  {
			Error("Please input number corrently, instead of character");
			return 0,err
		}
		return newInput,err
}


var Error = color.New(color.FgHiRed).Add(color.Bold).PrintlnFunc()
var Success = color.New(color.FgHiGreen).Add(color.Bold).PrintlnFunc()
var Notif = color.New(color.FgHiYellow).Add(color.Bold).PrintlnFunc()
