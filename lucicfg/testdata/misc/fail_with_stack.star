def func():
  return stacktrace()

def another_func(t):
  fail("boom", trace=t)

t = func()

another_func(t)

# Expect errors:
#
# Traceback (most recent call last):
#   //testdata/misc/fail_with_stack.star: in <toplevel>
#   //testdata/misc/fail_with_stack.star: in func
#   <builtin>: in stacktrace
# Error: boom
