package comparebt

type BinaryNode[T any] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func CompareBt(a *BinaryNode[int], b *BinaryNode[int]) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.value != b.value {
		return false
	}

	return CompareBt(a.left, b.left) && CompareBt(a.right, b.right)
}
