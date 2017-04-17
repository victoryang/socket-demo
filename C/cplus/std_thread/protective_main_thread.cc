#include "iostream"
#include "thread"
#include "unistd.h"

class test
{
public:
	test(){};
	~test(){};
	void operator()()
	{
		for (int i = 0; i < 10; ++i)
		{
			/* code */
			std::cout << i << std::endl;
			sleep(1);
		}
	}
};

class thread_guard
{
public:
	std::thread& t;
	explicit thread_guard(std::thread &t_):t(t_){};
	~thread_guard()
	{
		if (t.joinable())
		{
			/* code */
			std::cout << "joinable" << std::endl;
			t.join();
			std::cout << "join!!" << std::endl;
		}
	};
	thread_guard(thread_guard const&)=delete;
	thread_guard& operator=(thread_guard const&)=delete;
};

int main(int argc, char const *argv[])
{
	test hello;
	std::thread t(hello);
	thread_guard tg(t);
	t.detach();
	sleep(5);
	std::cout << "finished" << std::endl;
	return 0;
}