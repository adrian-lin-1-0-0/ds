package ds

type SegmentNode struct {
	sum   int
	start int
	end   int
	mid   int
	left  *SegmentNode
	right *SegmentNode
}

func (node *SegmentNode) Update(index, val int) {
	if node.start == index && node.end == index {
		node.sum = val
		return
	}
	if index <= node.mid {
		node.left.Update(index, val)
	} else {
		node.right.Update(index, val)
	}
	node.sum = node.left.sum + node.right.sum
}

func (node *SegmentNode) SumRange(start, end int) int {
	if node.start == start && node.end == end {
		return node.sum
	}
	if end <= node.mid {
		return node.left.SumRange(start, end)
	} else if start > node.mid {
		return node.right.SumRange(start, end)
	}
	return node.left.SumRange(start, node.mid) + node.right.SumRange(node.mid+1, end)
}

func NewSegmentTree(vals []int) *SegmentNode {
	var buildTree func(start, end int) *SegmentNode
	buildTree = func(start, end int) *SegmentNode {
		if start == end {
			return &SegmentNode{
				start: start,
				end:   end,
				sum:   vals[start],
			}
		}

		mid := (start + end) / 2
		left := buildTree(start, mid)
		right := buildTree(mid+1, end)
		return &SegmentNode{
			start: start,
			end:   end,
			mid:   mid,
			sum:   left.sum + right.sum,
			left:  left,
			right: right,
		}
	}

	return buildTree(0, len(vals)-1)
}
