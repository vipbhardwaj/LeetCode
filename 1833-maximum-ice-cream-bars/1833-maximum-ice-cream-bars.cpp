class Solution {
public:
    int maxIceCream(vector<int>& costs, int coins) {
        sort(costs.begin(), costs.end());
        
        int res = 0, i = 0;
        while(i < costs.size()) {
            if(costs[i] > coins)
                return res;
            coins -= costs[i];
            res++;
            i++;
        }
        return res;
    }
};