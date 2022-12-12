class Solution {
public:
    int minFlips(string s) {
        int n = s.size();
        s += s;
        string s1, s2;
        for(int i=0; i < s.size(); i++) {
            s1.push_back(i % 2 ? '0' : '1');
            s2.push_back(i % 2 ? '1' : '0');
        }
        
        int res1 = 0, res2 = 0, res = INT_MAX;
        for(int i=0; i < s.size(); i++) {
            if(s[i] != s1[i])
                res1++;
            if(s[i] != s2[i])
                res2++;
            if(i >= n) {
                if(s[i-n] != s1[i-n])
                    res1--;
                if(s[i-n] != s2[i-n])
                    res2--;
            }
            if(i >= n-1)
                res = min({res, res1, res2});
        }
        return res;
    }
};