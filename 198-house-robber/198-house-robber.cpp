class Solution {
public:
    int rob(vector<int>& nums) {
        int n=nums.size();
        vector<int> v;
        if(n == 0) return 0;
        else v.push_back(nums[0]);
        
        if(n == 1) return nums[0];
        else v.push_back(nums[1]);
        
        if(n == 2) return max(v[n-1], v[n-2]);
        else v.push_back(nums[2] + v[0]);
        
        for(int i=3; i<n; i++) {
            v.push_back(nums[i] + max(v[i-2], v[i-3]));
        }
        return max(v[n-1], v[n-2]);
    }
};