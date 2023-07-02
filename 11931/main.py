n = int(input())
n_list = []
for i in range(0, n):
    a = int(input())
    n_list.append(a)
new_list = [max(n_list), min(n_list)]
for num in n_list:
    before_num = new_list[0]
    for j in range(0, len(new_list)):
        num_2 = new_list[j]
        if num_2 < num < before_num:
            new_list.insert(j, num)
            break
        before_num = num_2
print(new_list)