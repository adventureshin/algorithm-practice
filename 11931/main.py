from heapq import heapify, heappop

n = int(input())
numbers = [-int(input()) for _ in range(n)]
heapify(numbers)
while numbers:
    print(-heappop(numbers))
