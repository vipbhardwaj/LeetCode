class Solution {
public:
    int minFallingPathSum(vector<vector<int>>& grid) {
        vector<vector<int>> dp = grid;
        for(int j=1; j < dp.size(); j++) {
            for(int k=0; k < dp.size(); k++) {
                int e = INT_MAX;
                for(int x=0; x < dp.size(); x++) {
                    if(x == k)
                        continue;
                    e = min(e, dp[j-1][x]);
                }
                dp[j][k] += e;
            }
        }
        return *min_element(begin(dp[dp.size() - 1]), end(dp[dp.size() - 1]));
    }
};