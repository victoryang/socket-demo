#include <iostream>
#include "array.h"
using namespace std;

int partition(int *array, int begin, int end)
{
	int pos;
	int target = array[begin];
	int i,j;
	for (i = begin, j = end; i < j;)
	{
		/* code */
		if (array[j] < target)
		{
			/* code */
			array[i] = array[j];
			i++;
			for(;array[i] < target && i<j; i++);
			if(i<j)
			{
				array[j] = array[i];
				j--;
			}
		} else {
			j--;
		}
	}
	pos = i;
	array[pos] = target;
	return pos;
}

void quick_sort(int *array, int begin, int end)
{
	int pos;
	if (begin < end)
	{
		/* code */
		pos = partition(array, begin, end);
		quick_sort(array, begin, pos-1);
		quick_sort(array, pos+1, end);
	}
}

int main(int argc, char const *argv[])
{
	quick_sort(a, 1, MAX-1);
	PRINTF;
	return 0;
}