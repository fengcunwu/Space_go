//这是一个演示go语言中切片的dome，类似于c++中的vector，
package main

import "fmt"

func print_arr(arr []int){
  fmt.Printf("len = %d, cap = %d, arr = %v\n", len(arr), cap(arr), arr)
}

func main(){
  var arr = []int{1,2,3,4,5,6}
  print_arr(arr)
  s := arr
  print_arr(s)
  s1 := make([]int, 1, 5)
  print_arr(s1)
  s2 := arr[:2]
  print_arr(s2)
  s3 := arr[1:]
  print_arr(s3)
  s4 := arr[2:4]
  print_arr(s4)
  s5 := arr[0:]
  print_arr(s5)
  for i:= 7; i <= 11; i++{
     s = append(s, i)
  }
  print_arr(s)
  s6 := make([]int, len(s), cap(s) * 2)
  copy(s6, s)
  print_arr(s6)
}
