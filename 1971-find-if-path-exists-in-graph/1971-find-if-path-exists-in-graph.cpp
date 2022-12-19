class Solution {
public:
    unordered_map<int, vector<int>> paths;
    
    bool validPath(int n, vector<vector<int>>& edges, int source, int destination) {
        for(vector<int> v: edges) {
            paths[v[0]].push_back(v[1]);
            paths[v[1]].push_back(v[0]);
        }
        
        unordered_set<int> s;
        return dfs(source, destination, s);
    }
    
    bool dfs(int e, int d, unordered_set<int>& v) {
        if(e == d)
            return true;
        if(v.find(e) == v.end()) {
            v.insert(e);
            for(int next: paths[e])
                if(dfs(next, d, v))
                    return true;
        }
        return false;
    }
};