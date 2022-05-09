import { HttpClient } from '@angular/common/http';

export class ZCache {

	done = false;
	src: any = {};

	cbList: (() => void)[] = [];

	constructor(
		private hc: HttpClient,
		private decode: any,
		url: string,
	) {
		(async () => {
			const ab = await this.fetch(url);
			if (!ab) {
				return;
			}

			this.src = decode(ab);
			this.cbList.forEach(v => {
				v();
			});
			this.cbList.length = 0;
			this.done = true;
		})();
	}

	async fetch(url: string): Promise<Uint8Array|null> {
		url += '?' + (new Date()).toISOString().substring(0, 10).replace(/-/g, '');

		const post = this.hc.get(url, {
			observe: 'response',
			responseType: 'arraybuffer',
		}).toPromise();
		const res = await post;

		if (res?.status !== 200 || !res.body) {
			return null;
		}

		const re = new Uint8Array(res.body);
		return re;
	}

	async load(): Promise<any> {
		if (this.done) {
			return this.src;
		}

		// await new Promise(resolve => setTimeout(resolve, 5000));

		return new Promise((resolve, reject) => {
			this.cbList.push(() => {
				resolve(this.src);
			});
		});
	}
}
