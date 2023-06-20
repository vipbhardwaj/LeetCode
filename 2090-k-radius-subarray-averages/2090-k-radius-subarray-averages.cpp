class Solution {
public:
    vector<int> getAverages(vector<int>& nums, int k) {
        int n = nums.size();
        if(n < 2*k + 1)
            return vector<int>(n, -1);
        
        long winSum = 0;
        vector<int> v(k, -1);
        for(int i=0; i < (2*k + 1); i++)
            winSum += nums[i];
        
        for(int i=k; i < n-k; i++) {
            v.push_back(winSum / (2*k + 1));
            winSum -= nums[i - k];
            if(n > i+k+1)
                winSum += nums[i + k + 1];
            else
                break;
        }
        while(k--)
            v.push_back(-1);
        return v;
    }
};