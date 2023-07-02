n = int(input())
n_list = [1000000]
for i in range(0,n):
    a = int(input())
    n_list.append(a)
for num in n_list:
    print(num)
    for i in range(1,len(n_list)):
        if num < n_list[i]:
            n_list.insert(i-1,num)
            break
print(n_list)
