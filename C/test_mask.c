#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[])
{
	int i = 31;

	printf("i & 1 = %d\n", i&1);
	printf("i & ~1 = %d\n", i&~1);
	return 0;
}
