#ifndef __OMNIPAGEC_H__
#define __OMNIPAGEC_H__

#define ERR_BUFFER_SIZE 1024
#define CODEPAGE_BUFFER_SIZE 254
#define OUTPUTFMT_BUFFER_SIZE 254

#ifdef __cplusplus
extern "C" {
#endif


void RecSample(void);



typedef void* nuancePtr;

nuancePtr nuanceNew(void);
void nuanceFree(nuancePtr h);

int nuanceSetLicense(nuancePtr n,
                     const char *licenceFile,
                     const char *oemCode,
                     char *errBuff,
                     const int errSize);

void nuanceQuit(nuancePtr n);

int nuanceInit(nuancePtr n,
               const char *company,
               const char *product,
               char *errBuff,
               const int errSize);

int nuanceLoadFormTemplateLibrary(nuancePtr n,
                                  const char *templateFile,
                                  char *errBuff,
                                  const int errSize);

int nuancePreprocessImgWithTemplate(nuancePtr n,
                                    const char *imgFile,
                                    char *errBuff,
                                    const int errSize);

int nuanceGetZoneCount(nuancePtr n);

int nuanceGetZoneData(nuancePtr n,
                      const int zoneID,
                      char *zoneName,
                      const int zoneNameSize,
                      char *zoneText,
                      const int zoneTextSize);

void nuanceFreeImgWithTemplate(nuancePtr n);

int nuanceOCRImgToFile(nuancePtr n,
    const char *imgFile,
    const char *outputFile,
    const int nPage,
    const char *auxDocumentFile,
    char *errBuff,
    const int errSize);

int nuanceOCRImgToTextFile(nuancePtr n,
                       const char *imgFile,
                       const char *outputFile,
                       const int nPage,
                       const char *auxDocumentFile,
                       char *errBuff,
                       const int errSize);

int nuanceSetLanguagePtBr(nuancePtr n,
                          char *errBuff,
                          const int errSize);

int nuanceCountPages(nuancePtr n,
                     const char *imgFile,
                     int *nPages,
                     char *errBuff,
                     const int errSize);

int nuanceSetCodePage(nuancePtr n,
                      const char *codePage,
                      char *errBuff,
                      const int errSize);

int nuanceSetOutputFormat(nuancePtr n,
                          const char *outputFormat,
                          char *errBuff,
                          const int errSize);


#ifdef __cplusplus
}
#endif
#endif
