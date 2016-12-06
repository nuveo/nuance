#ifndef __OMNIPAGE_HPP__
#define __OMNIPAGE_HPP__

#include <cstdio>
#include <cstring>
#include <KernelApi.h>
#include <recpdf.h>
#include <RecApiPlus.h>

class omnipage {
private:
    HFORMTEMPLATEPAGE *hFormTemplateArray;
    HFORMTEMPLATECOLLECTION hFormTmplCollection;
    HPAGE hPage;
    int hFormTemplateArrayLen;
    int ZoneCount;

public:
    omnipage(void);
    ~omnipage();

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

    int OCRImgToTextFile(const char *imgFile,
                     const char *outputFile,
                     const int nPage,
                     const char *auxDocumentFile,
                     char *errBuff,
                     const int errSize);

    int OCRImgToFile(const char *imgFile,
        const char *outputFile,
        const int nPage,
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

    int SetLanguagePtBr(char *errBuff, const int errSize);

    int CountPages(const char *imgFile,
                   int *nPages,
                   char *errBuff,
                   const int errSize);

    int SetCodePage(const char *codePage,
                    char *errBuff,
                    const int errSize);

    int SetOutputFormat(const char *outputFormat,
                        char *errBuff,
                        const int errSize);

    void Quit(void);
};

#endif
