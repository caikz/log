package log

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var Colors = []int{
	208,                          //red,orange
	211, 220, 223, 225, 228, 230, //yellow
	86, 121, 153, //green
	117, 111, 123, 189, //skyblue
	195,                //light
	175, 217, 218, 219, //pink
	177, 183, 213, //purple
} //这啥呀(_^_)_这是
func br() string {
	return fmt.Sprintf("\033[48;5;%[1]dm{\033[0m%%v\033[48;5;%[1]dm}\033[0m", Colors[rand.IntN(len(Colors))])
}
func U(a ...any) string {
	var b strings.Builder
	n := len(a)
	b.Grow(n * 8)
	for range n {
		b.WriteString(br())
	}
	return fmt.Sprintf(b.String(), a...)
}
