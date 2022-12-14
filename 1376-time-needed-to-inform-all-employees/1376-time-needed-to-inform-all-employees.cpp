class Solution {
public:
    int res = 0;
    
    int numOfMinutes(int n, int headID, vector<int>& manager, vector<int>& informTime) {
        unordered_map<int, vector<int>> m;
        for(int i=0; i < manager.size(); i++)
            m[manager[i]].push_back(i);
        
        countTime(headID, m, informTime, 0);
        return res;
    }
    
    void countTime(int h, unordered_map<int, vector<int>>& m, vector<int>& t, int time) {
        // cout << h << " - " << time << endl;
        if(!m.count(h))
            res = max(res, time);
        for(int i: m[h])
            countTime(i, m, t, time + t[h]);
    }
};