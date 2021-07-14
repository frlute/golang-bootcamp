# golang-bootcamp
golang code bootcamp, 本项目主要是通过 `golang` 来练习一些日常遇到的代码示例。

## Data Structure

- Array
- String
- LinkedList
- Queue
- Stack
- Tree
- Heap

- Graph
## Arithmetic

- 排序
  - 基于比较的排序
    - [快速排序]()
  - 
- 搜索
  - 二分查找
    - [二分查找](./DSA/search/BinarySearch/BinarySearch.go)

### 算法思想与技巧

- 分治思想
- 双指针
  - 快慢指针
  - 左右指针
  - 滑动窗口
- 哨兵

## Concurrency

- [worker pool](concurrency/workerpool/worker_pool.go)
  - [ ] process function add timeout limit
- [pipeline](concurrency/pipeline/pipeline.go)
- [pipeline with cancel](concurrency/pipeline/pipelineWithCancel.go)
- [semaphore](concurrency/semaphore/semaphore.go)

## TO-DO-LIST

- [x] 提交时跑测试用例，或者 commit 时通过 CI 跑测试用例
  - [x] 目前已经增加了 test 和 coverage 命令，方便之后 CI 集成
  - [ ] 增加单元测试覆盖度统计
- [ ] 参考 https://github.com/Workiva/go-datastructures
- [ ] There are many patterns can be implemented using Go, for example Fan-in Fan-out, Generator, Job Queue, Pipeline, Semaphore, Worker Pool etc.

## 参考文档

[data structures](https://en.wikipedia.org/wiki/List_of_data_structures)

[algorithms](https://en.wikipedia.org/wiki/Introduction_to_Algorithms)