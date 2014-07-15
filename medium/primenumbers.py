import sys


f = open(sys.argv[1])
for l in f:
    primes = []
    for num in range(2, int(l)):
        if all(num % div != 0 for div in range(2, num)):
            primes.append(str(num))
    print ','.join(primes)
