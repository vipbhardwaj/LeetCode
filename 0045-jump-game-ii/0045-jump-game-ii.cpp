class Solution {
public:
    int jump(vector<int>& nums) {
        int jumps = 0, currMax = 0, lastJump = 0;
        
        for(int i=0; i < nums.size() - 1; i++) {
            currMax = max(currMax, i + nums[i]);
            
            if(lastJump == i) {
                lastJump = currMax;
                jumps++;
            }
        }
        return jumps;
    }
};