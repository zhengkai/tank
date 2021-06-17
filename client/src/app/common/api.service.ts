import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { pb } from '../../pb';

export type TankMap = { [key: number]: pb.ITank }

@Injectable({
	providedIn: 'root'
})
export class ApiService {

	srcDone = false;
	src: TankMap = {};

	cbList: (() => void)[] = [];

	urlGateway = '/data/list.pb';

	constructor(
		private httpClient: HttpClient,
	) {

		console.log('writen by zhengkai https://soulogic.com');

		(async () => {
			await this.list();
			this.cbList.forEach(v => {
				v();
			});
			this.cbList.length = 0;
			this.srcDone = true;
		})();
	}

	async data(): Promise<TankMap> {

		if (this.srcDone) {
			return this.src;
		}

		return new Promise((resolve, reject) => {
			this.cbList.push(() => {
				resolve(this.src);
			});
		});
	}

	async list(): Promise<void> {

		const uri = this.urlGateway + '?' + (new Date()).toISOString().substring(0, 10);

		const post = this.httpClient.get(uri.replace(/-/g, ''), {
			observe: 'response',
			responseType: 'arraybuffer',
		}).toPromise();
		const res = await post;

		let rsp: pb.TankList;
		if (res.status !== 200 || !res.body) {
			return;
		}

		const re = new Uint8Array(res.body);
		rsp = pb.TankList.decode(re);

		rsp?.list?.forEach(v => {
			const id = v?.base?.ID || 0;
			if (id) {
				this.src[id] = v;
			}
		});
	}
}
