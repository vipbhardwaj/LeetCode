class Solution {
public:
    int lenLongestFibSubseq(vector<int>& arr) {
        unordered_map<int, int> m;
        for(int i=0; i < arr.size(); i++)
            m[arr[i]] = i;
        vector<vector<int>> dp(arr.size(), vector<int>(arr.size(), 2));
        // for(int i=0; i < arr.size(); i++)
        //     for(int j=i+1; j < arr.size(); j++)
        //         dp[i][j] = 2;
        int ans = 0;
        for(int i=0; i < arr.size(); i++) {
            for(int j=i+1; j < arr.size(); j++) {
                if(m.count(arr[i] + arr[j])) {
                    dp[j][m[arr[i] + arr[j]]] = dp[i][j] + 1;
                    ans = max(ans, dp[i][j]+1);
                }
            }
        }
        return ans;
    }
};