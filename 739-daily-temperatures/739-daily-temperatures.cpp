class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& t) {
        // set<int> high = {t[t.size() - 1]};
        map<int, int> index;
        index[t[t.size() - 1]] = t.size() - 1;
        vector<int> v(t.size(), 0);
        int minIndex, maxTemp = t[t.size() - 1];
        
        for(int i = t.size() - 2; i >= 0; i--) {
            index[t[i]] = i;
            maxTemp = max(maxTemp, t[i]);
            if(t[i] < t[i+1]) {
                v[i] = 1;
                continue;
            } if(t[i] == maxTemp) {
                v[i] = 0;
                continue;
            }
            auto it = index.find(t[i]);
            minIndex = INT_MAX;
            
            while(++it != index.end())
                minIndex = min(minIndex, it->second);
            if(minIndex != INT_MAX)
                v[i] = minIndex - i;
        }
        return v;
    }
};