class Solution {
public:
    unordered_map<int, bool> v;
    unordered_map<int, vector<int>> g;
    
    int minTime(int n, vector<vector<int>>& edges, vector<bool>& hasApple) {
        for (vector<int> edge: edges) {
            g[edge[0]].push_back(edge[1]);
            g[edge[1]].push_back(edge[0]);
        }
        return dfs(0, 0, hasApple);
    }
    
    int dfs(int node, int myCost, vector<bool>& hasApple) {
	    if (v[node])
	        return 0;
	    
        v[node] = true;
        int childrenCost = 0;
        
        for (int x: g[node])
            if(!v[x])
                childrenCost += dfs(x, 2, hasApple);
        
        if (!childrenCost && !hasApple[node]) 
            return 0;
	    
        return (childrenCost + myCost);
    }
};