package chapter2

import (
    "errors"
    "github.com/exced/CtCI6-Go/libs/list"
)

func removeNode(n *list.Node) error {
    if n == nil || n.Next == nil {
        return errors.New("removeNode failed")
    }
    if n.Next.Next != nil {
        n.Next.Next.Prev = n
    }
    n.Value = n.Next.Value
    n.Next = n.Next.Next
    return nil
}