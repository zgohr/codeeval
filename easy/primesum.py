

if __name__ == '__main__':

    total = count = 0
    i = 2
    while(count < 1000):
        for n in range(2, i / 2 + 1):
            if i % n is 0:
                break
        else:
            count += 1
            total += i
        i += 1
    print total
