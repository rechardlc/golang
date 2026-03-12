package main

import (
	"fmt"
	"sort"
)

// 49. 字母异位词分组
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

/*
思路(计数法)：
1. 字母异位词的特征是各字符出现频率一致。
2. 在 Go 中，固定长度的数组 [26]int 是值类型且可以作为 map 的 key（这与 TS/JS 不同）。
3. 遍历字符串数组，统计每个字符频率并作为 key 存入 map。
4. 最后收集 map 中的所有 values 即可。
*/
func groupAnagrams(strs []string) [][]string {
	// 1. make(map[KeyType]ValueType) 初始化 map。
	// 在 TS 中，Array 是引用类型，不能直接作为 Map 的 Key (会按引用比较)。
	// 在 Go 中，数组是值类型，只要元素可比较，数组就可作为 Key。
	mp := make(map[[26]int][]string)
	for _, s := range strs {
		cnt := [26]int{} // 2. 局部变量初始化为零值，相当于 TS 的 new Array(26).fill(0)
		for i := 0; i < len(s); i++ {
			// 3. s[i] 获取的是 byte (uint8)，在此题 ASCII 场景下比 range 遍历更高效。
			cnt[s[i]-'a']++
		}
		// 4. 这里的 cnt 会被值拷贝到 map 内部作为 Key。
		mp[cnt] = append(mp[cnt], s)
	}
	// 5. 预分配结果切片的容量，类似 TS 中 new Array(mp.size)
	res := make([][]string, 0, len(mp))
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("Input: %v\nOutput: %v\n", strs, groupAnagrams(strs))
	// fmt.Printf("Sort Method Output: %v\n", groupAnagramsBySort(strs))
}

/*
思路 (排序法)：
1. 遍历字符串数组，将每个字符串转换为 byte 切片并按 ASCII 排序。
2. 排序后的字符串作为 map 的 key。
3. 相同的 key 归为一类。
*/
func groupAnagramsBySort(strs []string) [][]string {
	// 1. 在 Go 中 map 的 key 必须是可比较类型，string 是首选。
	mp := make(map[string][]string)

	for _, s := range strs {
		// 2. 将字符串转为 []byte 进行排序
		// 注意：Go 的 string 是不可变的，必须转成 slice 才能操作
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		// 3. 排序后的 []byte 转回 string 作为标识 key
		sortedS := string(b)
		mp[sortedS] = append(mp[sortedS], s)
	}

	res := make([][]string, 0, len(mp))
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
