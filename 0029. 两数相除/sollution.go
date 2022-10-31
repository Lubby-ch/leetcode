package main

import (
	"fmt"
	"math"
	"time"
)

func divide(dividend int, divisor int) int {
	var ans int
	var sign = 1
	if (dividend < 0 && divisor >= 0) || (dividend >= 0 && divisor < 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = int(math.Abs(float64(dividend)))
	}
	if divisor < 0 {
		divisor = int(math.Abs(float64(divisor)))
	}
	for dividend >= divisor {
		var count int
		var tmp = divisor
		for {
			if tmp<<1 > dividend {
				break
			}
			tmp <<= 1
			count++
		}
		dividend -= tmp
		ans += int(math.Pow(2, float64(count)))
	}
	if sign == -1 {
		ans = -ans
	}
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}

func main() {
	//fmt.Println(divide(-14, -3))
	a := make(map[uint32][]uint32)
	v := a[0]
	v = append(v, 1)
	a[0] = v
	now := time.Now()
	fmt.Println(now, now.UTC(), len(a[0]), a[0])
}
