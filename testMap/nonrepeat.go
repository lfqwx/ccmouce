package main

import "fmt"

func nonrepeat(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(s) {
		//1.条件1，看是否存在非重复项
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		//条件2，看是否超过字符串最大长度
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		//maxlength = i - start + 1
		lastOccured[ch] = i //这里相当有开始对map进行赋值了,将字符放入map等待下一次判断
		fmt.Printf("值为：%d:\n", lastOccured[ch])
	}
	return maxlength
}

func main() {
	fmt.Println(nonrepeat("aaab"))
	fmt.Println(nonrepeat("岁月流逝"))
}
