import { Component, OnInit } from '@angular/core';
import { ApiService } from '../common/api.service';
import { pb } from '../../pb';

interface Row {
	id: number;
	name: string;
	tier: number;
	nation: number;
	battle: number;
	battleRate: number;
	battleWidth: string;
	num: number;
	numShow: string;
	rate: number;
	numWidth: string;
	rank: number;
}

function uniq(value: any, index: any, self: any) {
	return self.indexOf(value) === index;
}

@Component({
	selector: 'app-list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

	src: pb.ITank[] = [];
	li: Row[] = [];
	totalNum = '';

	idx = 1;
	key = '';
	byBattle = false;
	higher = true;

	tierList = [5, 6, 7, 8, 9, 10];
	typeList = [
		{
			id: pb.TankEnum.type.LT,
			name: 'LT',
		},
		{
			id: pb.TankEnum.type.MT,
			name: 'MT',
		},
		{
			id: pb.TankEnum.type.HT,
			name: 'HT',
		},
		{
			id: pb.TankEnum.type.TD,
			name: 'TD',
		},
		{
			id: pb.TankEnum.type.SPG,
			name: 'SPG',
		},
	];
	keyList = [
		{
			id: 'dmg',
			name: '伤害',
		},
		{
			id: 'win',
			name: '胜率',
		},
		{
			id: 'xp',
			name: '经验',
		},
		{
			id: 'frag',
			name: '击毁',
		},
		{
			id: 'spot',
			name: '点亮',
		},
		{
			id: 'survived',
			name: '幸存率',
		},
	];
	shopList = [
		{
			id: pb.TankEnum.shop.Gold,
			name: '金币车',
		},
		{
			id: pb.TankEnum.shop.Silver,
			name: '银币车',
		},
		{
			id: pb.TankEnum.shop.Premium,
			name: '特种车',
		},
	];
	nationList = [
		{
			id: pb.TankEnum.nation.C,
			name: '中',
		},
		{
			id: pb.TankEnum.nation.S,
			name: '苏',
		},
		{
			id: pb.TankEnum.nation.M,
			name: '美',
		},
		{
			id: pb.TankEnum.nation.D,
			name: '德',
		},
		{
			id: pb.TankEnum.nation.Y,
			name: '英',
		},
		{
			id: pb.TankEnum.nation.F,
			name: '法',
		},
		{
			id: pb.TankEnum.nation.I,
			name: '意',
		},
		{
			id: pb.TankEnum.nation.V,
			name: '瑞',
		},
		{
			id: pb.TankEnum.nation.B,
			name: '波',
		},
		{
			id: pb.TankEnum.nation.J,
			name: '捷',
		},
		{
			id: pb.TankEnum.nation.R,
			name: '倭',
		},
	];

	selectTier: number[] = [];
	selectType: number[] = [];
	selectNation: number[] = [];
	selectShop: number[] = [];

	constructor(
		public api: ApiService,
	) {
		this.loadSearch();
	}

	async ngOnInit(): Promise<void> {
		this.src = await this.api.list();
		this.select();
	}

	_click(arr: number[], el: number, multi = false) {
		const idx = arr.indexOf(el);
		if (idx == -1) {
			if (!multi) {
				arr.length = 0;
			}
			arr.push(el);
		} else {
			arr.splice(idx, 1);
		}
		this.select();
	}

	clickTier(id: number, ev: MouseEvent) {
		this._click(this.selectTier, id, ev.ctrlKey);
	}

	clickShop(id: any, ev: MouseEvent) {
		this._click(this.selectShop, id, ev.ctrlKey);
	}

	clickType(id: number, ev: MouseEvent) {
		this._click(this.selectType, id, ev.ctrlKey);
	}

	clickNation(t: any, ev: MouseEvent) {
		this._click(this.selectNation, t, ev.ctrlKey);
	}

	clickByBattle() {
		this.byBattle = !this.byBattle;
		this.select();
	}

	clickHigher() {
		this.higher = !this.higher;
		this.select();
	}

	clickKey(t: string) {
		this.key = t;
		this.select();
	}

	clickReset() {
		this.selectShop.length = 0;
		this.selectTier.length = 0;
		this.selectType.length = 0;
		this.selectNation.length = 0;
		this.higher = true;
		this.select();
	}

	loadSearch() {
		console.log('search', window.location.search);

		const us = new URLSearchParams(window.location.search);

		this.selectTier.length = 0;
		us.get('tier')?.split(',').filter(uniq).forEach((a) => {
			const i = parseInt(a);
			if (this.tierList.includes(i)) {
				this.selectTier.push(i);
			}
		});

		this.selectType.length = 0;
		us.get('type')?.split(',').filter(uniq).forEach((a) => {
			const i = parseInt(a);
			for (const row of this.typeList) {
				if (row.id === i) {
					this.selectType.push(i);
					break;
				}
			}
		});

		this.selectShop.length = 0;
		us.get('shop')?.split(',').filter(uniq).forEach((a) => {
			const i = parseInt(a);
			for (const row of this.shopList) {
				if (row.id === i) {
					this.selectShop.push(i);
					break;
				}
			}
		});

		this.key = this.keyList[0].id;
		const key = us.get('key');
		for (const row of this.keyList) {
			if (key === row.id) {
				this.key = key;
				break;
			}
		}

		this.byBattle = !!us.get('battle');
		this.higher = !us.get('higher');
	}

	buildURL() {
		const arg = [];

		const tier = this.selectTier.sort((a, b) => a - b).join(',');
		if (tier?.length) {
			arg.push('tier=' + tier);
		}

		const ty = this.selectType.sort((a, b) => a - b).join(',');
		if (ty?.length) {
			arg.push('type=' + ty);
		}

		const shop = this.selectShop.sort((a, b) => a - b).join(',');
		if (shop?.length) {
			arg.push('shop=' + shop);
		}

		const nation = this.selectNation.sort((a, b) => a - b).join(',');
		if (nation?.length) {
			arg.push('nation=' + nation);
		}
		console.log('nation', nation);

		if (this.key != this.keyList[0].id) {
			arg.push('key=' + this.key);
		}

		if (this.byBattle) {
			arg.push('battle=1');
		}

		if (!this.higher) {
			arg.push('higher=0');
		}

		let search = '';
		if (arg.length) {
			search = '?' + arg.join('&');
		}
		if (search !== window.location.search) {
			window.history.pushState('', '', search);
		}
	}

	select() {
		this.li.length = 0;

		let maxBattle = 0;
		let maxNum = 0;

		this.buildURL();

		this.src.forEach((v: pb.ITank) => {

			const d = this.higher ? v?.statsHigher : v?.stats;
			if (!d?.battles) {
				return;
			}

			const tier = v.base?.tier || 0;
			if (this.selectTier.length && !this.selectTier.includes(tier)) {
				return;
			}
			if (this.selectType.length && !this.selectType.includes(v?.base?.type || 0)) {
				return;
			}
			if (this.selectShop.length && !this.selectShop.includes(v?.base?.shop || 0)) {
				return;
			}

			const o = {
				id: v.base?.ID || 0,
				name: v.base?.name || '',
				tier,
				nation: v.base?.nation || 0,
				battle: d.battles,
				battleRate: 0,
				battleWidth: '0',
				num: 0,
				numShow: '',
				rate: 0,
				numWidth: '0',
			} as Row;

			this.calcNum(d, o);

			if (maxBattle < o.battle) {
				maxBattle = o.battle;
			}
			if (maxNum < o.num) {
				maxNum = o.num;
			}

			this.li.push(o);
		});

		if (this.selectNation.length) {
			this.totalNum = ` / ${this.li.length}`
			this.li = this.li.filter((v: Row) => {
				return this.selectNation.includes(v.nation);
			});
		} else {
			this.totalNum = '';
		}

		this.sort(maxBattle, maxNum);
	}

	sort(maxBattle: number, maxNum: number): void {

		this.li.sort((a: Row, b: Row) => {
			if (a.num === b.num) {
				return a.id < b.id ? 1 : -1;
			}
			return a.num < b.num ? 1 : -1;
		});

		let rank = 0;
		this.li.forEach((v: Row) => {
			rank++;
			v.rank = rank;
		});

		if (this.byBattle) {
			this.li.sort((a: Row, b: Row) => {
				if (a.battle === b.battle) {
					return a.id < b.id ? 1 : -1;
				}
				return a.battle < b.battle ? 1 : -1;
			});
		}

		this.li.forEach((v: Row) => {
			v.rate = v.num / maxNum;
			v.numWidth = (v.rate * 100).toFixed(2) + '%';
			v.battleRate = v.battle / maxBattle;
			v.battleWidth = (v.battleRate * 100).toFixed(2) + '%';
		});
	}

	calcNum(s: pb.ITankStats, r: Row): number {

		let i = 0;

		const b = s.battles || 1;

		let num = 0;

		switch (this.key) {

		case 'dmg':
			num = s?.damageDealt || 0;
			r.num = num / b;
			r.numShow = '' + Math.round(r.num);
			break;

		case 'win':
			num = s?.wins || 0;
			r.num = num / b;
			r.numShow = (r.num * 100).toFixed(2) + '%';
			break;

		case 'xp':
			num = s?.xp || 0;
			r.num = num / b;
			r.numShow = '' + Math.round(r.num);
			break;

		case 'frag':
			num = s?.frags || 0;
			r.num = num / b;
			r.numShow = (r.num).toFixed(2);
			break;

		case 'spot':
			num = s?.spotted || 0;
			r.num = num / b;
			r.numShow = (r.num).toFixed(2);
			break;

		case 'survived':
			const x2 = s?.survivedBattles || 0;
			r.num = x2 / b;
			r.numShow = (r.num * 100).toFixed(2) + '%';
			break;
		}
		return i;
	}
}
