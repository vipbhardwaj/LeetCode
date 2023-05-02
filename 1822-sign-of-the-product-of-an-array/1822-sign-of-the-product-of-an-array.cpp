class Solution {
public:
    int arraySign(vector<int>& nums) {
        int odd = 0;
        for(int i: nums) {
            if(!i)
                return 0;
            if(i < 0)
                odd++;
        }
        if(odd % 2)
            return -1;
        return 1;
    }
};