class Solution {
public:
    bool canVisitAllRooms(vector<vector<int>>& rooms) {
        // vector<bool> keys(rooms.size(), false);
        vector<bool> vis(rooms.size(), false);
        dfs(0, rooms, vis);
        for(bool b: vis)
            if(!b)
                return false;
        return true;
    }
    
    void dfs(int r, vector<vector<int>>& rooms, vector<bool>& vis) {
        if(vis[r])
            return;
        vis[r] = true;
        
        for(int i: rooms[r])
            dfs(i, rooms, vis);
    }
};