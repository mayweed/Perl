// Write a merge sort of...
package main

import (
    "fmt"
    "math/rand"
    "time"
    )

func merge(leftArray,rightArray []int)[]int{
    var output []int
    var i,j int
    //this loop counter not right!!
    for k:=0;k<(len(leftArray)+len(rightArray));k++{
        if leftArray[i] < rightArray[j]{
            output[k]=leftArray[i]
            i+=1
        }else if rightArray[j] < leftArray[i]{
            output[k]=rightArray[j]
            j+=1
        }
    }
    return output
}

//func mergeSort
//middle:=len(array)/2
//fmt.Println(merge(array[:middle],array[middle:]))

func main() {
    //cf go doc example...
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    //generate an array with 10 random positive values
    var array []int
    for i:=0;i<10;i++{
        array=append(array,r.Int())
    }
    fmt.Println(array)
    fmt.Println(merge(array[:5],array[5:]))
}