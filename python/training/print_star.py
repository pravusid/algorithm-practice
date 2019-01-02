def print_stars(num):
    for i in range(num):
        for _ in range(i):
            print("*", end="")
        print("")


print_stars(10)
