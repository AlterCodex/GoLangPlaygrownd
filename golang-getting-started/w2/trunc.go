package main
import "fmt"


func main() {
	var input float64
	for{
		fmt.Printf("Insert any float to be truncate, negative to exit \n")
		fmt.Scan(&input)
		fmt.Printf("the trunctated value for %f is %d \n",input,truncate(input))
		if(input<0){
			break;
		}
		
	}
	
}

func truncate( input float64 ) int {
	return int(input)
}


