#ifndef __NUANCE_H__
#define __NUANCE_H__

#define ERR_BUFFER_SIZE 1024

#ifdef __cplusplus
extern "C" {
#endif

void errMsg(RECERR rc, char* errBuff, int errBuffSize);
void Quit(void);
int SetLicense(const char *licenceFile, const char *oemCode, char *errStr, int errSize);
int InitNuance(const char *company,const char *product, char *errStr, int errSize);
int LoadFormTemplateLibrary(const char *templateFile, char *errStr, int errSize);

#ifdef __cplusplus
}
#endif
#endif
