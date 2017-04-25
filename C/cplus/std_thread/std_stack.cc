#include <iostream>
#include <stack>
using namespace std;

int main(int argc, char const *argv[])
{
	stack<int> s;
	s.push(5);
	s.push(3);
	cout << s.top() << endl;
	s.pop();
	cout << s.top() << endl;
	return 0;
}