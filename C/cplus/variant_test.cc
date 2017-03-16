#include<variant>
#include<iostream>

int main()
{
	std::variant<int, float> w, v;
	w = 12;
	std::out << "w is " << w <<std::endl;
	return 0;	
}
