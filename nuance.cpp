#include "nuance.h"

#include <cstdio>

#include <KernelApi.h>
#include <recpdf.h>

int SetLicense(const char *licenceFile, const char *oemCode) {
    RECERR rc = kRecSetLicense(licenceFile, oemCode);
    if (rc != REC_OK) {
        printf("Error code = %X\n", rc);
        return -1;
    }
    return 0;
}

int InitPDF(const char *company,const char *product) {

    printf("Initializing the Engine -- kRecInit()\n");

    RECERR rc = kRecInit(company, product);
    if ((rc != REC_OK) &&
        (rc != API_INIT_WARN) &&
        (rc != API_LICENSEVALIDATION_WARN)) {

        printf("Error code = %X\n", rc);
        return -1;
    }

    if (rc == API_INIT_WARN) {
        printf("Module initialization warning. One or more\n");
        printf("recognition modules haven't been initialized properly.\n");
        printf("For more information, see kRecGetModulesInfo()\n");
    }

    printf("Initialize RecPDF API -- rPdfInit()\n");

    rc = rPdfInit();
    if (rc != REC_OK) {
        printf("Error code = %X\n", rc);
        return -1;
    }

    return 0;
}
