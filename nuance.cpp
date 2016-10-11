#include <cstdio>
#include <cstring>
#include <KernelApi.h>
#include <recpdf.h>

#include "nuance.h"

HFORMTEMPLATEPAGE *formTemplateArray; // TODO: transform into an array with the loaded instances.
int formTemplateArrayLen = 0;

void Quit(void) {
    kRecQuit();
}

int SetLicense(const char *licenceFile, const char *oemCode, char *errStr, int errSize) {

    RECERR rc = kRecSetLicense(licenceFile, oemCode);
    if (rc != REC_OK) {
        errMsg(rc, errStr, errSize);
        Quit();
        return -1;
    }
    return 0;
}

int InitNuance(const char *company,const char *product, char *errStr, int errSize) {

    RECERR rc = kRecInit(company, product);
    if ((rc != REC_OK) &&
        (rc != API_INIT_WARN) &&
        (rc != API_LICENSEVALIDATION_WARN)) {

        errMsg(rc, errStr, errSize);
        printf("kRecInit %s\n", errStr);
        Quit();
        return -1;
    }

    return 0;
}

int LoadFormTemplateLibrary(const char *templateFile, char *errStr, int errSize) {

    RECERR rc = kRecLoadFormTemplateLibrary(0, templateFile, TRUE, &formTemplateArray, &formTemplateArrayLen);
    if (rc != REC_OK) {
        errMsg(rc, errStr, errSize);
        Quit();
        return -1;
    }
    return 0;
}

void errMsg(RECERR rc, char* errBuff, int errBuffSize) {
    LONG ErrExt;
    char szBuff[ERR_BUFFER_SIZE];
    char ErrStr[ERR_BUFFER_SIZE];
    const char *symb = NULL;

    memset(szBuff, 0, sizeof(szBuff));
    memset(ErrStr, 0, sizeof(ErrStr));
    memset(errBuff, 0, errBuffSize);

    kRecGetLastError(&ErrExt, ErrStr, sizeof(ErrStr));

    kRecGetErrorInfo(rc, &symb);
    sprintf(szBuff + strlen(szBuff), "%s: ", symb);

    int actlen = strlen(szBuff);
    int remlen = sizeof(szBuff) - actlen - 1;

    kRecGetErrorUIText(rc, ErrExt, ErrStr, szBuff + actlen, &remlen);

    strncpy(errBuff, szBuff, errBuffSize);
}
