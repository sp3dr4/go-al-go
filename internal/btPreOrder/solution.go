package btpreorder

type BinaryNode[T any] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func walk(curr *BinaryNode[int], path *[]int) *[]int {
	if curr == nil {
		return path
	}

	*path = append(*path, curr.value)
	walk(curr.left, path)
	walk(curr.right, path)
	return path
}

func PreOrderSearch(head *BinaryNode[int]) *[]int {
	return walk(head, &[]int{})
}
