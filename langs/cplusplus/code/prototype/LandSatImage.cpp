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
		return new LandSatImage(1);
	}

protected:
	LandSatImage(int dummy){
		_id = _count++;
	}

private:
	// Mechanism for initializing an Image subclass - this causes the 
	// default ctor to be called, which registers the subclass's prototype
	static LandSatImage _landSatImage;
	// This is only called when the private static data member is inited
	LandSatImage() {
		addPrototype(this);
	}
	// Nominal(名义上的) "state" per instance mechanism
	int _id;
	static int _count;
};

// Register the subclass's prototype
LandSatImage LandSatImage::_landSatImage;
// initialize the "state" per instance mechanism
int LandSatImage::_count = 1;