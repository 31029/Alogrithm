package binarytree

// Leetcode-700: 二叉搜索树中的搜索 【https://leetcode.cn/problems/search-in-a-binary-search-tree/description/】
/* 描述：
给定二叉搜索树（BST）的根节点 root 和一个整数值 val。
你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 null 。   
*/
func SearchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }

    if val > root.Val {
        return SearchBST(root.Right, val)
    } else if val < root.Val {
        return SearchBST(root.Left, val)
    } else if val == root.Val {
        return root
    }

    return nil
}
