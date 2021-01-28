#include "Soldier.h"

Soldier::Soldier(string name)
{
    this->_name = name;
    this->_prt_gun = nullptr;
}

Soldier::~Soldier()
{
    if (this->_prt_gun == nullptr)
    {
        return;
    }
    delete this->_prt_gun;
    this->_prt_gun = nullptr;
}

void Soldier::addGun(Gun *ptr_gun)
{
    this->_prt_gun = ptr_gun;
}

void Soldier::addBulletToGun(int num)
{
    this->_prt_gun->addBullet(num);
}

bool Soldier::fire()
{
    return this->_prt_gun->shoot();
}
