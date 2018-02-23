package templates

//HostServiceModel data for generating the host.service.ts
type HostServiceModel struct {
	// GRPC Hosts
	LocalHost string
	TestHost  string
	DemoHost  string
	ProdHost  string
	
	// HTTP Hosts
	LocalHttpsHost string
	TestHttpsHost  string
	DemoHttpsHost  string
	ProdHttpsHost  string
}

//InternalEnvironmentServiceTemplate is used to generate the host.service.ts
const HostServiceTemplate = `
import {Injectable} from '@angular/core';
import {EnvironmentService, Environment} from '../../environment';

@Injectable()
export class HostService {
	private _host: string;
	private _httpsHost: string;

	constructor(private environmentService: EnvironmentService) {
	}

	host(): string {
		if (this._host) {
			return this._host;
		}

		switch (this.environmentService.getEnvironment()) {
			case Environment.LOCAL:
				this._host = '{{ .LocalHost }}';
				break;
			case Environment.TEST:
				this._host = '{{ .TestHost }}';
				break;
			case Environment.DEMO:
				this._host = '{{ .DemoHost }}';
				break;
			case Environment.PROD:
				this._host = '{{ .ProdHost }}';
				break;
		}
		return this._host;
	}

	httpsHost(): string {
		if (this._httpsHost) {
			return this._httpsHost;
		}

		switch (this.environmentService.getEnvironment()) {
			case Environment.LOCAL:
				this._httpsHost = '{{ .LocalHttpsHost }}';
				break;
			case Environment.TEST:
				this._httpsHost = '{{ .TestHttpsHost }}';
				break;
			case Environment.DEMO:
				this._httpsHost = '{{ .DemoHttpsHost }}';
				break;
			case Environment.PROD:
				this._httpsHost = '{{ .ProdHttpsHost }}';
				break;
		}
		return this._httpsHost;
	}

	hostWithScheme(): string {
		const scheme = this.environmentService.getEnvironment() === Environment.LOCAL ? 'http://' : 'https://';
		return scheme + this.host()
	}

	httpsHostWithScheme(): string {
		const scheme = this.environmentService.getEnvironment() === Environment.LOCAL ? 'http://' : 'https://';
		return scheme + this.httpsHost()
	}
}
`
