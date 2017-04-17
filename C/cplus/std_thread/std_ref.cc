#include "iostream"
#include "thread"

class test
{
public:
	test(){};
	~test(){};
	void print()
	{
		std::cout << "hello world" << std::endl;
	}
};

void print(int& i)
{
	std::cout << "in print: "<< ++i << std::endl;
};

int main(int argc, char const *argv[])
{
	//test x;
	int i = 3;
	//std::thread t(&test::print, &x);
	std::thread t(print, std::ref(i));
	t.join();
	std::cout << "in main: " << i << std::endl;
	return 0;
}