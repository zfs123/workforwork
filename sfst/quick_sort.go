func quick_sort(a []int, p int, r int) {
	if p >= r {
		return
	}
	mid := (p + r) >> 1
	i := p
	j := r
	for {
		for {
			if a[i] < a[mid] {i = i + 1} else {break}
		}

		for {
			if a[j] > a[mid] {j = j - 1} else {break}
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
		} else {
			break
		}
	}

	quick_sort(a, p, j)
	quick_sort(a, j + 1, r)
}
/*
void quick_sort(int a[], int l, int r) {
    if (l >= r) {
        return;
    }
    int i = l - 1, j = r + 1, x = a[l + r >> 1];
    while (i < j) {
        do i ++ ; while(a[i] < x);
        do j -- ; while(a[j] > x);
        if (i < j) {
            swap(a[i], a[j]);
        }
    }
    quick_sort(a, l, j);
    quick_sort(a, j + 1, r);
}
*/