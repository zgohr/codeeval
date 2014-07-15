import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        arr = [int(x) for x in l.split(',')]
        max_sum = 0
        positive_sum = 0
        for v in arr:
            # Keep adding integers, disregard if sum goes negative
            positive_sum = max((v + positive_sum, 0))
            max_sum = max((max_sum, positive_sum))
        print max_sum
