class Solution {
public:
    int countAsterisks(string s) {
        bool flag = true;
        int res = 0;
        
        for(char c: s) {
            if(c == '|')
                flag = !flag;
            if(c == '*' && flag)
                res++;
        }
        return res;
    }
};