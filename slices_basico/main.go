package main

import (
	"fmt"
)

func main()  {
	
	var nums []int

	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}

	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasilia")

	fmt.Println(capitais, len(capitais), cap(capitais))

	cidade := make([]string, 4)
	cidade[0] = "New York"
	cidade[1] = "São Paulo"
	cidade[2] = "Tokio"
	cidade[3] = "Cingapura"

	fmt.Println(cidade, len(cidade), cap(cidade))	

	for indice, cidade := range cidade {
		fmt.Printf("Cidade [%d]  =  %s \n\r", indice, cidade)


	}


}