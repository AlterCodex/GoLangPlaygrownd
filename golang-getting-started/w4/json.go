package main
import ("fmt"
		"encoding/json"
		)

type Person struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

func main() {
	for{
		var p Person
		fmt.Printf("Give the Name of a new Person to marshal (ctrl+c to quit)\n")
		fmt.Scanf("%s",&p.Name)
		fmt.Printf("Give the Address of a new Person to marshal (ctrl+c to quit)\n")
		fmt.Scanf("%s",&p.Address)
		brray, err:=json.Marshal(p)
		//fmt.Println(p)
		if err==nil {
			fmt.Printf("%s \n",brray)
		}


	}
}