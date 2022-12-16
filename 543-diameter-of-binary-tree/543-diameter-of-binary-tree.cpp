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
    int diameterOfBinaryTree(TreeNode* root) {
        int d = 0;
        dfs(root, d);
        return d;
    }

    int dfs(TreeNode* r, int& d) {
        if(!r)
            return 0;
        int ld = dfs(r->left, d);
        int rd = dfs(r->right, d);
        d = max(d, rd + ld);
        return max(rd, ld) + 1;
    }
};