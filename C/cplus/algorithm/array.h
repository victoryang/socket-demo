#ifndef ARRAY_H
#define ARRAY_H
#define MAX 10
int a[MAX] = {0,3,4,7,2,9,1,6,5,8};

#define PRINTF for(int i = 1; i<MAX; i++)                           \
               {													\
					cout << a[i] << endl;							\
			   }
#endif // ARRAY_H