#ifndef TEST_H
#define TEST_H

#include<iostream>

#define DISALLOW_COPY_AND_ASSIGN(TypeName)      \
        TypeName(const TypeName&);                      \
        void operator=(const TypeName&)

namespace myClass {
	namespace sub1 {
		void print (){std::cout << "in sub1()" << std::endl;}
	}
	class A                                                   
	{                                                         
        	public:                                                           
			A();                                      
                	A(int x){a=x; std::cout << "create function" << std::endl;};
        	private:
                	int a;                                    
			DISALLOW_COPY_AND_ASSIGN(A);
	};
};

#endif
