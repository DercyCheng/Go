package interview

//处理特殊情况：当任一数为0时，直接返回另一个数。
//提取公共因子2：计算shift记录a和b中公共因子2的数量，并将两数除以这些因子。
//确保a为奇数：去除a中剩余的2因子，使其变为奇数。
//主循环处理：
//去除b中的2因子：使b变为奇数。
//交换并相减：保持a ≤ b，通过不断相减缩小问题规模。
//返回结果：将结果a左移shift位恢复公共因子。

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	shift := 0
	for (a&1) == 0 && (b&1) == 0 {
		a >>= 1
		b >>= 1
		shift++
	}

	for (a & 1) == 0 {
		a >>= 1
	}

	for b != 0 {
		for (b & 1) == 0 {
			b >>= 1
		}
		if a > b {
			a, b = b, a
		}
		b -= a
	}

	return a << shift
}
