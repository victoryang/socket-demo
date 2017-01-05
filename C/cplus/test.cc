#include"test.h"
using namespace std;

void test(A &a)
{
	cout << "no need to copy in test()" << endl;
}

void test1(A a)
{
	cout << "need copy in test1()" <<endl;
}

int main (int argc, char* argv[])
{	
	A a(1);
	test(a);
	test1(a);
	return 0;
}
