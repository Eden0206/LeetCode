package main

import "fmt"

type PersonInfo struct {
	Id int
	Name string
	Address string
}


func main() {
	personA:=make(map[string][2]PersonInfo)
	personB:=make(map[string]PersonInfo)
	personA["test1"]=[2]PersonInfo{
		{Id:001,Name:"Bob",Address:"Road one"},
		{Id:002,Name:"Jack",Address:"Road two"},
	}
	personB["test2"]=PersonInfo{
		Id:003,Name:"Many",Address:"Road three",
	}

	v,ok:=personA["test1"]
	if !ok{
		fmt.Println("err")
	}
	fmt.Println(v)

	v2,ok:=personB["test2"]
	if !ok{
		fmt.Println("err")
	}
	fmt.Println(v2)


	var m = map[string]PersonInfo{
		"test3": {003, "Peter","Road1",},
		"test4": {004, "Jarry","Road2",},
	}

	for _,v:=range m{
		fmt.Println(v)
	}
}
