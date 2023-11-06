class SeatManager {
public:
    set<int> s;
    SeatManager(int n) {
        for(int i=1; i <= n; i++)
            s.insert(i);
    }
    
    int reserve() {
        int r = *s.begin();
        s.erase(r);
        return r;
    }
    
    void unreserve(int sn) {
        s.insert(sn);
        return;
    }
};

/**
 * Your SeatManager object will be instantiated and called as such:
 * SeatManager* obj = new SeatManager(n);
 * int param_1 = obj->reserve();
 * obj->unreserve(seatNumber);
 */