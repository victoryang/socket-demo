#include <iostream>
#include <stdint.h>

using namespace std;

class bitmap
{
public:
	bitmap(int total){
		len = total/8+1;
		m_value = new uint8_t[len];
		cout << "len is: " << len << endl;
	};
	void add(int64_t value) {
		cout << "adding " << value << "..." << endl;
		int32_t offset = value & 7;
		int64_t index = value >> 3;
		uint8_t bitpos = 0x1 << offset;
		m_value[index] |= bitpos;
	};
	bool contain(int64_t value){
		cout << "checking " << "..." << endl;
		int32_t offset = value & 7;
		int64_t index = value >> 3;
		uint8_t bitpos = 0x1 << offset;
		return m_value[index] & bitpos;
	};
	void clear(int64_t value){
		cout << "removing " << value << "..." << endl;
		int32_t offset = value & 7;
		int64_t index = value >> 3;
		uint8_t bitpos = ~(0x1 << offset);
		m_value[index] &= bitpos;
	};
	~bitmap(){
		delete[] m_value;
	};
private:
	uint8_t *m_value;
	int64_t len;
};

int main(int argc, char const *argv[])
{
	bitmap b(50);
	b.add(30);
	cout << "bitmap contain 30: " << (b.contain(30)?"true":"false")<<endl;
	b.clear(30);
	cout << "bitmap contain 30: " << (b.contain(30)?"true":"false")<<endl;
	return 0;
}