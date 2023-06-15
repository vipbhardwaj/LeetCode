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
    long long kthLargestLevelSum(TreeNode* root, int k) {
        
        priority_queue<long long, vector<long long>> pq;
        queue<TreeNode*> q;
        q.push(root);
        
        while(!q.empty()) {
            long long l = q.size(), sum = 0;
            
            while(l--) {
                if(q.front()->left)
                    q.push(q.front()->left);
                if(q.front()->right)
                    q.push(q.front()->right);
                
                sum += q.front()->val;
                q.pop();
            }
            pq.push(sum);
        }
        if(pq.size() < k)
            return -1;
        while(--k) {
            cout << pq.top() << ", ";
            pq.pop();
        }
        return pq.top();
    }
};