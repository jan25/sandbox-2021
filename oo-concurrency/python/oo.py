
class Shape:
    def area(self):
        raise Exception('Not implemented')

    def volume(self):
        raise Exception('Not implemented')

class Rectrangle(Shape):
    def __init__(self, width, height):
        self.width, self.height = width, height
    
    def area(self):
        return self.width * self.height

class Square(Rectrangle):
    def __init__(self, side):
        super().__init__(side, side)

class Cube(Shape):
    def __init__(self, width):
        self.width = width

    def area(self):
        return self.width * self.width * 6

    def volume(self):
        return self.width ** 3

r = Rectrangle(2, 3)
print(r.area())
# print(r.volume())

sq = Square(3)
print(sq.area())

c = Cube(5)
print(c.area(), c.volume())
