class Solution {
public:
    int largestAltitude(vector<int>& gain) {
        int res = 0, sum = 0;
        
        for(int i: gain) {
            sum += i;
            res = max(res, sum);
        }
        return res;
    }
};