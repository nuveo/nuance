#ifndef __OMNIPAGEC_H__
#define __OMNIPAGEC_H__

#define ERR_BUFFER_SIZE 1024
#define CODEPAGE_BUFFER_SIZE 254
#define OUTPUTFMT_BUFFER_SIZE 254

#ifdef __cplusplus
extern "C" {
#endif


void RecSample(void);



typedef void* omnipagePtr;

omnipagePtr omnipageNew(void);
void omnipageFree(omnipagePtr h);

int omnipageSetLicense(omnipagePtr n,
                     const char *licenceFile,
                     const char *oemCode,
                     char *errBuff,
                     const int errSize);

void omnipageQuit(omnipagePtr n);

int omnipageInit(omnipagePtr n,
               const char *company,
               const char *product,
               char *errBuff,
               const int errSize);

int omnipageLoadFormTemplateLibrary(omnipagePtr n,
                                  const char *templateFile,
                                  char *errBuff,
                                  const int errSize);

int omnipagePreprocessImgWithTemplate(omnipagePtr n,
                                    const char *imgFile,
                                    char *errBuff,
                                    const int errSize);

int omnipageGetZoneCount(omnipagePtr n);

int omnipageGetZoneData(omnipagePtr n,
                      const int zoneID,
                      char *zoneName,
                      const int zoneNameSize,
                      char *zoneText,
                      const int zoneTextSize);

void omnipageFreeImgWithTemplate(omnipagePtr n);

int omnipageOCRImgToFile(omnipagePtr n,
    const char *imgFile,
    const char *outputFile,
    const int nPage,
    const char *auxDocumentFile,
    char *errBuff,
    const int errSize);

int omnipageOCRImgToTextFile(omnipagePtr n,
                       const char *imgFile,
                       const char *outputFile,
                       const int nPage,
                       const char *auxDocumentFile,
                       char *errBuff,
                       const int errSize);

int omnipageSetLanguagePtBr(omnipagePtr n,
                          char *errBuff,
                          const int errSize);

int omnipageCountPages(omnipagePtr n,
                     const char *imgFile,
                     int *nPages,
                     char *errBuff,
                     const int errSize);

int omnipageSetCodePage(omnipagePtr n,
                      const char *codePage,
                      char *errBuff,
                      const int errSize);

int omnipageSetOutputFormat(omnipagePtr n,
                          const char *outputFormat,
                          char *errBuff,
                          const int errSize);


#ifdef __cplusplus
}
#endif
#endif
