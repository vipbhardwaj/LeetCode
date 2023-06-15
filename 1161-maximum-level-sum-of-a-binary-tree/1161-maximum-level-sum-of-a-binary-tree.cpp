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
    int maxLevelSum(TreeNode* root) {
        int lvl = 1, sum = 0, res = lvl, maxSum = INT_MIN;
        vector<int> v;
        queue<TreeNode*> q, p;
        p.push(root);
        
        while(!p.empty()) {
            while(!p.empty()) {
                q.push(p.front());
                p.pop();
            }
            while(!q.empty()) {
                sum += q.front()->val;
                if(q.front()->left)
                    p.push(q.front()->left);
                if(q.front()->right)
                    p.push(q.front()->right);
                q.pop();
            }
            if(sum > maxSum) {
                res = lvl;
                maxSum = sum;
            }
            sum = 0;
            lvl++;
        }
        return res;
    }
};