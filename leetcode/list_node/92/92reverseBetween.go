package main

func main() {
	node := &ListNode{Val: 3, Next: &ListNode{Val: 5,
		Next: nil}}
	reverseBetween(node, 1, 2)
}

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
// Related Topics 链表
// 👍 997 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left <= 0 || right <= 0 || right <= left {
		return head
	}

	var newHead *ListNode

	// 逆转的长度
	//len := right - left + 1
	var preLeft int
	if left-1 < 1 {
		preLeft = left
	} else {
		preLeft = left - 1
	}
	postRight := right + 1

	var preLeftNode *ListNode
	var postRightNode *ListNode
	var leftNode *ListNode
	var rightNode *ListNode

	// 从1开始
	newHead = head
	for i := 1; newHead != nil; i++ {
		// 找到逆转的头节点
		if preLeft == i {
			preLeftNode = newHead
		}
		if left == i {
			leftNode = newHead
		}
		// 找到逆转的尾节点
		// 如果没有下一位, 则需要用最后一位填上
		if postRight == i {
			postRightNode = newHead
		}
		if right == i {
			rightNode = newHead
		}
		newHead = newHead.Next

		if newHead == nil && i+1 == postRight {
			postRight = i
			postRightNode = rightNode
		}
	}
	newHead = nil

	// 开始翻转
	var point *ListNode = leftNode
	rightNode = leftNode
	for point != nil {
		if postRight != right {
			if point == postRightNode {
				break
			}
		}
		tempNode := point.Next
		point.Next = newHead
		newHead = point
		point = tempNode
	}
	leftNode = newHead

	// 翻转后 源链表位置调整
	if preLeftNode != nil && preLeft != left {
		preLeftNode.Next = leftNode
	}

	if rightNode != nil && right != postRight {
		rightNode.Next = postRightNode
	}

	if left != preLeft {
		return head
	} else {
		return newHead
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region end(Prohibit modification and deletion)
