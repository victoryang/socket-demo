#include <iostream>

#define DISALLOW_COPY_AND_ASSIGN(TypeName)		\
		TypeName(const TypeName &);				\
		void operator=(const TypeName &)

namespace myClass 
{
	class A
	{
		public:
			A();
		private:
			DISALLOW_COPY_AND_ASSIGN(A);
	};
}  //namespace

int main ()
{
	
	return 0;
}
