class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& t) {
        set<int> high = {t[t.size() - 1]};
        unordered_map<int, int> index;
        index[t[t.size() - 1]] = t.size() - 1;
        vector<int> v(t.size(), 0);
        int minIndex;
        
        for(int i = t.size() - 2; i >= 0; i--) {
            high.insert(t[i]);
            index[t[i]] = i;
            auto it = high.find(t[i]);
            minIndex = INT_MAX;
            
            while(++it != high.end())
                minIndex = min(minIndex, index[*it]);
            if(minIndex != INT_MAX)
                v[i] = minIndex - i;
        }
        return v;
    }
};