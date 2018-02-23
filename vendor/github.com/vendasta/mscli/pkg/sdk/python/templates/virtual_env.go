package templates

// VirtualEnvTemplate Stores the name of the sdk
type VirtualEnvTemplate struct {
	Name string
}

// SDKVirtualEnvTemplate generates the virtual env readme.md
const SDKVirtualEnvTemplate = `# {{ .Name }} Python SDK

# Developing
Start by creating a new virtual environment while in the sdk/python folder by running ` + "`virtualenv --python python2.7 v-env`" + `.

If you already have a virtual environment activated, you can disable it by running ` + "`deactivate`" + `.

Next, active the specific {{ .Name }} virtual environment by running ` + "`source v-env/bin/activate`" + `.

With your virtual environment activate, we can now install our dependencies to the virtual environment. Cd into /src/ and run ` + "`pip install -e .`" + `.

If you need to add a new dependency, you must add it to the dependency list in ` + "`setup.py`" + `
`
