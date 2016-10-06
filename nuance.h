#ifndef __NUANCE_H__
#define __NUANCE_H__

#ifdef __cplusplus
extern "C" {
#endif

int SetLicense(char *licenseFile, char *oemCode);
int InitPDF(char *company,char *product);

#ifdef __cplusplus
}
#endif
#endif
