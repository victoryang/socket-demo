#include "iostream"
#include "memory"

using namespace std;

int main(int argc, char const *argv[])
{
	auto_ptr<int> a(new int(5));
	shared_ptr<int> b(new int(3));
	cout << "hello world" << *a << endl;
	cout << "hello world" << *b << endl;
	return 0;
}