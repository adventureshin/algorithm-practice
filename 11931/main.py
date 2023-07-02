n = int(input())
numbers = [int(input()) for _ in range(n)]
numbers.sort(reverse=True)
print('\n'.join(str(number) for number in numbers))
