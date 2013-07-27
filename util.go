package main

func Min(a ...int) int {
  min := int(^uint(0) >> 1) // largest possible int
  for _,i := range a {
    if i < min {
      min = i
    }
  }
  return min
}

func Max(a ...int) int {
  max := -(int(^uint(0)>>1) -1) // smallest int
  for _,i := range a {
    if i > max {
      max = i
    }
  }
  return max
}

func Percent(percent, of int) int {
  return of * percent / 100
}
