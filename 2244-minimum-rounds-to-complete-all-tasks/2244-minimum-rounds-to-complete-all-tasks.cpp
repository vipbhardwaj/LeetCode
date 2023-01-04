class Solution {
public:
    int minimumRounds(vector<int>& tasks) {
        unordered_map<int, int> m;
        for(int i: tasks)
            m[i]++;
        
        int res = 0;
        for(auto i: m) {
            if(i.second == 1)
                return -1;
            else
                while(i.second > 0) {
                    if(i.second == 2 || i.second == 4)
                        i.second -= 2;
                    else
                        i.second -= 3;
                    res++;
                }
        }
        return res;
    }
};