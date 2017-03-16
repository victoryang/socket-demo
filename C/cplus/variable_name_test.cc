#include <iostream>

using namespace std;


class A
{
public:
	A(){ _a = 0;};
	~A(){};
	A operator()(int c)
	{
		A tmp;
		tmp._a = c;
		cout << "in operator" <<endl;
		return tmp;
	}
public:
	int _a;
};

int main(int argc, char const *argv[])
{
	A a,b;
	b = a(3);
	cout << b._a <<endl;
	return 0;
}