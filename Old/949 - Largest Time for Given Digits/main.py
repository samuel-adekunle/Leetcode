from functools import reduce


def largestTimeFromDigits(A: list) -> str:
    first = list(filter(lambda x: x in range(3), A))
    if len(first) == 0:
        return ""
    a = max(first)
    A.remove(a)
    if a == 2:
        second = list(filter(lambda x: x in range(3), A))
    else:
        second = A.copy()
    if len(second) == 0:
        return ""
    b = max(second)
    A.remove(b)
    third = list(filter(lambda x: x in range(6), A))
    if len(third) == 0:
        return ""
    c = max(third)
    A.remove(c)
    d = A[0]
    return f"{a}{b}:{c}{d}"


if __name__ == '__main__':
    A = [1, 3, 3, 4]
    res = largestTimeFromDigits(A)
    print(res)
