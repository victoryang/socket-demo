package main

import "fmt"
import "container/list"

var gPath = [6][6]int{
        {0 , 0, 0, 0, 1, 1},
        {1,  1, 0, 0, 1, 0},
        {0 , 1, 1, 1, 1, 0},
        {0 , 0, 1, 0, 1, 2},
        {0 , 0, 1, 0, 1, 0},
        {0 , 0, 1, 1, 1, 0},
    }

var gValue [6][6]int

func check_pos_valid(x int, y int) bool {
    if x < 0 || x > 5 || y < 0 || y > 5 {
        return false
    }

    if gPath[x][y] == 0 {
        return false
    }

    if gValue[x][y] == 1 {
        return false
    }

    return true
}

func find_path_recursive_bfs(x int, y int) bool {
    if !check_pos_valid(x, y) {
        return false
    }

    gValue[x][y] = 1
    if gPath[x][y] == 2 {   
        return true
    }

    if find_path_recursive_bfs(x, y-1) {
        return true
    }

    if find_path_recursive_bfs(x-1, y) {
        return true
    }

    if find_path_recursive_bfs(x+1, y) {
        return true
    }

    if find_path_recursive_bfs(x, y+1) {
        return true
    }

    gValue[x][y] = 0
    return false
}

/*==========================================================*/

type Stack struct {
    v   *list.List
}

func NewStack() *Stack{
    v := list.New()
    return &Stack{
        v: v,
    }
}

func (s *Stack) Push(value interface{}) {
    s.v.PushBack(value)
}

func (s *Stack) Pop() interface{} {
    e := s.v.Back()
    if e != nil {
        s.v.Remove(e)
        return e.Value
    }
    
    return nil
}

func (s *Stack) Top() interface{} {
    e := s.v.Back()
    if e != nil {
        return e.Value
    }
    
    return nil
}

type Point struct {
    X,Y int
}

func getXY(v interface{}) (int, int) {
    if v == nil {
        return -1,-1
    }

    if p,ok := v.(Point); ok {
        return p.X, p.Y
    }

    return -1, -1
}

func find_path_loop_non_recursive(p Point) {
    s := NewStack()
    s.Push(p)

    for s.Top() != nil {
        x,y := getXY(s.Top())
        gValue[x][y] = 1
        if gPath[x][y] == 2 {
            _ = s.Pop()
            break
        }

        move := false
        if check_pos_valid(x, y-1) {
            s.Push(Point{x, y-1})
            move = true
        }

        if check_pos_valid(x-1, y) {
            s.Push(Point{x-1, y})
            move = true
        }

        if check_pos_valid(x+1, y) {
            s.Push(Point{x+1, y})
            move = true
        }

        if check_pos_valid(x, y+1) {
            s.Push(Point{x, y+1})
            move = true
        }

        if move == false {
            s.Pop()
            gValue[x][y] = 0
        }
    }
}

/*===================================================*/

func main() {
    //find_path_recursive_bfs(1, 0)
    find_path_loop_non_recursive(Point{1, 0})
    fmt.Println("res: ")
    for _,v :=range gValue {
        fmt.Println(v)
    }
}