package main

import "testing"

func TestAddFind(t *testing.T) {
    ip := ImageDb{}
    ip.Add("image 1", 0, 100, 200)
    ip.Add("image 2", 50, 50, 50)
    ip.Add("image 3", 255, 255, 255)
    ip.Add("image 4", 0, 0, 0)

    var testCases = []struct {
        r int
        g int
        b int
        expected string
    }{
        {30, 30, 10, "image 4"},
        {200, 200, 200, "image 3"},
        {0, 100, 200, "image 1"},
        {0, 100, 100, "image 2"},
    }

    for _, tc := range testCases {
        actual := ip.Find(tc.r, tc.g, tc.b)
        if actual != tc.expected {
            t.Errorf("Expected %q, got %q for RGB(%d, %d, %d)", tc.expected, actual, tc.r, tc.g, tc.b)
        }
    }

}

func TestSumSquares(t *testing.T) {
    var testCases = []struct {
        r1 int
        g1 int
        b1 int
        r2 int
        g2 int
        b2 int
        expected int
    }{
        {0, 0, 0, 2, 3, 1, 14},
        {50, 50, 50, 50, 50, 50, 0},
        {10, 5, 0, 5, 11, 0, 61},
    }

    for i, tc := range testCases {
        actual := SumSquares(tc.r1, tc.g1, tc.b1, tc.r2, tc.g2, tc.b2)
        if actual != tc.expected {
            t.Errorf("Expected %q, got %q for test case %d", tc.expected, actual, i)
        }
    }
}
