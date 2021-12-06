from typing import List


class Board:
    def __init__(self):
        self.rows: List[List[int]] = []
        self.cols: List[List[int]] = []

    def add_row(self, row: List[int]):
        self.rows.append(row)

    def add_cols(self):
        cols = []
        for i in range(5):
            col = []
            for row in self.rows:
                col.append(row[i])
            cols.append(col)
        self.cols = cols

    def check_rows(self) -> bool:
        for row in self.rows:
            if len(row) == 0:
                return True
        return False

    def check_cols(self) -> bool:
        for row in self.cols:
            if len(row) == 0:
                return True
        return False


def main():
    with open("./day4.txt", "r") as f:
        lines: List[str] = [line for line in f.readlines()]
    # part 1
    numbers: List[int] = [int(i) for i in lines[0].strip().split(",")]
    boards: List[Board] = []
    b = Board()
    for line in lines[1:]:
        if line == "\n":
            continue
        else:
            b.add_row([int(i) for i in line.strip("\n").split(" ") if i != ""])
        if len(b.rows) == 5:
            b.add_cols()
            boards.append(b)
            b = Board()

    stop = False
    board_num = 0
    win_num = 0
    for number in numbers:
        for num, board in enumerate(boards):
            for row in board.rows:
                while number in row:
                    row.remove(number)
            for col in board.cols:
                while number in col:
                    col.remove(number)
            if board.check_rows() or board.check_cols():
                stop = True
                board_num = num
                win_num = number
                break
        if stop:
            break

    winner = boards[board_num]
    c = 0
    for row in winner.rows:
        c += sum(row)

    print("answer to part 1", c * win_num)


if __name__ == "__main__":
    main()
