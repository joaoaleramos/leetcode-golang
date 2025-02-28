package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists and returns the merged sorted list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to act as the starting point of the merged list.
	root := &ListNode{}
	// cur is a pointer to the current node in the merged list.
	cur := root

	// Iterate through both lists until one of them is exhausted.
	for list1 != nil && list2 != nil {
		// Compare the values of the current nodes of both lists.
		if list1.Val < list2.Val {
			// If list1's value is smaller, append it to the merged list.
			cur.Next = list1
			// Move to the next node in list1.
			list1 = list1.Next
		} else {
			// If list2's value is smaller or equal, append it to the merged list.
			cur.Next = list2
			// Move to the next node in list2.
			list2 = list2.Next
		}
		// Move to the next node in the merged list.
		cur = cur.Next
	}

	// If list1 still has remaining nodes, append them to the merged list.
	if list1 != nil {
		cur.Next = list1
	} else {
		// If list2 still has remaining nodes, append them to the merged list.
		cur.Next = list2
	}

	// Return the merged list, which starts at root.Next (skipping the dummy node).
	return root.Next
}
