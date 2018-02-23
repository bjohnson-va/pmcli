package templates

//NPMIgnoreTemplate is used to generate npmignore file
const NPMIgnoreTemplate = `tsconfig.json
.npmignore

**/*.ts
!**/*.d.ts

*.tgz
package/
`
