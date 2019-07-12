//这是一个演示go语言中range的dome，我认为和c++11中的范围for一致
package main

import "fmt"

func main(){
  nums := []int{1,2,3,4}
  sum := 0;
  for _, num := range nums{
    sum += num;
    fmt.Print(num)
    fmt.Print(" ")
  }
  fmt.Println()
  fmt.Println("sum :", sum);
  kvs := map[int]string{1 : "apple", 2 : "banana"};
  for k, v := range kvs{
    fmt.Printf("%d--->%s\n", k, v);
  }
  for i, c := range "go"{
    fmt.Println(i, c);
  }
}
