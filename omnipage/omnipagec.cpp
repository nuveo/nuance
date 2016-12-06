#include "omnipage.hpp"
#include "omnipagec.h"

omnipagePtr omnipageNew(void) {
    omnipage *n = new omnipage();
    return (void*)n;
}

void omnipageFree(omnipagePtr n) {
    omnipage *ptr = (omnipage*)n;
    delete ptr;
}

void omnipageQuit(omnipagePtr n) {
    omnipage *ptr = (omnipage*)n;
    ptr->Quit();
}

int omnipageInit(omnipagePtr n,
               const char *company,
               const char *product,
               char *errBuff,
               const int errSize) {

    omnipage *ptr = (omnipage*)n;

    return ptr->Init(company,
                     product,
                     errBuff,
                     errSize);
}


int omnipageSetLicense(omnipagePtr n,
                     const char *licenceFile,
                     const char *oemCode,
                     char *errBuff,
                     const int errSize) {

    omnipage *ptr = (omnipage*)n;

    return ptr->SetLicense(licenceFile,
                           oemCode,
                           errBuff,
                           errSize);
}

int omnipageLoadFormTemplateLibrary(omnipagePtr n,
                                  const char *templateFile,
                                  char *errBuff,
                                  const int errSize) {

    omnipage *ptr = (omnipage*)n;

    return ptr->LoadFormTemplateLibrary(templateFile,
                                        errBuff,
                                        errSize);
}

int omnipagePreprocessImgWithTemplate(omnipagePtr n,
                                    const char *imgFile,
                                    char *errBuff,
                                    const int errSize) {

    omnipage *ptr = (omnipage*)n;

    return ptr->PreprocessImgWithTemplate(imgFile, errBuff, errSize);
}

int omnipageGetZoneCount(omnipagePtr n) {
    omnipage *ptr = (omnipage*)n;
    return ptr->getZoneCount();
}

int omnipageGetZoneData(omnipagePtr n,
                      const int zoneID,
                      char *zoneName,
                      const int zoneNameSize,
                      char *zoneText,
                      const int zoneTextSize) {

    omnipage *ptr = (omnipage*)n;
    return ptr->getZoneData(zoneID,
                            zoneName,
                            zoneNameSize,
                            zoneText,
                            zoneTextSize);

}

void omnipageFreeImgWithTemplate(omnipagePtr n) {
    omnipage *ptr = (omnipage*)n;
    ptr->FreeImgWithTemplate();
}

int omnipageOCRImgToFile(omnipagePtr n,
    const char *imgFile,
    const char *outputFile,
    const int nPage,
    const char *auxDocumentFile,
    char *errBuff,
    const int errSize) {

        omnipage *ptr = (omnipage*)n;
        return ptr->OCRImgToFile(imgFile,
                                 outputFile,
                                 nPage,
                                 auxDocumentFile,
                                 errBuff,
                                 errSize);

}


int omnipageOCRImgToTextFile(omnipagePtr n,
                       const char *imgFile,
                       const char *outputFile,
                       const int nPage,
                       const char *auxDocumentFile,
                       char *errBuff,
                       const int errSize) {

    omnipage *ptr = (omnipage*)n;
    return ptr->OCRImgToTextFile(imgFile,
                             outputFile,
                             nPage,
                             auxDocumentFile,
                             errBuff,
                             errSize);

}

int omnipageSetLanguagePtBr(omnipagePtr n, char *errBuff, const int errSize) {

    omnipage *ptr = (omnipage*)n;
    return ptr->SetLanguagePtBr(errBuff, errSize);
}

int omnipageCountPages(omnipagePtr n,
                     const char *imgFile,
                     int *nPages,
                     char *errBuff,
                     const int errSize) {

    omnipage *ptr = (omnipage*)n;
    return ptr->CountPages(imgFile,
                           nPages,
                           errBuff,
                           errSize);
}

int omnipageSetCodePage(omnipagePtr n,
                      const char *codePage,
                      char *errBuff,
                      const int errSize) {

    omnipage *ptr = (omnipage*)n;
    return ptr->SetCodePage(codePage,
                            errBuff,
                            errSize);
}

int omnipageSetOutputFormat(omnipagePtr n,
                          const char *outputFormat,
                          char *errBuff,
                          const int errSize) {

    omnipage *ptr = (omnipage*)n;
    return ptr->SetOutputFormat(outputFormat,
                                errBuff,
                                errSize);
}
