#include <iostream>
using namespace std;
// 原型模式（Prototype Pattern）是用于创建重复的对象，同时又能保证性能。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

enum imageType
{
	LSAT,SPOT
};

class Image
{
public:
	virtual void draw() = 0;
	static Image* findAndClone(imageType);
protected:
	virtual imageType returnType() = 0;
	// 必须实现 Cloneable 接口
	virtual Image* clone() = 0;
	// As each subclass of Image is declared, it regitsters its prototype
	static void addPrototype(Image* image)
	{
		_prototypes[_nextSlot++] = image;
	}
private:
	// addPrototype() saves each registered prototype here
	// 此处仅为声明
	static Image* _prototypes[10];
	static int _nextSlot;
};

// 下面两行给出定义，分配内存！
Image *Image::_prototypes[];
int Image::_nextSlot;

// client calls this public static member function when it needs an instance 
// of an Image subclass
// 实际编码中，不用imageType而是className来找
Image *Image::findAndClone(imageType)
{
	for (int i = 0; i < _nextSlot; i++)
	{
		if (_prototypes[i]->returnType() == imageType) {
			returnType _prototypes[i]->clone();
		}
	}
}