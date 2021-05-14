// 重载operator new，operator delete，operator new[]，operator delete

void* myAlloc(size_t size) {
	return malloc(size);
}

void myFree(void* ptr) {
	return free(ptr);
}

//它们不可用被声明于一个Namespace内
inline void* operator new(size_t size) {
	cout << "global new()\n";
	return myAlloc(size);
}

inline void* operator new[](size_t size) {
	cout << "global new[]()\n";
	return myAlloc(size);
}


inline void* operator delete(void* ptr) {
	cout << "global delete()\n";
	return myFree(size);
}

inline void* operator delete[](void* ptr) {
	cout << "global delete[]()\n";
	return myFree(size);
}
