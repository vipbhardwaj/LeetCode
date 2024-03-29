class Solution {
public:
    vector<int> getAverages(vector<int>& nums, int k) {
        ios::sync_with_stdio(false);
        cin.tie(nullptr);
        cout.tie(nullptr);
        
        int n = nums.size();
        if(n < 2*k + 1)
            return vector<int>(n, -1);
        
        long winSum = 0;
        vector<int> v(k, -1);
        for(int i=0; i < 2 * k; i++)
            winSum += nums[i];
        
        for(int i=k; i < n-k; i++) {
            winSum += nums[i + k];
            v.push_back(winSum / (2*k + 1));
            winSum -= nums[i - k];            
        }
        while(k--)
            v.push_back(-1);
        return v;
    }
};