#include <iostream>
#include "array.h"

using namespace std;

void set_sorted_value(int *a, int len)
{
	for (int i = len; i > 1; i = i >> 1/*recursively compared*/)
	{
		/* code */
		if (a[i>>1] >= a[i])    //always compare the current node with parent node
			break;
		int tmp = a[i];
		a[i] = a[i>>1];
		a[i>>1] = tmp;
	}
};

void buildMaximumHeap(int *array, int length)
{
	int index = 1;
	for(;index < length; index++)
		set_sorted_value(array, index);
};

int adjust_normal_postion(int *a, int index)
{
	int left = index << 1;
	int right = left + 1;
	int swap;

	if(a[index] >= a[left]) {
		if(a[index] >= a[right]) {
			return -1;
		} else {
			swap = right;
		}
	} else {
		if (a[index] >= a[right]) {
			swap = left;
		} else {
			swap = (a[left] > a[right]) ? left : right;
		}
	}

	int tmp = a[swap];
	a[swap] = a[index];
	a[index] = tmp;

	return swap;
};

bool adjust_leaf_position(int *a, int index)
{
	if (a[index] > a[index << 1])
		return true;
	int tmp = a[index];
	a[index] = a[index << 1];
	a[index << 1] = tmp;
	return false;
};

void rebuildHeap(int *array, int index, int len)
{
	int swap = 0;
	if (len < index << 1)  return;
	if (len == index << 1) {
		adjust_leaf_position(array, index);
		return;
	}

	if (-1 != (swap = adjust_normal_postion(array, index))) {
		rebuildHeap(array, swap, len);
	}
};

void heapSort(int *array, int length)
{
	if (length == 0 || NULL == array) return;
	buildMaximumHeap(array, length);

	for (int i = length-1; i > 1; i--)
	{
		int tmp = a[i];
		a[i] = a[1];
		a[1] = tmp;

		rebuildHeap(array, 1, i-1);
	}
}

int main(int argc, char const *argv[])
{
	/* code */
	//int a[9] = {0,3,5,4,6,2,8,7,1};
	heapSort(a,MAX);
/*	for (int i = 1; i < 9; ++i)
	{
		cout << a[i] << endl;
	}*/
	PRINTF;
	return 0;
}