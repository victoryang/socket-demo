#include <iostream>
#include "stdlib.h"
#include "array.h"
using namespace std;

void merge(int *array, int begin, int mid, int end)
{
	int *tmp = new int[end-begin+1]();
	int i = begin;
	int j = mid+1;
	int k = 0;
	while(i<=mid&&j<=end)
		tmp[k++] = (array[i]<array[j])?array[i++]:array[j++];
	while(i<=mid)
		tmp[k++] = array[i++];
	while(j<=end)
		tmp[k++] = array[j++];
	for (int i = 0; i < end-begin+1; ++i)
	{
		/* code */
		array[begin+i] = tmp[i];
	}
	delete[] tmp;
	tmp = 0;
}

void mergeSort(int *array, int begin, int end)
{
	if (begin < end) {
		int mid = (begin + end)/2;
		mergeSort(array, begin, mid);
		mergeSort(array, mid+1, end);
		merge(array, begin, mid, end);
	}
}

int main(int argc, char const *argv[])
{
	mergeSort(a, 1, MAX-1);
	PRINTF;
	return 0;
}