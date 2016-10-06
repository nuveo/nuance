#include <cstdio>
#include <KernelApi.h>
#include <recpdf.h>

#include "nuance.h"
#include <cstring>

HFORMTEMPLATEPAGE *formTemplateArray;
int formTemplateArrayLen = 0;

void Quit(void) {
    printf("Quit -- kRecQuit\n");
    kRecQuit();
}

int SetLicense(const char *licenceFile, const char *oemCode) {

    printf("Set License -- kRecSetLicense()\n");

    RECERR rc = kRecSetLicense(licenceFile, oemCode);
    if (rc != REC_OK) {
        char errStr[1024] = "";
        errMsg(rc, errStr);
        printf("%s\n", errStr);
        Quit();
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

        char errStr[1024] = "";
        errMsg(rc, errStr);
        printf("%s\n", errStr);
        Quit();
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
        char errStr[1024] = "";
        errMsg(rc, errStr);
        printf("%s\n", errStr);
        Quit();
        return -1;
    }

    return 0;
}

int LoadFormTemplateLibrary(const char *templateFile) {

    printf("Open template library -- kRecLoadFormTemplateLibrary()\n");

    RECERR rc = kRecLoadFormTemplateLibrary(0, "fgv.ftl", TRUE, &formTemplateArray, &formTemplateArrayLen);
    if (rc != REC_OK) {
        char errStr[1024] = "";
        errMsg(rc, errStr);
        printf("%s\n", errStr);
        Quit();
        return -1;
    }
    return 0;
}

void errMsg(RECERR rc, char* errStr) {
    LONG ErrExt;
    char szBuff[1024];
    char ErrStr[512];
    const char *symb = NULL;

    memset(szBuff, 0, sizeof(szBuff));
    memset(ErrStr, 0, sizeof(ErrStr));

    kRecGetLastError(&ErrExt, ErrStr, sizeof(ErrStr));

    kRecGetErrorInfo(rc, &symb);
    sprintf(szBuff + strlen(szBuff), "%s: ", symb);

    int actlen = strlen(szBuff);
    int remlen = sizeof(szBuff) - actlen - 1;

    kRecGetErrorUIText(rc, ErrExt, ErrStr, szBuff + actlen, &remlen);

    strcpy(errStr, szBuff);
}
