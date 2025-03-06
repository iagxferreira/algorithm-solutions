vector<int> reverseArray(vector<int> a) {
    stack<int> buffer;
    vector<int> reversed_array;
    for(int item : a){
        buffer.push(item);
    }
    for(int item : a){
        reversed_array.push_back(buffer.top());
        buffer.pop();
    }
    return reversed_array;
}