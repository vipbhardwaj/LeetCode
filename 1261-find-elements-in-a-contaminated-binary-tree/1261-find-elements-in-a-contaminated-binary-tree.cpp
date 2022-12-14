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
class FindElements {
public:
    unordered_set<int> s;
    
    FindElements(TreeNode* root) {
        dfs(root, 0);
    }
    
    void dfs(TreeNode* r, int d) {
        if(!r)
            return;
        s.insert(d);
        dfs(r->left, 2*d + 1);
        dfs(r->right, 2*d + 2);
    }
    
    bool find(int target) {
        if(s.find(target) != s.end())
            return true;
        return false;
    }
};

/**
 * Your FindElements object will be instantiated and called as such:
 * FindElements* obj = new FindElements(root);
 * bool param_1 = obj->find(target);
 */