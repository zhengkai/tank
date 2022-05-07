import { Injectable } from '@angular/core';
import { formatNumber } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { pb } from '../../pb';
import { tanksgg } from './tanksgg';

export type TankMap = { [key: number]: pb.ITank }

@Injectable({
	providedIn: 'root'
})
export class ApiService {

	srcDone = false;
	src: TankMap = {};

	historyCache: { [key: number]: pb.TankStatHistory } = {};
	tanksgg: { [key: number]: string } = {};

	buildTime = '';

	// historyCache

	cbList: (() => void)[] = [];

	urlGateway = '/data/list.pb';

	constructor(
		private httpClient: HttpClient,
	) {
		console.log('writen by zhengkai https://soulogic.com');

		for (const k of Object.keys(tanksgg)) {
			const v = tanksgg[+k];
			if (v && v?.startsWith('#')) {
				delete tanksgg[+k];
			}
		}
		this.tanksgg = tanksgg;

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

	historyExists(tankID: number): boolean {
		return !!this.historyCache[tankID];
	}
	async history(tankID: number): Promise<pb.TankStatHistory|null> {

		if (!tankID) {
			return null;
		}

		const url = `/data/history/${tankID}.pb${this.dateSuffix()}`

		const post = this.httpClient.get(url, {
			observe: 'response',
			responseType: 'arraybuffer',
		}).toPromise();
		const res = await post;
		if (res?.status !== 200 || !res.body) {
			return null;
		}

		const re = new Uint8Array(res.body);
		const rsp = pb.TankStatHistory.decode(re);
		this.historyCache[tankID] = rsp;

		return rsp;
	}

	async list(): Promise<void> {

		const uri = this.urlGateway + this.dateSuffix();

		const post = this.httpClient.get(uri.replace(/-/g, ''), {
			observe: 'response',
			responseType: 'arraybuffer',
		}).toPromise();
		const res = await post;

		if (res?.status !== 200 || !res.body) {
			return;
		}

		const re = new Uint8Array(res.body);
		const rsp = pb.TankList.decode(re);

		this.buildTime = rsp.buildTime;

		rsp?.list?.forEach(v => {
			const id = v?.base?.ID || 0;
			if (id) {
				if (id === 16913 && v?.base) {
					v.base.shop = 2;
				}
				this.src[id] = v;
			}
			// console.log(`${id}: "#${(v?.base?.name || '').replace(/('|")/g, '')}",`);
			if (v?.base?.shop === pb.TankEnum.shop.Premium) {
				v.base.shop = pb.TankEnum.shop.Gold;
			}
		});
	}

	dateSuffix(): string {
		return '?' + (new Date()).toISOString().substring(0, 10);
	}

	statsNum(key: string, s: pb.ITankStats): [number, string] {

		const b = s.battles || 1;

		let num = 0;
		let show = '';

		switch (key) {

		case 'dmg':
			num = (s?.damageDealt || 0) / b;
			show = '' + formatNumber(Math.round(num), 'en-US');
			break;

		case 'win':
			num = (s?.wins || 0) / b;
			show = (num * 100).toFixed(2) + '%';
			break;

		case 'xp':
			num = (s?.xp || 0) / b;
			show = '' + formatNumber(Math.round(num), 'en-US');
			break;

		case 'frag':
			num = (s?.frags || 0) / b;
			show = num.toFixed(2);
			break;

		case 'spot':
			num = (s?.spotted || 0) / b;
			show = num.toFixed(2);
			break;

		case 'survived':
			num = (s?.survivedBattles || 0) / b;
			show = (num * 100).toFixed(2) + '%';
			break;

		case 'p1':
			num = s?.p1 || 0;
			show = '' + formatNumber(Math.round(num), 'en-US');
			break;

		case 'p2':
			num = s?.p2 || 0;
			show = '' + formatNumber(Math.round(num), 'en-US');
			break;

		case 'p3':
			num = s?.p3 || 0;
			show = '' + formatNumber(Math.round(num), 'en-US');
			break;

		case 'battle':
			num = s?.battles || 0;
			show = '' + formatNumber(num, 'en-US');
			break;

		default:
			console.log('unknown key', key)
		}
		return [num, show];
	}
}
