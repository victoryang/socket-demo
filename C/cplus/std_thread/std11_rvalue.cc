#include <iostream>
#include <string>

using namespace std;

void func (string && a, string && b)
{
	string c = b;
	b = a;
	a = c;
};

int main(int argc, char const *argv[])
{
	string a("a");
	string b("b");
	cout << "a: " << a << " b: " << b << endl;
	func(move(a),move(b));
	cout << "a: " << a << " b: " << b << endl;
	return 0;
}