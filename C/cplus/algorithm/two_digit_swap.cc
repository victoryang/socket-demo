#include <iostream>
using namespace std;

int main(int argc, char const *argv[])
{
	int a = 1, b = 2;
	cout << "before swap... a: " << a << " b: " << b << endl;
	a = a ^ b;
	b = b ^ a;
	a = a ^ b;
	cout << "after swap... a: " << a << " b: " << b << endl;
	return 0;
}