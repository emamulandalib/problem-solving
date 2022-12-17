namespace LeetCodeDotnet.Solutions;

public class SlidingWindowMaximum {
    private List<int> _data = new List<int>();

    public int[] MaxSlidingWindow(int[] nums, int k) {
        var n = nums.Length;

        if (n * k == 0) {
            return new int[] {};
        }

        if (k == 1) {
            return nums;
        }

        _data = nums.ToList();
        BuildMax();
        for (var i = 0; i < n; i++) {
            
        }
        return new int[]{};
    }
    
    private void BuildMax() {
        var n = _data.Count;
        for (var i = (n / 2) - 1; i > 0; i--) {
            Heapify(i, n);
        }
    }

    private void Push(int number) {
        _data.Add(number);
        var i = _data.Count - 1;
        while (i > 0 && _data[i] > _data[Parent(i)]) {
            (_data[i], _data[Parent(i)]) = (_data[Parent(i)], _data[i]);
            i = Parent(i);
        }
    }

    private int Parent(int i) {
        return (i - 1) / 2;
    }
    
    private int Pop() {
        var m = _data[0];
        _data[0] = _data.Last();
       _data.RemoveAt(_data.Count - 1); 
       Heapify(0, _data.Count);
        return m;
    }
    
    private void Heapify(int i, int n) {
        var largest = i;
        var left = 2 * i + 1;
        var right = 2 * i + 2;
        
        if (left < n && _data[left] > _data[largest]) {
            largest = left;
        }
        
        if (right < n && _data[right] > _data[largest]) {
            largest = right;
        }
        
        if (largest != i) {
            (_data[i], _data[largest]) = (_data[largest], _data[i])
            Heapify(largest, n)
        }
    }
}