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
    int res = 0;
    int longestUnivaluePath(TreeNode* root) {
        dfs(root);
        return res;
    }
    
    int dfs(TreeNode* t) {
        if(!t)
            return 0;
        int r = dfs(t->right), l = dfs(t->left), lst = 0, rst = 0;
        
        if(t->right && t->val == t->right->val) {
            rst += r+1;
        }
        if(t->left && t->val == t->left->val) {
            lst += l+1;
        }
        res = max(res, lst + rst);
        return max(lst, rst);
    }
};