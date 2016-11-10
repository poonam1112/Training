package main

import (
    "fmt"
)

//This function take integer value as arguement 
func modifyValue(i int64)  {
    
    a:=make([]int,10)
	s:=make([]int,0)//creating slice/..
	
	//dividing each digit of integer
	for i>0{
		x:=0
		x=int(i%10)
		i=i/10
		//appending last digits of integer to the slice as a first element(append in reverse order)
		s=append(s,x)
		
		a[x]=a[x]+1 //increaing the count for that x index
	}
	// reversing slice, So that we get the original integer values
	for i := len(s)/2-1; i >= 0; i-- {
    opp := len(s)-1-i
    s[i], s[opp] = s[opp], s[i]
    }
    
//fmt.Println(s)// original slice
fmt.Println(a)// slice which contain count of digits
}
func main() {
    var i int64
    // taking user integer input...
	fmt.Println("enter the series of digits")
	fmt.Scanf("%d",&i)
	//function for creating slices...
    modifyValue(i)
}