package certs

/*
IMPORTANT: due to golang's encapsulation directives, variables in this file MUST start with a capital letter, else
they will not be visible from other packages

reference: http://golangtutorials.blogspot.com/2011/06/structs-in-go-instead-of-classes-in.html
*/

/*
The following SSL/TLS certificate and key were generated for the embedded TLS server, by following the procedure
outlined here:

http://www.akadia.com/services/ssh_test_certificate.html

*/

var Certificate = `-----BEGIN CERTIFICATE-----
MIIDNTCCArygAwIBAgIJALlZ4nvuxhiVMAkGByqGSM49BAEwgYgxCzAJBgNVBAYT
AkNBMRUwEwYDVQQIEwxTYXNrYXRjaGV3YW4xEjAQBgNVBAcTCVNhc2thdG9vbjER
MA8GA1UEChMIVmVuZGFzdGExFTATBgNVBAMTDEJyYWQgSm9obnNvbjEkMCIGCSqG
SIb3DQEJARYVYmpvaG5zb25AdmVuZGFzdGEuY29tMB4XDTE5MTEyMDIzMDQzNFoX
DTI5MTExNzIzMDQzNFowgYgxCzAJBgNVBAYTAkNBMRUwEwYDVQQIEwxTYXNrYXRj
aGV3YW4xEjAQBgNVBAcTCVNhc2thdG9vbjERMA8GA1UEChMIVmVuZGFzdGExFTAT
BgNVBAMTDEJyYWQgSm9obnNvbjEkMCIGCSqGSIb3DQEJARYVYmpvaG5zb25AdmVu
ZGFzdGEuY29tMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEIUeqLI6IZDYy0Z+N1wkj
+gEkgULD5qhOQf9C2wpKJnFY6gAL3psGStPa1jgfWs5dJmNhrAIdjP5/F+0LLtFi
Hw9Yz9/enh0hTnDnx9vGTfoi7Jpm2SeytbkUlb/pgT2mo4HwMIHtMB0GA1UdDgQW
BBTdDJuikd9AMwU9KtcT2DufViQFLTCBvQYDVR0jBIG1MIGygBTdDJuikd9AMwU9
KtcT2DufViQFLaGBjqSBizCBiDELMAkGA1UEBhMCQ0ExFTATBgNVBAgTDFNhc2th
dGNoZXdhbjESMBAGA1UEBxMJU2Fza2F0b29uMREwDwYDVQQKEwhWZW5kYXN0YTEV
MBMGA1UEAxMMQnJhZCBKb2huc29uMSQwIgYJKoZIhvcNAQkBFhViam9obnNvbkB2
ZW5kYXN0YS5jb22CCQC5WeJ77sYYlTAMBgNVHRMEBTADAQH/MAkGByqGSM49BAED
aAAwZQIxAO5dNFgBnB/DL6TKIKyyZH+Js+9m/L6Z4JUJG4upYNexliIoKyvKHt+/
Vl7uSh+wEwIwCv2CG2aCqo1IqVnidL1mfHO9FMIC4PS7lfQMd5aoytWZ8jBsfary
FH7yrIOGO8FY
-----END CERTIFICATE-----`

var PrivateKey = `-----BEGIN EC PARAMETERS-----
BgUrgQQAIg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDAmxtmnSe6FfBkSVNlnEsYDXgjghikWgLJCzddpBxMxXKLDKuUtU33G
ZsR3uNOh+TegBwYFK4EEACKhZANiAAQhR6osjohkNjLRn43XCSP6ASSBQsPmqE5B
/0LbCkomcVjqAAvemwZK09rWOB9azl0mY2GsAh2M/n8X7Qsu0WIfD1jP396eHSFO
cOfH28ZN+iLsmmbZJ7K1uRSVv+mBPaY=
-----END EC PRIVATE KEY-----`
