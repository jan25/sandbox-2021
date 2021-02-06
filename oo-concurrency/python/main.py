
def f(*args, **kwargs):
    print(args, kwargs)
    print(kwargs['name'])

if __name__ == '__main__':
    print('Hi!')
    print(f(1, 2, 'asdf', name='neil', age=43))