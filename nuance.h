#ifndef __NUANCE_H__
#define __NUANCE_H__

#ifdef __cplusplus
extern "C" {
#endif

void errMsg(RECERR rc, char* errStr);
void Quit(void);
int SetLicense(const char *licenseFile, const char *oemCode);
int InitPDF(const char *company,const char *product);
int LoadFormTemplateLibrary(const char *templateFile);

#ifdef __cplusplus
}
#endif
#endif
