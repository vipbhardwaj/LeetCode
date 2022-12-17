class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        unordered_set<string> o = {"/", "*", "+", "-"};
        stack<int> s;
        
        for(string t: tokens) {
            if(o.find(t) == o.end()) {
                s.push(stoi(t));
                continue;
            }
            long x = s.top(); s.pop();
            long y = s.top(); s.pop();
            long z;
            switch (t[0]) {
                case '/':
                    z = y / x;
                    break;
                case '*':
                    z = y * x;
                    break;
                case '+':
                    z = y + x;
                    break;
                case '-':
                    z = y - x;
                    break;
            }
            s.push((int)z);
        }
        return s.top();
    }
};