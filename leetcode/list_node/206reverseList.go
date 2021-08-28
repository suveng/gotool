package main

func main() {
	node := readListNode()
	reverseList(node)
}

func readListNode() *ListNode {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	return head
}

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//
// Related Topics 递归 list_node
// 👍 1921 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
思路:
1. 遍历源链表,
1.1 每次遍历都将原链表的头节点插入新链表的尾部
1.2 将原链表的头节点指向下一个节点

*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode

	for head != nil {
		tempHead := head.Next
		head.Next = newHead
		newHead = head
		head = tempHead
	}

	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region end(Prohibit modification and deletion)
