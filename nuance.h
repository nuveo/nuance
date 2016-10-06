#ifndef __NUANCE_H__
#define __NUANCE_H__

#ifdef __cplusplus
extern "C" {
#endif

int SetLicense(const char *licenseFile, const char *oemCode);
int InitPDF(const char *company,const char *product);

#ifdef __cplusplus
}
#endif
#endif
