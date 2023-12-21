package main
import ("fmt"
		)


func main() {
	var sli = make([]int,0,100)
	for{
		input:=0
		fmt.Printf("insert an number ctrl+c to exit \n")
		fmt.Scanf("%d%*c",&input)
		next := input
		for i,v := range sli {
			if(v >= next) {
				sli[i]=next
				next = v
			}
		}
		sli = append(sli,next)
		for _,v := range sli{
			fmt.Printf("%d ",v)
		}
		fmt.Printf("\n")

	}
}
