class Solution {
public:
    int maximumGain(string s, int x, int y) {
        if (x < y) {
            // Swap points
            int temp = x;
            x = y;
            y = temp;
            // Reverse the string to maintain logic
            reverse(s.begin(), s.end());
        }

        int pts = 0, aCount= 0, bCount = 0;
        // return recur(s, 1, ch, hc, x, y, 0);
        for(char c: s) {
            // if (c == hc and st.top() == ch) {
            //     while (st.)
            //     cout << st.top() << " - " << c << endl;
            //     st.pop();
            //     pts += x;
            // }
            // if (c == ch and st.top() == hc) {
            //     cout << st.top() << " - " << c << endl;
            //     st.pop();
            //     pts += y;
            // }
            // st.push(c);
            if (c == 'a') {
                aCount++;
            } else if ( c== 'b') {
                if (aCount) {
                    aCount--;
                    pts += x;
                } else {
                    bCount++;
                }
            } else {
                pts += min(aCount, bCount) * y;
                aCount = bCount = 0;
            }
        }
        return pts + (min(aCount, bCount) * y);
    }
    
    int recur(string s, int i, char ch, char hc, int x, int y, int pts) {
        if (i == s.size()) {
            return pts;
        }
        
        cout << s[i] << ", " << pts << endl;
        if (s[i] == ch and s[i-1] == hc) {
            pts += recur(s, i+1, ch, hc, x, y, pts+min(x, y));
        } else if (s[i] == hc and s[i-1] == ch) {
            pts += recur(s, i+1, ch, hc, x, y, pts+max(x, y));
        } else {
            pts += recur(s, i+1, ch, hc, x, y, pts);
        }
        
//         recur(s, i+1, ch, hc, x, y, pts);
        
//         if (s[i] == ch and s[i-1] == hc) {
//             pts += min(x, y);
//         } else if (s[i] == hc and s[i-1] == ch) {
//             pts += max(x, y);
//         }
//         cout << s[i] << ", " << pts << endl;

        return pts;
        
//         cout << s[i] << ", " << pts << endl;
//         // if (s[i] == ch and s[i-1] == hc) {
//         //     pts+=min(x, y);
//         // }
//         if (s[i] == hc and s[i-1] == ch) {
//             pts+=max(x, y);
//         }
        
//         pts += recur(s, i+1, ch, hc, x, y, pts);
        
//         if (s[i] == ch and s[i-1] == hc) {
//             pts += min(x, y);
//         }
//         // else if (s[i] == hc and s[i-1] == ch) {
//         //     pts += max(x, y);
//         // }
//         cout << s[i] << ", " << pts << endl;
    }
};