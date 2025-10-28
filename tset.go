package main


import "fmt"
func main() {
	 str := "rida rida rida"
	 split(str)
}

func split(s string) {
	slays := [][]string{}
	str := []string{}
	for _, v := range s {
		if  string(v)!= " "{
       str = append(str, string(v))
		} else {
			slays = append(slays, str)
			str = []string{}
			// fmt.Println(slays)
			// slays = []string{}
		}
		
	}
	slays = append(slays, str)
	fmt.Println(slays)
	
	return 
}