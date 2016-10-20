#include "nuance.hpp"
#include "nuancec.h"

nuancePtr nuanceNew(void) {
    nuance *n = new nuance();
    return (void*)n;
}

void nuanceFree(nuancePtr n) {
    nuance *ptr = (nuance*)n;
    delete ptr;
}

void nuanceQuit(nuancePtr n) {
    nuance *ptr = (nuance*)n;
    ptr->Quit();
}

int nuanceInit(nuancePtr n,
               const char *company,
               const char *product,
               char *errBuff,
               const int errSize) {

    nuance *ptr = (nuance*)n;

    return ptr->Init(company,
                     product,
                     errBuff,
                     errSize);
}


int nuanceSetLicense(nuancePtr n,
                     const char *licenceFile,
                     const char *oemCode,
                     char *errBuff,
                     const int errSize) {

    nuance *ptr = (nuance*)n;

    return ptr->SetLicense(licenceFile,
                           oemCode,
                           errBuff,
                           errSize);
}

int nuanceLoadFormTemplateLibrary(nuancePtr n,
                                  const char *templateFile,
                                  char *errBuff,
                                  const int errSize) {

    nuance *ptr = (nuance*)n;

    return ptr->LoadFormTemplateLibrary(templateFile,
                                        errBuff,
                                        errSize);
}

int nuancePreprocessImgWithTemplate(nuancePtr n,
                                    const char *imgFile,
                                    char *errBuff,
                                    const int errSize) {

    nuance *ptr = (nuance*)n;

    return ptr->PreprocessImgWithTemplate(imgFile, errBuff, errSize);
}

int nuanceGetZoneCount(nuancePtr n) {
    nuance *ptr = (nuance*)n;
    return ptr->getZoneCount();
}

int nuanceGetZoneData(nuancePtr n,
                      const int zoneID,
                      char *zoneName,
                      const int zoneNameSize,
                      char *zoneText,
                      const int zoneTextSize) {

    nuance *ptr = (nuance*)n;
    return ptr->getZoneData(zoneID,
                            zoneName,
                            zoneNameSize,
                            zoneText,
                            zoneTextSize);

}

void nuanceFreeImgWithTemplate(nuancePtr n) {
    nuance *ptr = (nuance*)n;
    ptr->FreeImgWithTemplate();
}

int nuanceOCRImgToText(nuancePtr n,
                       const char *imgFile,
                       const char *outputFile,
                       const int nPage,
                       const char *auxDocumentFile,
                       char *errBuff,
                       const int errSize) {

    nuance *ptr = (nuance*)n;
    return ptr->OCRImgToText(imgFile,
                             outputFile,
                             nPage,
                             auxDocumentFile,
                             errBuff,
                             errSize);

}

int nuanceSetLanguagePtBr(nuancePtr n, char *errBuff, const int errSize) {

    nuance *ptr = (nuance*)n;
    return ptr->SetLanguagePtBr(errBuff, errSize);
}

int nuanceCountPages(nuancePtr n,
                     const char *imgFile,
                     int *nPages,
                     char *errBuff,
                     const int errSize) {

    nuance *ptr = (nuance*)n;
    return ptr->CountPages(imgFile,
                           nPages,
                           errBuff,
                           errSize);
}
