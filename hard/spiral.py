import sys


def spiral(n, m, arr, spiraled, total_length):
    utilized_indexes = []
    # Top
    for x in range(n):
        utilized_indexes.append(x)
        spiraled.append(arr[x])
        if len(spiraled) is total_length:
            return spiraled
    # Right
    for x in range(2, m + 1):
        i = n * x - 1
        utilized_indexes.append(i)
        spiraled.append(arr[i])
        if len(spiraled) is total_length:
            return spiraled
    # Bottom
    for x in range(n * m - 2, n * m - n - 1, -1):
        utilized_indexes.append(x)
        spiraled.append(arr[x])
        if len(spiraled) is total_length:
            return spiraled
    # Left
    for x in range(m - 2, 0, -1):
        i = x * n
        utilized_indexes.append(i)
        spiraled.append(arr[i])
        if len(spiraled) is total_length:
            return spiraled

    remainder_arr = [arr[int(i)]
                     for i in range(total_length) if i not in utilized_indexes]

    return spiral(n - 2, m - 2, remainder_arr, spiraled, total_length)


if __name__ == "__main__":
    for l in open(sys.argv[1]):
        m, n, arr = l.rstrip().split(';')
        m, n = int(m), int(n)
        arr = arr.split(' ')
        print ' '.join(spiral(n, m, arr, [], len(arr)))
