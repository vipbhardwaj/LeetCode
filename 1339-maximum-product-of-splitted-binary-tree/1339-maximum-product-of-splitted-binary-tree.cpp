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
    int s;
    unsigned long long ans = 0;
    int maxProduct(TreeNode* root) {
        dfs(root);
        s = root->val;
        findMaxProd(root);
        return (int)(ans % 1000000007);
    }
    
    int dfs(TreeNode* r) {
        if(!r)
            return 0;
        r->val += (dfs(r->left) + dfs(r->right));
        return r->val;
    }

    void findMaxProd(TreeNode* r) {
        if(!r)
            return;
        ans = max(ans, ((unsigned long long)r->val * (unsigned long long)(s - r->val)));
        findMaxProd(r->left);
        findMaxProd(r->right);
    }
};