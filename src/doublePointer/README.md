# 双指针 用法

## 概述

分为两种 `快慢指针` 和 `左右指针` 两种

### 快慢指针

快慢指针一般都初始化指向链表的头结点 head，前进时快指针 fast 在前，慢指针 slow 在后，巧妙解决一些链表中的问题。

示例:

[判定链表中是否含有环](https://github.com/frlute/leetcode/blob/main/src/141.linked-list-cycle.go)

[已知链表中含有环，返回这个环的起始位置](https://github.com/frlute/leetcode/blob/main/src/142.linked-list-cycle-ii.go)

[寻找链表的中点](https://github.com/frlute/leetcode/blob/main/src/876.middle-of-the-linked-list.go)

[删除链表的倒数第N个节点](https://github.com/frlute/leetcode/blob/main/src/19.remove-nth-node-from-end-of-list.go)

### 左右指针

左右指针在数组中实际是指两个索引值，一般初始化为 `left = 0, right = nums.length - 1`。

示例见：

[二分查找](../search/binarySearch/search.go)

*只要数组有序，就应该想到双指针技巧。*

[两数之和](https://github.com/frlute/leetcode/blob/main/src/167.two-sum-ii-input-array-is-sorted.go)

[反转数组](https://github.com/frlute/leetcode/blob/main/src/344.reverse-string.go#L15)

## 进阶-滑动窗口

[最小子字符串](https://github.com/frlute/leetcode/blob/main/src/76.minimum-window-substring.go)

[字符串的排列](https://github.com/frlute/leetcode/blob/main/src/567.permutation-in-string.go)

[找所有字母异位词](https://github.com/frlute/leetcode/blob/main/src/438.find-all-anagrams-in-a-string.go)

[最长无重复子串](https://github.com/frlute/leetcode/blob/main/src/3.longest-substring-without-repeating-characters.go)

### 参考

[滑动窗口技巧](https://labuladong.github.io/algo/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%8A%80%E5%B7%A7%E8%BF%9B%E9%98%B6.html)