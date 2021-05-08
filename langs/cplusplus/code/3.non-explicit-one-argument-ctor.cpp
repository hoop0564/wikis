
class Fraction // 分数
{
public:
	// saying: two parameter(形参), one argument(实参)
	// explicit 可以出现在下面的构造函数前面
	Fraction(int num, int den=1): m_numberator(num), m_denominator(den){}
	~Fraction();

	Fraction operator+(const Fraction& f) {
		return Fraction(...);
	}
	
private:	
	int m_numberator; // 分子
	int m_denominator; // 分母
};

int main()
{
	Fraction f(3, 5);
	Fraction d2 = f + 4; // 调用non-explicit ctor 将4转为 Fraction
						 // 然后调用operator+，这些都是编译器自动完成的
	return 0;
}