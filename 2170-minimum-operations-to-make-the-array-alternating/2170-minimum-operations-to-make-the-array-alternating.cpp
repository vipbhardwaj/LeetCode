class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        array<int, 3> even = top2Freq(nums, 0), odd = top2Freq(nums, 1);
        
        if(even[0] != odd[0])
            return nums.size() - (even[1] + odd[1]);
        return nums.size() - max(even[1] + odd[2], odd[1] + even[2]);
    }
    
    array<int, 3> top2Freq(vector<int>& nums, int start) {
        int first = 0, second = 0, count[100001] = {0};
        
        for(int i = start; i < nums.size(); i+=2) {
            if(++count[nums[i]] >= count[first]) {
                if(first != nums[i])
                    second = first;
                first = nums[i];
            } else if(count[nums[i]] > count[second])
                second = nums[i];
        }
        return {first, count[first], count[second]};
    }
};