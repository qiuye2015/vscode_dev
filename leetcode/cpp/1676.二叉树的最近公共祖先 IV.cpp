/*
 * @lc app=leetcode.cn id=1676 lang=cpp
 *
 * [1676] 二叉树的最近公共祖先 IV
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, vector<TreeNode*> &nodes) {
        if (root==nullptr){ return nullptr;}
        if (count(nodes.begin(),nodes.end(),root)) return root;
        auto left=lowestCommonAncestor(root->left,nodes);
        auto right=lowestCommonAncestor(root->right,nodes);
        if (left&&right){ return root;}
        return left?left:right;
    }
};
// @lc code=end

