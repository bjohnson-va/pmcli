package templates

// SdkTemplate Stores the data used by the SDKSetupTemplate
type SdkTemplate struct {
	Name       string
	SDKName    string
	URL        string
	PythonName string
	Version    string
}

// SDKSetupTemplate is used to generate the Python SDK
const SDKSetupTemplate = `
import os
import sys

from setuptools import setup
from setuptools import find_packages

REQUIREMENTS = [
    'httplib2 >= 0.9.1',
    'oauth2client >= 2.0.1',
    'protobuf >= 3.0.0b2, != 3.0.0.b2.post1',
    'googleapis-common-protos',
    'six',
    'enum34',
    'vax',
    'vobject'
]

GRPC_EXTRAS = [
    'grpcio >= 1.0rc1'
]

if sys.version_info[:2] == (2, 7) and 'READTHEDOCS' not in os.environ:
    REQUIREMENTS.extend(GRPC_EXTRAS)

with open(os.path.join(os.path.dirname(os.path.realpath(__file__)), '{{ .SDKName }}', 'VERSION')) as f:
    version = f.read().strip()

found_packages = find_packages()
grpc_package = ['{{ .SDKName }}._generated.grpc.{{ .PythonName }}', '{{ .SDKName }}._generated.grpc.{{ .PythonName }}.{{ .Version }}']
proto_package = ['{{ .SDKName }}._generated.proto.{{ .PythonName }}', '{{ .SDKName }}._generated.proto.{{ .PythonName }}.{{ .Version}}']
found_packages = found_packages + grpc_package + proto_package
package_data = {pkg_name: ['README.md', 'VERSION'] for pkg_name in found_packages}

# include our secret JSON files:
package_data[''] = ['*.json']

setup(
    name='{{ .Name }}',
    version=version,
    description='Python API Client library for {{ .Name }}',
    author='Vendasta',
    url='{{ .URL }}',
    packages=package_data,
    install_requires=REQUIREMENTS,
    extras_require={'grpc': GRPC_EXTRAS},
    package_data=package_data,
)
`
