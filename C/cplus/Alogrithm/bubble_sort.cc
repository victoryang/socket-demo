#include <iostream>
#include "array.h"
using namespace std;

int main(int argc, char const *argv[])
{
	/* code */
	for (int i = 1; i < MAX; ++i)
	{
		/* code */
		bool swap = false;
		for (int j = MAX-1; j > 1; j--)
		{
			/* code */
			if (a[j] < a[j-1])
			{
				/* code */
				int tmp = a[j-1];
				a[j-1] = a[j];
				a[j] = tmp;
				swap = true;
			}
		}
		/*if swap is false, means there's no heavy bubbles in the [j..n] sections any more*/
		if (!swap)
			break;
	}
	
	PRINTF
	return 0;
}