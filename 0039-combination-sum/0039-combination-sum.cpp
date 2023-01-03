class Solution {
public:
    vector<vector<int>> ans;
    
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        vector<int> temp;
        solve(0, temp, target, candidates);
        return ans;
    }
    
    void solve(int i, vector<int>& temp, int sum, vector<int>& c) {
        if(!sum) {
            ans.push_back(temp);
            return;
        }
        if(sum < 0 || i == c.size())
            return;
        
        solve(i+1, temp, sum, c);
        temp.push_back(c[i]);
        solve(i, temp, sum - c[i], c);
        temp.pop_back();
    }
};