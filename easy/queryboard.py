import sys


class Board(object):
    board = [[0 for x in range(256)] for x in range(256)]

    def SetCol(self, col, val):
        for i, x in enumerate(self.board):
            self.board[i][int(col)] = int(val)

    def SetRow(self, row, val):
        for i, x in enumerate(self.board[int(row)]):
            self.board[int(row)][i] = int(val)

    def QueryCol(self, col):
        print sum([x[int(col)] for x in self.board])

    def QueryRow(self, row):
        print sum(self.board[int(row)])


if __name__ == "__main__":
    b = Board()

    f = open(sys.argv[1])
    for l in f:
        args = l.rstrip().split(' ')
        getattr(b, args.pop(0))(*args)
