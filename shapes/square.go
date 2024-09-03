package shapes

import (
	"bufio"
	"fmt"
	"go-cli-app/utils"
	"math"
	"time"
)

func Square(reader *bufio.Reader) {
		topA:
		fmt.Print("Input rectangle width : ");
		width,err := utils.InputNumber(reader)
		if err != nil {goto topA}
			
		curcumference := 4 * width  
		area := width  * width
		hypotenuse := math.Sqrt(math.Pow(float64(width),2) * math.Pow(float64(width),2)) 

		fmt.Println("\ncircumference : ", curcumference );
		fmt.Println("area : ", area );
		fmt.Println("diagonal count : ", 2 );
		fmt.Println("diagonal width (hypotenuse) : ", hypotenuse);

		utils.Notif("\nwe will redirect you to main menu soon ....");
		
		time.Sleep(time.Millisecond * 4000)
}