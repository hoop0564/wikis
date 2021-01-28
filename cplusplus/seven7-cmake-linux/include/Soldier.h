#pragma once
#include <string>
#include "Gun.h"
using namespace std;

class Soldier
{
private:
    string _name;
    Gun *_prt_gun;

public:
    Soldier(string name);
    ~Soldier();
    void addGun(Gun *ptr_gun);
    void addBulletToGun(int num);
    bool fire();
};
