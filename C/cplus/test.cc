#include"test.h"
namespace myClass {
void test(A &a)
{
	std::cout << "no need to copy in test()" << std::endl;
}

void test1(A a)
{
	std::cout << "need copy in test1()" <<std::endl;
}
}  //namespace
int main (int argc, char* argv[])
{	
	::myClass::A a(1);
	::myClass::test(a);
	namespace sub = ::myClass::sub0; 
	::sub::print();
        //test1(a);
	return 0;
}
