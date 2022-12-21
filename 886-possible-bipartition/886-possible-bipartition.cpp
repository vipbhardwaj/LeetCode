class Solution {
public:
    unordered_map<int, vector<int>> m;
    
    bool possibleBipartition(int n, vector<vector<int>>& dislikes) {
        for(vector<int> v: dislikes) {
            m[v[0]].push_back(v[1]);
            m[v[1]].push_back(v[0]);
        }
        vector<int> v(n+1, -1);
        
        for(int i=1; i <= n; i++)
            if(v[i] == -1)
                if(!dfs(i, 0, v))
                    return false;
        return true;
    }
    
    bool dfs(int node, int color, vector<int>& v) {
        v[node] = color;
        
        for(int i: m[node]) {
            if(v[node] == v[i])
                return false;
            if(v[i] == -1)
                if(!dfs(i, 1 - color, v))
                    return false;
        }
        return true;
    }
};