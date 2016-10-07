#include <cstdio>
#include <KernelApi.h>
#include <recpdf.h>

#include "nuance.h"
#include <cstring>

HFORMTEMPLATEPAGE *formTemplateArray;
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

int InitPDF(const char *company,const char *product, char *errStr, int errSize) {

    RECERR rc = kRecInit(company, product);
    if ((rc != REC_OK) &&
        (rc != API_INIT_WARN) &&
        (rc != API_LICENSEVALIDATION_WARN)) {

        errMsg(rc, errStr, errSize);
        Quit();
        return -1;
    }

    rc = rPdfInit();
    if (rc != REC_OK) {
        errMsg(rc, errStr, errSize);
        Quit();
        return -1;
    }

    return 0;
}

int LoadFormTemplateLibrary(const char *templateFile, char *errStr, int errSize) {

    RECERR rc = kRecLoadFormTemplateLibrary(0, "fgv.ftl", TRUE, &formTemplateArray, &formTemplateArrayLen);
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
