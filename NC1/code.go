package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	// write code here
	a := []byte(s)
	b := []byte(t)

	ret := []byte{}

	i := 0
	flag := 0

	for i < len(a) && i < len(b) {
		if res := bToI(a[len(a)-i-1]) + bToI(b[len(b)-i-1]) + flag; res > 9 {
			flag = 1
			ret = append(ret, iToB(res-10))
		} else {
			flag = 0
			ret = append(ret, iToB(res))
		}

		i++
	}

	if len(a) > len(b) {
		for i < len(a) {
			if res := bToI(a[len(a)-i-1]) + flag; res > 9 {
				flag = 1
				ret = append(ret, iToB(res-10))
			} else {
				flag = 0
				ret = append(ret, iToB(res))
			}
			i++
		}
	} else if len(a) < len(b) {
		for i < len(b) {
			if res := bToI(b[len(b)-i-1]) + flag; res > 9 {
				flag = 1
				ret = append(ret, iToB(res-10))
			} else {
				flag = 0
				ret = append(ret, iToB(res))
			}

			i++
		}
	}

	if flag == 1 {
		ret = append(ret, iToB(1))
	}

	reverse(ret)
	return string(ret)
}

func iToB(m int) byte {
	return byte(m + '0')
}

func bToI(n byte) int {
	return int(n - '0')
}

func reverse(arr []byte) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		tmp := arr[l-i-1]
		arr[l-i-1] = arr[i]
		arr[i] = tmp
	}
}
