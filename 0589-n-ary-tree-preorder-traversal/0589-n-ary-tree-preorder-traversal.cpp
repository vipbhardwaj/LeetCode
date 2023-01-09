/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val) {
        val = _val;
    }

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
public:
    vector<int> v;
    
    vector<int> preorder(Node* root) {
        if(!root)
            return {};
        preOrd(root);
        return v;
    }
    
    void preOrd(Node* r) {
        v.push_back(r->val);
        for(Node* c: r->children)
            preOrd(c);
    }
};