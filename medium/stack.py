import sys


class Stack(object):
    numbers = []

    def push(self, num):
        self.numbers.append(num)

    def pop(self):
        return self.numbers.pop()


if __name__ == "__main__":
    f = open(sys.argv[1])
    s = Stack()
    for l in f:
        nums = l.rstrip().split(' ')
        to_print = []
        for n in nums:
            s.push(n)
        i = 0
        while(True):
            try:
                n = s.pop()
            except IndexError:
                break
            if i % 2 is 0:
                to_print.append(n)
            i += 1
        print ' '.join(to_print)
