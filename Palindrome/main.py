"""Script to test the is_palindrome function."""


# essentially a recursive stack
def is_palindrome_rec(word: str):
    """Returns True if s is a palindrome and False otherwise."""
    if len(word) <= 1:
        return True
    return word[0] == word[-1] and is_palindrome_rec(word[1:-1])


# iterating using essentially two pointers
def is_palindrome_iter(word: str):
    """Returns True if s is a palindrome and False otherwise."""
    for i in range(len(word) // 2):
        if word[i] != word[-i - 1]:
            return False
    return True


if __name__ == "__main__":
    print("Enter a string:")
    string = input()
    if is_palindrome_rec(string) and is_palindrome_iter(string):
        print("The string is a palindrome")
    else:
        print("The string is not a palindrome")
