import sys


def sort(arr):
    """
    Insertion sort moves higher numbers up the stack
    """
    for i, v in enumerate(arr):
        j = i
        while j > 0 and arr[j - 1] > arr[j]:
            arr[j - 1], arr[j] = arr[j], arr[j - 1]
            j -= 1
    return arr


for l in open(sys.argv[1]):
    numbers = [float(x) for x in l.split()]
    print ' '.join(["{0:.3f}".format(x) for x in sort(numbers)])
