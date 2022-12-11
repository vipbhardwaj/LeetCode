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
    int res = INT_MIN;
    int maxPathSum(TreeNode* root) {
        dfs(root);
        return res;
    }
    
    int dfs(TreeNode* t) {
        if(!t)
            return 0;
        int l = max(0, dfs(t->left));
        int r = max(0, dfs(t->right));
        res = max(res, t->val + l + r);
        return max(t->val + l, t->val + r);
    }
};