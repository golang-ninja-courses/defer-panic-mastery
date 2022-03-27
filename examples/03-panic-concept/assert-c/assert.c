#include <assert.h>
#include <stdlib.h>

void strcpy_strict(char *dst, char *src) {
    assert(dst && "dst string ptr is NULL");
    assert(src && "src string ptr is NULL");

    // Copying.
    while(*dst++ = *src++);
}

// gcc -o assert.out assert.c
// gcc -DNDEBUG -o assert.out assert.c
int main() {
    strcpy_strict(NULL, "hello world");
}
