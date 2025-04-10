package Leetcode

import (
	"strconv"
	"strings"
)

func GetDaysBetween(a, b string) int {
	parse := func(s string)(y,m,d int){
		ps := strings.Split(s,"-")
		y,_ = strconv.Atoi(ps[0])
		m,_ = strconv.Atoi(ps[1])
		d,_ = strconv.Atoi(ps[2])
	}
	y1,m1,d1 = parse(a)
	y2,m2,d2 = parse(b)
	
	days := func(y,m,d int)(n int){
		for y0:1970; y0<y; y0++{
			n+=365
			if (y0%4==0 && y0%100!=0) || (y0%400==0){
				n++
			}
		}
		md := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		if (y%4==0 && y%100!=0) || (y%400==0){
			md[1]=29
		}
		for i:=0; i<m-1; i++{
			n+=md[i]
		}
		return n+d-1
	}
	diff := days(y1,m1,d1) - days(y2,m2,d2)
	if diff<0 {
		return -n
	}
	return n
}