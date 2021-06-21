# golang-bootcamp
golang code bootcamp, 本项目主要是通过 `golang` 来练习一些日常遇到的代码示例。

## Data Structure


## Arithmetic

- 排序
  - [快速排序]()
  - 
- 搜索
  - 二分查找
## Concurrency

- [worker pool](concurrency/workerpool/worker_pool.go)
  - [ ] process function add timeout limit
- [pipeline](concurrency/pipeline/pipeline.go)
- [pipeline with cancel](concurrency/pipeline/pipelineWithCancel.go)
- [semaphore](concurrency/semaphore/semaphore.go)

## 算法思想等总结

### 已应用

- 双指针
  - 快慢指针
  - 左右指针
  - 滑动窗口
- 分治思想
  - 哨兵？

TO-DO-LIST

- [ ] 提交时跑测试用例，或者 commit 时通过 CI 跑测试用例
  - 目前已经增加了 test 和 coverage 命令，方便之后 CI 集成
- [ ] 参考 https://github.com/Workiva/go-datastructures
- [ ] There are many patterns can be implemented using Go, for example Fan-in Fan-out, Generator, Job Queue, Pipeline, Semaphore, Worker Pool etc.

## 备用地址

[data structures](https://en.wikipedia.org/wiki/List_of_data_structures)

[algorithms](https://en.wikipedia.org/wiki/Introduction_to_Algorithms)