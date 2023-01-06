class Solution {
public:
    bool canReach(vector<int>& arr, int start) {
        if(!arr[start])
            return true;
        if(arr[start] == -1)
            return false;
        
        int right = arr[start] + start, left = start - arr[start];
        arr[start] = -1;
        
        return ((right < arr.size() && canReach(arr, right)) || (left >= 0 && canReach(arr, left)));
    }
};