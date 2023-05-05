class Solution {
public:
    int maxVowels(string s, int k) {
        unordered_set<char> v = {'a', 'e', 'i', 'o', 'u'};
        int maxRes = 0, res = 0, i = 0, j = 0;
        
        for(; j < k; j++)
            if(v.find(s[j]) != v.end())
                res++;
        maxRes = max(res, maxRes);
        
        while(j < s.size()) {
            if(v.find(s[i]) != v.end())
                res--;
            i++;
            if(v.find(s[j]) != v.end())
                res++;
            j++;
            maxRes = max(maxRes, res);
        }
        return maxRes;
    }
};