package btpostorder

type BinaryNode[T any] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func walk(curr *BinaryNode[int], path *[]int) *[]int {
	if curr == nil {
		return path
	}

	walk(curr.left, path)
	walk(curr.right, path)
	*path = append(*path, curr.value)
	return path
}

func PostOrderSearch(head *BinaryNode[int]) *[]int {
	return walk(head, &[]int{})
}
