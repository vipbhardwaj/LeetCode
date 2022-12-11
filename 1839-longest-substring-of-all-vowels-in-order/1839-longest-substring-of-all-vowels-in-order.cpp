class Solution {
public:
    int longestBeautifulSubstring(string w) {
        bool a=false, e=false, i=false, o=false, u=false;
        int res = 0, temp = 0;
        for(char c: w) {
            switch(c) {
                // cout << temp << ", ";
                case 'a':
                    if(e || i || o || u) {
                        temp = 0;
                        e = false;
                        i = false;
                        o = false;
                        u = false;
                    }
                    a = true;
                    temp++;
                    break;
                case 'e':
                    if(!a || i || o || u) {
                        temp = 0;
                        a = false;
                        e = false;
                        i = false;
                        o = false;
                        u = false;
                    } else {
                        e = true;
                        temp++;
                    }
                    break;
                case 'i':
                    if(!a || !e || o || u) {
                        temp = 0;
                        a = false;
                        e = false;
                        i = false;
                        o = false;
                        u = false;
                    } else {
                        i = true;
                        temp++;
                    }
                    break;
                case 'o':
                    if(!a || !e || !i || u) {
                        temp = 0;
                        a = false;
                        e = false;
                        i = false;
                        o = false;
                        u = false;
                    } else {
                        o = true;
                        temp++;
                    }
                    break;
                case 'u':
                    if(!a || !e || !i || !o) {
                        temp = 0;
                        a = false;
                        e = false;
                        i = false;
                        o = false;
                        u = false;
                    } else {
                        u = true;
                        temp++;
                        res = max(res, temp);
                    }
                    break;
            }
        }
        return res;
        
//         unordered_map<char, int> m;
//         int res = 0;
        
//         for(int i = 0; i < w.size(); i++) {
//             if(w[i] == 'a' && m.size() < 2)
//                 m[w[i]]++;
//             else if(w[i] == 'e' && m.size() < 3 && m.size() > 0)
//                 m[w[i]]++;
//             else if(w[i] == 'i' && m.size() < 4 && m.size() > 1)
//                 m[w[i]]++;
//             else if(w[i] == 'o' && m.size() < 5 && m.size() > 2)
//                 m[w[i]]++;
//             else if(w[i] == 'u' && m.size() < 6 && m.size() > 3)
//                 m[w[i]]++;
//             else {
//                 if(m.size() == 5) {
//                     int t = 0;
//                     for(auto i: m)
//                         t += i.second;
//                     res = max(res, t);
//                 }
//                 m.clear();
//             }
//         }
//         if(m.size() == 5) {
//             int t = 0;
//             for(auto i: m)
//                 t += i.second;
//             res = max(res, t);
//         }
//         return res;
    }
};