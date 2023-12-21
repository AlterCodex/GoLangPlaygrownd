package main
import ("fmt"
		"strings")


func main() {
	var input string
	for{
		fmt.Printf("Insert any sting  to find ian  Ctrl+C to end\n")
		fmt.Scan(&input)
		if(findIan(input)){
			fmt.Printf("Found! \n")
		} else {
			fmt.Printf("Not Found! \n")
		}

		
	}
	
}


func findIan( input string ) bool {
	return strings.Contains(input,"i") && strings.Contains(input,"a")&& strings.Contains(input,"n")
}