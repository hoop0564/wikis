#include <iostream>
using namespace std;

class Component
{
public:
	Component(){
		cout << "Component ctor\n";
	};
	virtual ~Component(){
		cout << "Component dector\n";
	};
	
};

class Base {
public:
	Base() {
		cout << "Base ctor\n";
	}
	virtual ~Base() {
		cout << "Base dector\n";
	}
protected:
	Component c;
};

class Derived : Base
{
public:
	Derived(){
		cout << "Derived ctor\n";
	};
	virtual ~Derived(){
		cout << "Derived dector\n";
	};
protected:
	//Component c;
	
};

int main() {
	Derived d;
	return 0;
}