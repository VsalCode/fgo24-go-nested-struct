package main

import "fmt"

func main(){
	type extraAnotherStruct struct{
		val string
	}

	type anotherStruct struct{
		val extraAnotherStruct
	}

	type otherStruct struct{
		val []anotherStruct
	} 

	type strStruct struct {
		val [][]otherStruct
	}

	str := strStruct{
		val: [][]otherStruct{
			{
				{},
				{},
				{},
				{
					val: []anotherStruct{
						{},
						{
							val: extraAnotherStruct{
								val: "Hello",
							},
						},
					},
				},
			},
		},
	}

	fmt.Printf(str.val[0][3].val[1].val.val)

}