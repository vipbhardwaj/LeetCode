class SmallestInfiniteSet {
public:
    unordered_set<int> s;
    int smallest = 1;
    
    SmallestInfiniteSet() {}
    
    int popSmallest() {
        s.insert(smallest);
        int res = smallest;
        while(s.find(smallest) != s.end())
            smallest++;
        return res;
    }
    
    void addBack(int num) {
        if(s.find(num) == s.end())
            return;
        s.erase(num);
        smallest = min(smallest, num);
    }
};
/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * SmallestInfiniteSet* obj = new SmallestInfiniteSet();
 * int param_1 = obj->popSmallest();
 * obj->addBack(num);
 */