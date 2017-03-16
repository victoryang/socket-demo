#include <iostream>

using namespace std;

namespace {
	class tset
	{
	public:
		tset(){cout << "in constructor" << endl;};
		~tset(){};		
	};

	void init()
	{
		cout << "in init" << endl;
		tset a;
	}
}

namespace myspace {
	void func(){
		cout << "in func" << endl;
		init();
	}
}

int main(int argc, char const *argv[])
{
	::tset a;
	init();
	myspace::func();
	return 0;
}