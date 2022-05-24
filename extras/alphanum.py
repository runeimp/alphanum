
__author__ = ['RuneImp']
__version__ = '0.1.0'


import sys

from functools import reduce


# Code pulled from https://gist.github.com/robinhouston/99746cac543e6b3ea61f1d245e9b19cc
# and used as a reference for my Go library. Thank you Robin Houston!


def col_index(col:str) -> int:
	"""Take a spreadsheet-style column specifier, e.g. A, M, ZZ,
	and convert it to a one-based column index: A=1, Z=26, AA=27, etc.
	"""
	if not col or set(col) - set("ABCDEFGHIJKLMNOPQRSTUVWXYZ"):
		raise Exception("Bad column specifier: '" + col + "'")
	return reduce(lambda x,c: 26 * x + ord(c) - 64, col, 0)


def col_letters(index:int) -> str:
	"""Take a one-based column index and convert it to a
	spreadsheet-style column specifier: 1=A, 26=Z, 27=AA, etc.
	"""
	col_spec = ""
	while index:
		index, r = divmod(index - 1, 26)
		col_spec = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"[r] + col_spec
	return col_spec


if __name__ == '__main__':
	for arg in sys.argv[1:]:
		if arg.isalpha():
			arg = arg.upper()
			print(f"{arg} = {col_index(arg)}")
		else:
			print(f"{arg} = {col_letters(int(arg))}")


