package btinorder

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
	*path = append(*path, curr.value)
	walk(curr.right, path)
	return path
}

func InOrderSearch(head *BinaryNode[int]) *[]int {
	return walk(head, &[]int{})
}
