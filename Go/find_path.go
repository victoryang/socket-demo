package main

import "fmt"

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

    if gPath[x][y] == 2 {
        gValue[x][y] = 1
        return true
    }

    gValue[x][y] = 1

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

func find_path_loop_bfs(x int, y int) {
    var found bool = false
    for found == false {
        if check_pos_valid(x, y-1) {
            if gPath[x][y] == 2 {
                gValue[x][y] = 1
                found = true
            }
            gValue[x][y] = 1
            for {
                
            }
            gValue[x][y] = 0
        }
    }
}

func main() {
    find_path_recursive_bfs(1, 0)
    fmt.Println("res: ", gValue)
}