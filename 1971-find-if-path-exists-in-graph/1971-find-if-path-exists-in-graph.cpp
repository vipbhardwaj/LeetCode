class Solution {
public:
    unordered_map<int, vector<int>> paths;
    
    bool validPath(int n, vector<vector<int>>& edges, int source, int destination) {
        for(vector<int> v: edges) {
            paths[v[0]].push_back(v[1]);
            paths[v[1]].push_back(v[0]);
        }
        
        unordered_set<int> s;
        dfs(source, s);
        if(s.find(source) != s.end() && s.find(destination) != s.end())
            return true;
        return false;
    }
    
    void dfs(int e, unordered_set<int>& v) {
        if(v.find(e) != v.end())
            return;
        v.insert(e);
        for(int next: paths[e])
            dfs(next, v);
    }
};