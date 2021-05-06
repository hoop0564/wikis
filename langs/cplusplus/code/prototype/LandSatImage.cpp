class LandSatImage: public Image
{
public:
	imageType returnType() {
		return LSAT;
	}
	void draw() {
		cout << "LandSatImage::draw" << _id << endl;
	}
	// when clone() is called, call the one-argument ctor with a dummy(傀儡) arg
	Image *clone() {
		// 返回真正的类构造实现的对象！
		return new LandSatImage(1);
	}

protected:
	// dummy：傀儡，无所谓的参数，用于区分无参构造函数，实际编码中将为实际的类构造函数（有参or无参！）
	LandSatImage(int dummy){
		_id = _count++;
	}

private:
	// Mechanism for initializing an Image subclass - this causes the 
	// default ctor to be called, which registers the subclass's prototype
	static LandSatImage _landSatImage;
	// This is only called when the private static data member is inited
	// 注意此处无参构造是私有的，直接传了this
	LandSatImage() {
		addPrototype(this); // this = _landSatImage ！
	}
	// Nominal(名义上的) "state" per instance mechanism
	int _id;
	static int _count;
};

// Register the subclass's prototype
LandSatImage LandSatImage::_landSatImage;
// initialize the "state" per instance mechanism
int LandSatImage::_count = 1;