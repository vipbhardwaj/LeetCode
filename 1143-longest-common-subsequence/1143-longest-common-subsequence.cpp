class Solution {
public:
    int longestCommonSubsequence(string s1, string s2) {
        vector<vector<int>> v(s1.size() + 1, vector<int>(s2.size() + 1, 0));
        for(int i=0; i < s1.size(); i++)
            for(int j=0; j < s2.size(); j++)
                v[i+1][j+1] = (s1[i] == s2[j]) ? (1 + v[i][j]) : max(v[i+1][j], v[i][j+1]);
        return v[s1.size()][s2.size()];
    }
};