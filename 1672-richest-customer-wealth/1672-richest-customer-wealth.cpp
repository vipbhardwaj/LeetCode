class Solution {
public:
    int maximumWealth(vector<vector<int>>& a) {
        int sum, maxSum = 0;
        
        for(vector<int> v: a) {
            sum = 0;
            for(int i: v)
                sum += i;
            maxSum = max(maxSum, sum);
        }
        return maxSum;
    }
};