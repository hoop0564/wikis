#include "Gun.h"
#include "Soldier.h"
#include <iostream>
using namespace std;

void test()
{
    Soldier sanduo("xusanduo");
    sanduo.addGun(new Gun("AK47"));
    sanduo.addBulletToGun(20);
    sanduo.fire();
}

int main()
{
    cout << "this is a test4" << endl;
    test();
    return 0;
}