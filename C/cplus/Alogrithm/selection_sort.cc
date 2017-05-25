#include <iostream>
#include "array.h"

using namespace std;

int main(int argc, char const *argv[])
{
	for (int i = 1; i < MAX; ++i)
	{
		int min = 99;
		int min_pos = i;
		/*choose a minimum value from [i..MAX], swap the a[i] and a[min_pos]*/
		for (int j = i; j < MAX; ++j)
		{
			/* code */
			if (a[j] < min)
			{
				min_pos = j;
				min = a[j];
			}
		}
		int tmp = a[i];
		a[i] = a[min_pos];
		a[min_pos] = tmp;
	}
	PRINTF;
	return 0;
}