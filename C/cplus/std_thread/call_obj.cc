#include <iostream>
#include <thread>
using namespace std;

class X
{
public:
	X(){};
	~X(){};
	void func(int a)
	{
		cout << "hello world: " << a << endl;
	}
};


int main(int argc, char const *argv[])
{
	X x;
	thread t(&X::func, &x, 10);
	t.join();
	return 0;
}