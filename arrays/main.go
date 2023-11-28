package main
import "fmt"

func main(){
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length of is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 10)
	fmt.Println(intSlice[3])
	fmt.Printf("The length of is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice2 := []int32{80, 90}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	var intSlice3 []int32 = make([]int32, 2, 3)
	fmt.Println(intSlice3)

	var mMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(mMap)
	var myMap2 = map[string]int32{"Adaaaaaaammm": 45, "Rodddddyyy": 40, "MAttttt": 30}
	fmt.Println(myMap2["Adaaaaaam "])
	var age, ok = myMap2["Adaaaaaaammm"]
	if ok{
		fmt.Printf("The age of Jason is %v", age)
	} else{
		fmt.Println("No such user as Jason")
	}

	for name, age:= range myMap2{
		fmt.Printf("Name : %v\n Age: %v\n", name, age)
	}

	
	for i:=0 ; i<10; i++{
		fmt.Println(i)
	}
}