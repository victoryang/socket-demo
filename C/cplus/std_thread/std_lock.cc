#include <iostream>
#include <mutex>

using namespace std;

void swapInt(int& a, int& b)
{
	int c = a;
	a = b;
	b = c;
};

class test
{
public:
	test(){};
	test(int const& _data):data(_data){};
	~test(){};
	void swap(test& a)
	{
		lock(this->m, a.m);
		lock_guard<mutex> lock1(this->m,adopt_lock);
		lock_guard<mutex> lock2(a.m,adopt_lock);
		swapInt(this->data, a.data);
	};
	int data;
private:
	mutex m;
};

int main(int argc, char const *argv[])
{
	test a(1);
	test b(2);
	cout << a.data << " " << b.data << endl;
	a.swap(b);
	cout << a.data << " " << b.data << endl;
	return 0;
}