#include <iostream>
#include <dlfcn.h>
#include <cstdlib>
#include <unistd.h> // For sleep function

extern "C" {
    int NewGateway();
    void SayHello(int id, const char* name);
    void StopGateway(int id);
}

int main() {
    void* handle = dlopen("./go_program.so", RTLD_LAZY);
    if (!handle) {
        std::cerr << "Cannot open library: " << dlerror() << std::endl;
        return 1;
    }

    // Load the symbols
    int (*newGateway)() = (int (*)())dlsym(handle, "NewGateway");
    void (*sayHello)(int, const char*) = (void (*)(int, const char*))dlsym(handle, "SayHello");
    void (*stopGateway)(int) = (void (*)(int))dlsym(handle, "StopGateway");

    if (!newGateway || !sayHello || !stopGateway) {
        std::cerr << "Failed to load symbols: " << dlerror() << std::endl;
        dlclose(handle);
        return 1;
    }

    // Create a new Gateway instance and use it
    int id = newGateway();

    for (int i = 0; i < 10; ++i) {
        sayHello(id, "World");
        sleep(1); // Sleep for 1 second
    }

    // Stop the Gateway's goroutine
    stopGateway(id);

    // Cleanup
    dlclose(handle);
    return 0;
}