// conversion function, 转换函数

class Fraction // 分数
{
public:
	Fraction(int num, int den=1): m_numberator(num), m_denominator(den){}
	~Fraction();

	operator double() const { // return type 不用写，默认就会是double
		return (doube)(m_numberator/m_denominator);
	}
	
private:	
	int m_numberator; // 分子
	int m_denominator; // 分母
};

int main()
{
	Fraction f(3, 5);
	doube d = 4 + f; // 调用operator double()将非转为0.6
	return 0;
}