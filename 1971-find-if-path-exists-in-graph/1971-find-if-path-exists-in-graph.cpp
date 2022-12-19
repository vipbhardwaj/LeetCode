class Solution {
public:
    unordered_map<int, vector<int>> paths;
    
    bool validPath(int n, vector<vector<int>>& edges, int source, int destination) {
        for(vector<int> v: edges) {
            paths[v[0]].push_back(v[1]);
            paths[v[1]].push_back(v[0]);
        }
        
        vector<bool> s(n, false);
        return dfs(source, destination, s);
    }
    
    bool dfs(int e, int d, vector<bool>& v) {
        if(e == d)
            return true;
        if(!v[e]) {
            v[e] = true;
            for(int next: paths[e])
                if(dfs(next, d, v))
                    return true;
        }
        return false;
    }
};