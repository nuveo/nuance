#include "nuance.h"

#include <cstdio>

#include <KernelApi.h>
#include <recpdf.h>

int SetLicense(char *licenceFile, char *oemCode) {
    RECERR rc = kRecSetLicense(licenceFile, oemCode);
    if (rc != REC_OK)
    {
        printf("Error code = %X\n", rc);
        return -1;
    }
    return 0;
}
