# -*- coding: utf-8 -*-
# @Time:    2021/4/20 21:23
# @Author:  leeyoung
# @File:    weighted_edge.py
# @Content: 带权值的边


class WeightedEdge:
    def __init__(self, v, w, weight):
        self.v = v
        self.w = w
        self.weight = weight

    def __cmp__(self, other):
        return self.weight if self.weight < other.weight else other.weight

    def __lt__(self, other):
        return self.weight if self.weight < other.weight else other.weight

    def get_v(self):
        return self.v

    def get_w(self):
        return self.w

    def get_wight(self):
        return self.weight

    def to_string(self):
        return "(%d - %d: %d)" % (self.v, self.w, self.weight)


def main():
    weight_edge = WeightedEdge(1, 2, 20)
    print(weight_edge.to_string())


if __name__ == '__main__':
    main()
