#include <iostream>
using namespace std;

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
	virtual Image* clone() = 0;
	// As each subclass of Image is declared, it regitsters its prototype
	static void addPrototype(Image* image)
	{
		_prototypes[_nextSlot++] = image;
	}
private:
	// addPrototype() saves each registered prototype here
	static Image* _prototypes[10];
	static int _nextSlot;
};

Image *Image::_prototypes[];
int Image::_nextSlot;

// cpp
Image *Image::findAndClone(imageType)
{
	for (int i = 0; i < _nextSlot; i++)
	{
		if (_prototypes[i]->returnType() == imageType) {
			returnType _prototypes[i]->clone();
		}
	}
}