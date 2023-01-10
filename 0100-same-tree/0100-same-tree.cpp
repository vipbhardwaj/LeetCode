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
    void inorder(TreeNode* p, vector<TreeNode*>& hp) {
        if(p) {
            hp.push_back(p);
            inorder(p->left, hp);
            inorder(p->right, hp);
        } else hp.push_back(NULL);
    }
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        vector<TreeNode*> hq;
        vector<TreeNode*> hp;
        inorder(p, hp);
        inorder(q, hq);
        
        int l = hp.size();
        if(hq.size() != l) return false;
        
        for(int i=0; i<l; i++) {
            if(hq[i] == NULL && hp[i] != NULL) return false;
            if(hp[i] == NULL && hq[i] != NULL) return false;
            if(hq[i] != NULL && hp[i] != NULL && hq[i]->val != hp[i]->val) return false;
        }
        return true;
    }
};