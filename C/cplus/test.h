#ifndef TEST_H
#define TEST_H

#include<iostream>
using namespace std;

class A                                                   
{                                                         
        public:                                                           
		A();                                      
                A(int x){a=x; cout << "create function" << endl;};                                               
                A(const A &obj){ a=obj.a; cout << "Copy function" << endl;};
        private:
                int a;                                    
};

#endif
