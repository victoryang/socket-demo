#include "iostream"
#include "functional"
#include "unistd.h"

void func(int& n1, int& n2, const int& n3)
{
	std::cout << "in func: " << n1 << ' '<< n2 << ' ' << n3 << std::endl;
	n1++;
	n2++;
	//n3++;  compiler error
};

int main(int argc, char const *argv[])
{
	int a = 1, b = 2, c = 3;
	std::function<void()> f = std::bind(func, a, std::ref(b), std::cref(c));
	a = 10;
	b = 11;
	c = 12;
	std::cout << "in func: " << a << ' '<< b << ' ' << c << std::endl;
	f();
	std::cout << "in func: " << a << ' '<< b << ' ' << c << std::endl;
	return 0;
}