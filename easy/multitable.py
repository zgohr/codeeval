

if __name__ == '__main__':
    count = 0
    for x in range(1, 13):
        count += 1
        print ''.join([str(n * count).rjust(4) for n in range(1, 13)])
