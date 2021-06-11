import { Component, OnInit } from '@angular/core';
import { ApiService } from '../common/api.service';
import { pb } from '../../pb';

interface Row {
	id: number;
	name: string;
	tier: number;
	battle: number;
	battleRate: number;
	battleWidth: string;
	num: number;
	numShow: string;
	rate: number;
	numWidth: string;
	rank: number;
}

@Component({
	selector: 'app-list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

	src: pb.ITank[] = [];
	li: Row[] = [];

	higher = true;

	idx = 1;
	key = 'dmg';

	byBattle = false;

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
			name: '场均伤害',
		},
		{
			id: 'win',
			name: '胜率',
		},
		{
			id: 'xp',
			name: '场均经验',
		},
		{
			id: 'frag',
			name: '场均击毁',
		},
		{
			id: 'spot',
			name: '场均点亮',
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

	selectTier = [10];
	selectType = [
		pb.TankEnum.type.HT,
	];
	selectShop: pb.TankEnum.shop[] = [];

	constructor(
		public api: ApiService,
	) {
	}

	clickTier(i: number) {
		const idx = this.selectTier.indexOf(i);
		if (idx == -1) {
			this.selectTier.length = 0;
			this.selectTier.push(i);
		} else {
			this.selectTier.splice(idx, 1);
		}
		console.log('selectTier', this.selectTier);

		this.select();
	}

	clickShop(id: any) {
		const idx = this.selectShop.indexOf(id);
		if (idx == -1) {
			this.selectShop.length = 0;
			this.selectShop.push(id);
		} else {
			this.selectShop.splice(idx, 1);
		}
		console.log(this.selectShop);

		this.select();
	}

	clickType(t: any) {
		const ty = t.id;
		const idx = this.selectType.indexOf(ty);
		if (idx == -1) {
			this.selectType.length = 0;
			this.selectType.push(ty);
		} else {
			this.selectType.splice(idx, 1);
		}

		this.select();
	}

	clickByBattle() {
		this.byBattle = !this.byBattle;
		this.select();
	}

	clickKey(t: string) {
		this.key = t;

		this.select();
	}

	async ngOnInit(): Promise<void> {
		this.src = await this.api.list();
		this.select();
	}

	select() {
		this.li.length = 0;

		let maxBattle = 0;
		let maxNum = 0;

		this.src.forEach((v: pb.ITank) => {

			const d = this.higher ? v?.stats : v?.statsHigher;
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
