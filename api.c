
#include "string.h"

int helloFromC(char *buf) {

    strcpy(buf, "Hello from my c api.");
    return strlen(buf);
}

