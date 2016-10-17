package main

import (
    "fmt"
    li "github.com/goibibo/worktree"
)

type TwoArgs struct {
    X int
    Y int
}

func leaf2(i interface{}) interface{} {
    args := i.(TwoArgs)

    return args.X * args.Y
}

func merge2(results []interface{}) interface{} {
    var sum int
    for _, x := range results {
        sum += x.(int)
    }
    return sum
}

func leaf1(i interface{}) interface{} {
    args := i.(TwoArgs)

    return args.X + args.Y
}

func main() {

    // TWO Level work tree

    l2 := li.CommandTree{}
    l2.AddMapper(leaf2, TwoArgs{2, 3})
    l2.AddMapper(leaf2, TwoArgs{2, 2})
    l2.AddReducer(merge2)

    l1 := li.CommandTree{}
    l1.AddMapper(l2.Run, nil) // When nesting use nil for Run
    l1.AddMapper(leaf2, TwoArgs{2, 2})
    l1.AddReducer(merge2)

    fmt.Println(l1.Run(nil).(int))

}
