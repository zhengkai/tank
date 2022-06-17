import { Injectable } from '@angular/core';
import { formatNumber } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { pb } from '../../pb';
import { ZCache } from './cache';

export type TankMap = { [key: number]: pb.ITank }
export type IDMap = { [key: number]: pb.TankAlias }

// 数量太少而不展示的车
const ignoreList = [
	48641, // 252工程U 防卫者
	16913, // 百运
];

@Injectable({
	providedIn: 'root',
})
export class ApiService {

	src: TankMap = {};

	historyCache: { [key: number]: ZCache } = {};
	tanksgg: { [key: number]: string } = {};
	skill4ltu: { [key: number]: string } = {};
	icon: { [key: number]: string } = {};

	buildTime = '';

	// historyCache

	cbList: (() => void)[] = [];

	cacheTank: ZCache|null = null;

	cacheID: ZCache|null = null;
	mapID: { [key: number]: pb.TankAlias } = {};

	urlGateway = '/data/list.pb';

	constructor(
		public hc: HttpClient,
	) {
		console.log('writen by zhengkai https://soulogic.com');

		this.cacheTank = new ZCache(hc, (ab: Uint8Array) => {
			this.decodeList(ab);
		}, '/data/list.pb');

		this.cacheID = new ZCache(hc, (ab: Uint8Array) => {
			this.decodeID(ab);
		}, '/data/id.pb');
	}

	decodeList(ab: Uint8Array) {
		const rsp = pb.TankList.decode(ab);
		if (!rsp) {
			return;
		}
		this.buildTime = rsp.buildTime;
		rsp.list?.forEach(v => {
			const id = v?.base?.ID || 0;
			if (!id || ignoreList.includes(id)) {
				return;
			}
			if (id === 16913 && v?.base) {
				v.base.shop = 2;
			}
			this.src[id] = v;
			if (v?.base?.shop === pb.TankEnum.shop.Premium) {
				v.base.shop = pb.TankEnum.shop.Gold;
			}
		});
	}

	decodeID(ab: Uint8Array) {
		const rsp = pb.TankAliasList.decode(ab);
		rsp?.list?.forEach((v) => {
			const id = v?.ID;
			if (!id) {
				return;
			}
			this.mapID[id] = pb.TankAlias.fromObject(v);

			const s = v?.tanksgg || '';
			if (s) {
				this.tanksgg[id] = s;
			}

			const s4 = v?.skill4ltu || '';
			if (s4.length) {
				this.skill4ltu[id] = s4;
			}

			const icon = v?.wiki || '';
			this.icon[id] = icon;
		});
	}

	async list(): Promise<TankMap> {
		await this.cacheTank?.load();
		return this.src;
	}

	async id(): Promise<void> {
		await this.cacheID?.load();
	}

	historyExists(tankID: number): boolean {
		return !!this.historyCache[tankID];
	}
	async history(tankID: number): Promise<pb.TankStatHistory|null> {
		if (!tankID) {
			return null;
		}
		let c = this.historyCache[tankID];
		if (!c) {
			const url = `/data/history/${tankID}.pb`;
			c = new ZCache(this.hc, (ab: Uint8Array) => {
				return pb.TankStatHistory.decode(ab);
			}, url);
			this.historyCache[tankID] = c;
		}
		const rsp = (await c.load()) as pb.TankStatHistory;
		return rsp || null;
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
				console.log('unknown key', key);
		}
		return [num, show];
	}
}
