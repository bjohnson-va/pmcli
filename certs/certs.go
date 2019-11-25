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
MIICEDCCAZUCCQDv7VJCVLEQxjAKBggqhkjOPQQDAjBxMQswCQYDVQQGEwJDQTEV
MBMGA1UECAwMU2Fza2F0Y2hld2FuMRIwEAYDVQQHDAlTYXNrYXRvb24xETAPBgNV
BAoMCFZlbmRhc3RhMSQwIgYJKoZIhvcNAQkBFhViam9obnNvbkB2ZW5kYXN0YS5j
b20wHhcNMTkxMTI1MjI0MDI5WhcNMjIwMjI2MjI0MDI5WjBxMQswCQYDVQQGEwJD
QTEVMBMGA1UECAwMU2Fza2F0Y2hld2FuMRIwEAYDVQQHDAlTYXNrYXRvb24xETAP
BgNVBAoMCFZlbmRhc3RhMSQwIgYJKoZIhvcNAQkBFhViam9obnNvbkB2ZW5kYXN0
YS5jb20wdjAQBgcqhkjOPQIBBgUrgQQAIgNiAAR34GqH8QRs+XqjWTUeQKpKzSm2
jLsU+GPcYhMSyDrIzZDf0ipoGW2Wpx62gDV6tMnVpYHuwkcE8ZMhn0T8qyK0Fqhe
Az0c+o547tdNF/E6rpGK5GR/ykmve2Dpw1Ukx2EwCgYIKoZIzj0EAwIDaQAwZgIx
ANE5cZTP3vP+FmlyW1Z3Y0vwFH2ZkV+q7NxyS+kkOpj7Ja53uq/Oucxv9iuWeR5J
wwIxAPayNN+RORBhbYoA8613pZa+OarUKZIiExe2xOfGWd6ViJ2u7uPtu6KOJpdF
tIuuaA==
-----END CERTIFICATE-----`

var PrivateKey = `-----BEGIN EC PARAMETERS-----
BgUrgQQAIg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDDRsEs26S1Eb15sKjUCupVRBL/5qzGhBT5FpXb8iqWjsZd4cdMRCEhc
NxpsghyRHpugBwYFK4EEACKhZANiAAR34GqH8QRs+XqjWTUeQKpKzSm2jLsU+GPc
YhMSyDrIzZDf0ipoGW2Wpx62gDV6tMnVpYHuwkcE8ZMhn0T8qyK0FqheAz0c+o54
7tdNF/E6rpGK5GR/ykmve2Dpw1Ukx2E=
-----END EC PRIVATE KEY-----`
