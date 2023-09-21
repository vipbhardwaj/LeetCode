class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int m = nums1.size();
        int n = nums2.size();
        double median;
        vector<int> vect;
        
        if(m == 0) vect = nums2;
        else if(n == 0) vect = nums1;
        else {
            int i=0, j=0;
            
            while(i < m && j < n) {
                if(nums1[i] < nums2[j]) vect.push_back(nums1[i++]);
                else if(nums1[i] == nums2[j]) {
                    vect.push_back(nums1[i++]);
                    vect.push_back(nums2[j++]);
                } else vect.push_back(nums2[j++]);
            }
            while(i < m) vect.push_back(nums1[i++]);
            while(j < n) vect.push_back(nums2[j++]);
        }
        int size = vect.size();
        
        if(size % 2) return (double)vect[size / 2];
        else return ((double)vect[size / 2] + (double)vect[(size / 2) - 1]) / 2;
    }
};