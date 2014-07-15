

if __name__ == '__main__':

    for x in range(1000, 1, -1):
        for n in range(2, x / 2 + 1):
            if x % n is 0:
                break
        else:
            # Prime success! Check palindrome..
            if str(x)[::-1] == str(x):
                print x
                break
            continue
