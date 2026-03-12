package main // 1. 必写：声明当前文件属于 main 包，只有 main 包的代码才能直接编译成可执行文件

import "fmt" // 2. 导入：类似于 JS 的 import 或 require，导入后必须使用，否则编译报错

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

// 你可以按任意顺序返回答案。



// 示例 1：

// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 示例 2：

// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
// 示例 3：

// 输入：nums = [3,3], target = 6
// 输出：[0,1]

/**
 * twoSum 函数声明
 * - 参数 nums []int: 这是一个 int 类型的切片（Slice），类似于 JS 的 Array [number]
 * - 参数 target int: 目标整数值
 * - 返回值 []int: 返回一个 int 切片，没有返回值也要标注类型
 * 注意：Go 的大括号 { 必须和 func 在同一行，换行会报错！
 */
func twoSum(nums []int, target int) []int {
	// 3. 初始化 Map：类似于 JS 的 const m = new Map() 或 m = {}
	// 注意：Go 的 map 必须用 make 初始化，只声明 var m map[int]int 是 nil，直接赋值会崩溃（panic）
	m := make(map[int]int)

	// 4. 循环遍历：类似于 JS 的 nums.forEach((v, i) => ...) 或 for (const [i, v] of nums.entries())
	// 注意：range 返回的是 (索引, 值)，顺序不能乱。如果不需要索引，要写成 for _, v := range nums
	for i, v := range nums {
		
		// 5. Comma-ok 语法：这是 Go 的特色！类似于 JS 的 if (m.has(target - v)) { let idx = m.get(...) }
		// 作用：尝试从 map 取值，idx 是取出的值，ok 是布尔值（是否存在）
		// 注意：这里的 idx 和 ok 是局部变量，作用域只在这个 if 块里
		if idx, ok := m[target-v]; ok {
			
			// 6. 返回结果：类似于 JS 的 return [idx, i]
			// 注意：[]int{...} 是切片字面量，必须显式声明类型
			return []int{idx, i}
		}

		// 7. 写入 Map：类似于 JS 的 m[v] = i
		// 注意：Go 的 map 是线程不安全的，但在这种单线程逻辑中非常快
		m[v] = i
	}

	// 8. 空返回：类似于 JS 的 return null
	// 注意：切片的空值是 nil，你可以直接返回它
	return nil
}

func main() {
	// 9. 调用函数：注意变量名不要用关键字或保留词
	// 注意：你之前用了 init，在 Go 里 init 是个保留函数名，虽然不报错，但不是好习惯
	result := twoSum([]int{2, 7, 11, 15}, 9)

	// 10. 打印输出：fmt.Println 类似于 console.log
	fmt.Println(result)
}
