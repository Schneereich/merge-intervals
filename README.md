# merge-intervals
![GitHub](https://img.shields.io/github/license/Schneereich/merge-intervals)
![Travis (.org)](https://img.shields.io/travis/Schneereich/merge-intervals)
![Docker Pulls](https://img.shields.io/docker/pulls/schneereich/merge-intervals)

Golang application accepting integer intervals and returning a sorted list of merged overlapping intervals.

# Overview
The merge function in [interval/interval.go](interval/interval.go) accepts a list of intervals
(i.e. [25,30] [2,19] [14,23] [4,8]) and will produce a sorted list with merged overlapping intervals ([2,23] [25,30]).
If two intervals overlap, a combination of those two is created. 

[Documentation](https://godoc.org/github.com/Schneereich/merge-intervals/interval) of the interval package

Time complexity:
**O(n log n)** due to the sorting algorithm

Space complexity:
**O(1)** as sorting is happening in place and only constant additional space is needed.

To visit source and to get more details on the algorithm, click:
[solution on leetcode.com](https://leetcode.com/problems/merge-intervals/solution/)

## To Do
* Use a linked list for the intervals instead of the array and slice objects.
With very large lists, this should be more efficient. Caution: This seems not to be a standard in Golang.
* Catch more input errors (e.g. bigger start then end value for one interval) 

