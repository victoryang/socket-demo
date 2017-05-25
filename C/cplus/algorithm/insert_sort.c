#include <stdio.h>
#include "array.h"
//#define MAX 10

int main(int argc, char const *argv[])
{
	//int a[MAX] = {0,3,4,7,2,9,1,6,5,8};
	int i = 0, j = 0;

	for (i = 2; i < MAX; ++i)
	{
		/* code */
		a[0] = a[i];
		for (j = i-1; a[0] < a[j] && j >= 0; j--)
		{
			/* code */
			a[j+1] = a[j];
		}
		a[j+1] = a[0];
	}
	for(i = 1; i<MAX; i++)
	{
		
		printf("%d\n", a[i]);

	}
	return 0;
}