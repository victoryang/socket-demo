#include "iostream"
#include "thread"

class function
{
public:
	function(){};	
	~function(){};
	void operator()();
};

void function::operator()()
{
		std::cout << "in operator" << std::endl;
};

int	main(int argc, char const *argv[])
{
	function f;
	std::thread t(f);
	std::cout << "t.joinable: " << t.joinable() << std::endl;
	t.join();
	return 0;
}