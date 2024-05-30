#include <iostream>
#include <dlfcn.h>
#include <unistd.h> // For sleep()

int main() {
    void* handle = dlopen("./go_program.so", RTLD_LAZY);
    if (!handle) {
        std::cerr << "Cannot open library: " << dlerror() << '\n';
        return 1;
    }

    // Load the StartGateway function
    void (*startGateway)() = (void (*)())dlsym(handle, "StartGateway");
    if (!startGateway) {
        std::cerr << "Cannot load function: " << dlerror() << '\n';
        dlclose(handle);
        return 1;
    }

    // Load the StopGateway function
    void (*stopGateway)() = (void (*)())dlsym(handle, "StopGateway");
    if (!stopGateway) {
        std::cerr << "Cannot load function: " << dlerror() << '\n';
        dlclose(handle);
        return 1;
    }

    // Start the Go goroutine
    startGateway();

    // Let the Go goroutine run for some time
    sleep(10); // Runs for 10 seconds

    // Stop the Go goroutine
    stopGateway();

    // Cleanup
    dlclose(handle);
    return 0;
}
