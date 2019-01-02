def factorial(num):
    if num == 1:
        return 1
    return num * factorial(num - 1)

def factorial_tail(acc, num):
    if num == 1:
        return acc
    return factorial_tail(acc * num, num-1)

print(factorial(3))
print(factorial_tail(1, 3))
