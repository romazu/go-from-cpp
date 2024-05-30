#include <iostream>
#include <dlfcn.h>
#include <cstdlib>

extern "C" {
    void SayHello(const char* name);
}

int main() {
    // Load the Go shared library
    void* handle = dlopen("./go_program.so", RTLD_LAZY);
    if (!handle) {
        std::cerr << "Cannot open shared library: " << dlerror() << std::endl;
        return EXIT_FAILURE;
    }

    // Reset errors
    dlerror();

    // Load the symbol
    typedef void (*SayHelloFunc)(const char*);
    SayHelloFunc sayHello = (SayHelloFunc) dlsym(handle, "SayHello");
    const char* dlsym_error = dlerror();
    if (dlsym_error) {
        std::cerr << "Cannot load symbol 'SayHello': " << dlsym_error << std::endl;
        dlclose(handle);
        return EXIT_FAILURE;
    }

    // Call the Go function
    sayHello("World");

    // Close the shared library
    dlclose(handle);

    return EXIT_SUCCESS;
}
