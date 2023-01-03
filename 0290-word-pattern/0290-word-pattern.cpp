class Solution {
public:
    bool wordPattern(string pattern, string s) {
        unordered_map<char, string> m1;
        unordered_map<string, char> m2;
        string token;
        int i=0;
        
        for(char c:pattern) {
            token = "";
            while(i < s.size() && s[i] != ' ')
                token += s[i++];
            i++;
            
            if(!m1.count(c)) {
                m1[c] = token;
                if(m2.count(token)) {
                    if(c != m2[token])
                        return false;
                } else
                    m2[token] = c;
            } else
                if(token != m1[c])
                    return false;
        }
        if(i-1 != s.size())
            return false;
        return true;
    }
};