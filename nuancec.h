#ifndef __NUANCEC_H__
#define __NUANCEC_H__

#define ERR_BUFFER_SIZE 1024

#ifdef __cplusplus
extern "C" {
#endif

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

#ifdef __cplusplus
}
#endif
#endif
