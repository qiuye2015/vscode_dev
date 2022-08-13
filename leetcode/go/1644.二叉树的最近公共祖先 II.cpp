/*
 * @lc app=leetcode.cn id=1644 lang=cpp
 *
 * [1644] 二叉树的最近公共祖先 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode *ans;
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        dfs(root,p,q);
        return ans;
    }
    //以root为头结点的树中是否含有p或者q
    bool dfs(TreeNode* root, TreeNode* p, TreeNode* q){
        if(root==NULL){ return false; }
        auto l = dfs(root->left,p,q);
        auto r = dfs(root->right,p,q);
        //左树true，右树也是true，说明左树发现了p，右树发现了q
        //或者左树发现了q，右树发现了p
        if (l&&r) ans = root;
        //root就是p，在左树或右树发现了q，
        //或者root就是q，在左树或右树发现了p
        if ((root->val==p->val|| root->val==q->val)&&(l||r)){ ans=root; }
        //左树或者右树有p或q之一，或者root就是p或q就返回true
        return l || r || root->val==p->val || root->val==q->val;
    }
};
// @lc code=end

