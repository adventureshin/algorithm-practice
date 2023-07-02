n = int(input())
n_list = list()
for i in range(0,n):
    num = int(input())
    n_list.append(num)
first_n = n_list.pop(0)
p_list = [first_n]
for num in n_list:
    for i in range(0,len(p_list)):
        num_2 = p_list[i]
        if num == num_2:
            break
        elif num > num_2:
            p_list.insert(i,num)
            break
        elif i==len(p_list)-1:
            p_list.append(num)
print(p_list)
