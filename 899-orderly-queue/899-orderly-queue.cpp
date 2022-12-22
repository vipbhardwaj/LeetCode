class Solution {
public:
    string orderlyQueue(string s, int k) {
        if(k > 1) {
            sort(s.begin(), s.end());
            return s;
        }
        
        string t, ans = s;
        for(int i=0; i < s.size(); i++) {
            // cout << s.substr(i + 1, s.size()) << " + " << s.substr(0, i + 1) << endl;
            t = s.substr(i + 1, s.size()) + s.substr(0, i + 1);
            ans = min(t, ans);
        }
        return ans;
    }
};