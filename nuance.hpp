#ifndef __NUANCE_HPP__
#define __NUANCE_HPP__

#include <cstdio>
#include <cstring>
#include <KernelApi.h>
#include <recpdf.h>
#include <RecApiPlus.h>

class nuance {
private:
    HFORMTEMPLATEPAGE *hFormTemplateArray;
    HFORMTEMPLATECOLLECTION hFormTmplCollection;
    HPAGE hPage;
    int hFormTemplateArrayLen;
    int ZoneCount;

public:
    nuance(void);
    ~nuance();

    void errMsg(RECERR rc, char* errBuff, int errBuffSize);
    void errStrMsg(const char* msg, char* errBuff, int errBuffSize);

    int Init(const char *company,
             const char *product,
             char *errBuff,
             const int errSize);

    int SetLicense(const char *licenceFile,
                   const char *oemCode,
                   char *errBuff,
                   const int errSize);

    int LoadFormTemplateLibrary(const char *templateFile,
                                char *errBuff,
                                const int errSize);

    int PreprocessImgWithTemplate(const char *imgFile,
                                  char *errBuff,
                                  const int errSize);

    int OCRImgToText(const char *imgFile,
                     const char *outputFile,
                     const char *auxDocumentFile,
                     char *errBuff,
                     const int errSize);

    int getZoneCount(void);

    int getZoneData(const int zoneID,
                    char *zoneName,
                    const int zoneNameSize,
                    char *zoneText,
                    const int zoneTextSize);

    int FreeImgWithTemplate(void);

    void Quit(void);
};

#endif
