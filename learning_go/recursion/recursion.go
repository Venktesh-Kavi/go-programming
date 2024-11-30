package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func main() {
	fmt.Println("Enter space separated integers")
	r := bufio.NewReader(os.Stdin)
	ip, _ := r.ReadString('\n')
	ip = strings.TrimSpace(ip) // in go carriage returns (\r), tabs (\t), new lines (\n) and spaces are treated as spaces in go.
	strArr := strings.Split(ip, " ")
	res := make([]int, len(strArr))
	for i, str := range strArr {
		v, e := strconv.Atoi(str)
		if e != nil {
			log.Fatalf("error converting %s to int", str)
		}
		res[i] = v
	}
	fmt.Println(res)
	root := constructTree()
	o := inorder(root)
	fmt.Printf("inorder: %v\n", o)
}

func inorder(root *TreeNode) []int {
	res := []int{}
	/**
	 * The following code is faulty. The res slice is not being updated. even though res we initialise res to have a fix capabity, so that the backing array does not change.
	 * Slices references are also passed by value, which means that the backing array is shared between the caller and the callee. So, the backing array is shared between the caller and the callee.
	 * But the appended values are not visible to res in caller method. The reason is because the descriptor is itself is not updated. So len and capacity parameters denote the old value which is 0.
	 * Either return res value to update the descriptor or pass the slice as a pointer.
	 */
	//traverseFaulty(root, res)
	return traverseCorrect(root, res)
}

func traverseCorrect(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = traverseCorrect(root.Left, res)
	res = append(res, root.Value)
	res = traverseCorrect(root.Right, res)
	return res
}

func traverseFaulty(root *TreeNode, res []int) {
	if root == nil {
		return
	}
	traverseFaulty(root.Left, res)
	res = append(res, root.Value)
	traverseFaulty(root.Right, res)
}

func constructTree() *TreeNode {
	root := &TreeNode{
		Value: 1,
	}
	ln := &TreeNode{
		Value: 2,
	}
	root.Left = ln
	rn := &TreeNode{
		Value: 3,
	}
	root.Right = rn

	return root
}
