class Solution {
public:
    int longestPalindromeSubseq(string s) {
        string t = s;
        reverse(t.begin(), t.end());
        
        vector<vector<int>> v(s.size() + 1, vector<int>(t.size() + 1, 0));
        for(int i=0; i < s.size(); i++)
            for(int j=0; j < t.size(); j++)
                v[i+1][j+1] = (s[i] == t[j]) ? 1 + v[i][j] : max(v[i+1][j], v[i][j+1]);
        
        return v[s.size()][t.size()];
    }
};