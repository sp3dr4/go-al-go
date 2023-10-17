package btbfs

import (
	"fmt"

	"github.com/sp3dr4/go-al-go/internal/queue"
)

type BinaryNode[T any] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func Bfs(head *BinaryNode[int], needle int) bool {
	q := queue.NewMyQueue[BinaryNode[int]]()
	q.Enqueue(*head)

	for q.Len() > 0 {
		curr, err := q.Deque()
		if err != nil {
			panic(fmt.Errorf("should have found something, got %v", err))
		}

		if curr.value == needle {
			return true
		}

		if curr.left != nil {
			q.Enqueue(*curr.left)
		}
		if curr.right != nil {
			q.Enqueue(*curr.right)
		}
	}

	return false
}
