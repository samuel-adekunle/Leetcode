import os
import io
import sys

NEWLINE = '\n'
DEFAULT_CHUNK_SIZE = 128


# print out the last k lines of a file given the file handler
def tail(filename: str, k: int):
    file_size = os.stat(filename).st_size
    num_lines = 0
    chunk_size = DEFAULT_CHUNK_SIZE
    with open(file=filename, mode="rt") as f:
        f.seek(0, os.SEEK_END)  # seek to the end of the file
        while f.tell() > 0:
            if f.tell() < chunk_size:
                chunk_size = f.tell()
            f.seek(f.tell() - chunk_size)  # move back chunk_size bytes
            bytes = f.read(chunk_size)
            for i, b in enumerate(bytes[::-1]):
                if b == NEWLINE:
                    num_lines += 1
                if num_lines == k:
                    f.seek(f.tell() + chunk_size - i)
                    break
            if num_lines == k:
                break
            f.seek(f.tell() - chunk_size)

        # print all lines from current position
        chunk_size = DEFAULT_CHUNK_SIZE
        while f.tell() < file_size:
            if (delta := (file_size - f.tell())) < chunk_size:
                chunk_size = delta
            bytes = f.read(chunk_size)
            print(bytes)


if __name__ == "__main__":
    filename = sys.argv[1]
    tail(filename, 20)
