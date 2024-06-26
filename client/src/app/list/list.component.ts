import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { ApiService, TankMap } from '../common/api.service';
import { pb } from '../../pb';
import { Data } from '../matrix/matrix.component';

interface Row {
	id: number;
	name: string;
	shop: number;
	type: number;
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

const pk = ['p1', 'p2', 'p3'];

function uniq(value: any, index: any, self: any) {
	return self.indexOf(value) === index;
}

@Component({
	selector: 'app-list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.scss'],
})
export class ListComponent implements OnInit {

	@ViewChild('xclip') xclip!: ElementRef;

	init = false;

	lastCopy = 'empty';
	lastURI = '';

	srcDate = '';

	srcMap: TankMap = {};
	src: pb.ITank[] = [];
	li: Row[] = [];
	totalNum = '';

	clickY = false;

	matrixData: Data[] = [];
	matrixName = false;

	historyID = 0;
	historyKey = '';
	zeroStart = true;
	tankName = '';

	idx = 1;
	key = '';
	yKey = '';
	byBattle = false;
	higher = true;

	historyP = 'p1';

	tableType = 'table';
	tableList = [
		{
			id: 'table',
			name: '列表',
		},
		{
			id: 'matrix',
			name: '散点图',
		},
	];

	tierList = [3, 4, 5, 6, 7, 8, 9, 10];
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
	defaultKey = 'p3';
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
			id: 'p1',
			name: '1环',
		},
		{
			id: 'p2',
			name: '2环',
		},
		{
			id: 'p3',
			name: '3环',
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
			name: '幸存',
		},
	];
	historyKeyList = [
		{
			id: 'dmg',
			name: '伤害',
		},
	];
	yKeyList = [
		{
			id: 'battle',
			name: '出场',
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
		private elementRef: ElementRef,
	) {
		this.yKeyList.push(...this.keyList);
	}

	async ngOnInit(): Promise<void> {
		this.srcMap = await this.api.list();
		if (!this.init) {
			this.loadSearch();
		}
		this.srcDate = this.api.buildTime;
		this.src = Object.values(this.srcMap);
		this.init = true;
		this.select();
	}

	copyURL() {

		const search = this.buildURI();
		this.lastCopy = search;

		const el = this.xclip.nativeElement;
		el.value = `https://${window.location.host}/${search}`;
		el.select();
		el.setSelectionRange(0, 99999);
		document.execCommand('copy');
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

	clickTable(key: string) {
		this.tableType = key;
		this.updateURI();
	}

	clickTier(id: number, ev: MouseEvent) {
		this._click(this.selectTier, id, ev.ctrlKey);
	}

	clickShop(id: number, ev: MouseEvent) {
		this._click(this.selectShop, id, ev.ctrlKey);
	}

	clickType(id: number, ev: MouseEvent) {
		this._click(this.selectType, id, ev.ctrlKey);
	}

	clickNation(t: number, ev: MouseEvent) {
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
		this.historyKey = t;
		this.select();
	}

	clickHistoryKey(t: string) {
		this.historyKey = t;

		if (t === 'p1') {
			this.key = this.historyP;
		} else if (t !== 'battle') {
			this.key = t;
		}
		// this.select();
		this.updateURI();
	}

	clickZeroStart() {
		this.zeroStart = !this.zeroStart;
		this.updateURI();
	}

	clickYKey(t: string) {
		this.clickY = true;
		this.yKey = t;
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

	clickMatrixName() {
		this.matrixName = !this.matrixName;
		this.updateURI();
	}

	loadSearch() {

		const us = new URLSearchParams(window.location.search);

		this.selectTier.length = 0;
		(us.get('tier') || us.get('lv'))?.split(',').filter(uniq).forEach((a) => {
			const i = parseInt(a);
			if (this.tierList.includes(i)) {
				this.selectTier.push(i);
			}
		});

		this.selectType.length = 0;
		(us.get('type') || us.get('t'))?.split(',').filter(uniq).forEach((a) => {
			const i = parseInt(a);
			for (const row of this.typeList) {
				if (row.id === i) {
					this.selectType.push(i);
					break;
				}
			}
		});

		this.selectNation.length = 0;
		us.get('n')?.split(',').filter(uniq).forEach((a) => {
			const i = parseInt(a);
			for (const row of this.nationList) {
				if (row.id === i) {
					this.selectNation.push(i);
					break;
				}
			}
		});

		this.selectShop.length = 0;
		(us.get('shop') || us.get('s'))?.split(',').filter(uniq).forEach((a) => {
			const i = parseInt(a);
			for (const row of this.shopList) {
				if (row.id === i) {
					this.selectShop.push(i);
					break;
				}
			}
		});

		this.key = this.defaultKey;
		this.historyKey = this.defaultKey;
		const key = us.get('key') || us.get('k');
		for (const row of this.keyList) {
			if (key === row.id) {
				this.key = key;
				this.historyKey = key;
				break;
			}
		}

		this.yKey = 'win';
		if (this.key === 'win') {
			this.yKey = 'battle';
		}
		const ykey = us.get('y');
		for (const row of this.yKeyList) {
			if (ykey === row.id) {
				this.yKey = ykey;
				break;
			}
		}

		const tb = us.get('tb');
		for (const row of this.tableList) {
			if (tb === row.id) {
				this.tableType = tb;
				break;
			}
		}

		this.byBattle = !!(us.get('battle') || us.get('b'));
		this.higher = !(us.get('higher') || us.get('h'));
		this.matrixName = us.get('mn') !== '0';
		this.zeroStart = us.get('hz') !== '0';
		if (us.get('hb') === '1') {
			this.historyKey = 'battle';
		}
		this.history(us.get('hi'));
	}

	buildURI() {
		const arg = [];

		const tier = this.selectTier.sort((a, b) => a - b).join(',');
		if (tier?.length) {
			arg.push('lv=' + tier);
		}

		const ty = this.selectType.sort((a, b) => a - b).join(',');
		if (ty?.length) {
			arg.push('t=' + ty);
		}

		const shop = this.selectShop.sort((a, b) => a - b).join(',');
		if (shop?.length) {
			arg.push('s=' + shop);
		}

		const nation = this.selectNation.sort((a, b) => a - b).join(',');
		if (nation?.length) {
			arg.push('n=' + nation);
		}

		if (this.key != this.defaultKey) {
			arg.push('k=' + this.key);
		}

		if (this.key === 'win') {
			if (this.yKey !== 'battle') {
				arg.push('y=' + this.yKey);
			}
		} else {
			if (this.yKey !== 'win') {
				arg.push('y=' + this.yKey);
			}
		}

		if (this.tableType !== 'table') {
			arg.push('tb=' + this.tableType);
		}

		if (this.byBattle) {
			arg.push('b=1');
		}

		if (!this.higher) {
			arg.push('h=0');
		}

		if (this.historyID) {
			arg.push('hi=' + this.historyID);
		}

		if (this.historyKey === 'battle') {
			arg.push('hb=1');
		}

		if (!this.zeroStart) {
			arg.push('hz=0');
		}

		if (!this.matrixName) {
			arg.push('mn=0');
		}

		let search = '';
		if (arg.length) {
			search = '?' + arg.join('&');
		}
		return search;
	}

	updateURI() {
		const search = this.buildURI();
		this.lastURI = search;
		if (search !== window.location.search) {
			window.history.pushState('', '', search);
		}
	}

	select() {
		this.li.length = 0;

		let maxBattle = 0;
		let maxNum = 0;

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
				shop: v.base?.shop || 0,
				tier,
				nation: v.base?.nation || 0,
				type: v.base?.type || 0,
				battle: d.battles,
				battleRate: 0,
				battleWidth: '0',
				num: 0,
				numShow: '',
				rate: 0,
				numWidth: '0',
			} as Row;

			[o.num, o.numShow] = this.api.statsNum(this.key, d);

			if (maxBattle < o.battle) {
				maxBattle = o.battle;
			}
			if (maxNum < o.num) {
				maxNum = o.num;
			}

			this.li.push(o);
		});

		this.sort(maxBattle, maxNum);

		if (this.selectNation.length) {
			this.totalNum = ` / ${this.li.length}`;
			this.li = this.li.filter((v: Row) => {
				return this.selectNation.includes(v.nation);
			});
		} else {
			this.totalNum = '';
		}

		if (this.li.length) {
			this.matrix();
		}
		this.updateURI();
	}

	matrix() {

		if (this.key === this.yKey || !this.clickY) {
			if (this.key === 'win') {
				this.yKey = 'battle';
			} else {
				this.yKey = 'win';
			}
		}

		this.matrixData = this.li.map((v) => {
			return {
				nx: this.matrixAxis(this.yKey, v.id),
				ny: this.matrixAxis(this.key, v.id),
				type: v.type,
				id: v.id,
				x: 0,
				y: 0,
			} as Data;
		});
	}

	matrixAxis(key: string, id: number): number {

		const t = this.srcMap[id];
		if (!t) {
			return 0;
		}

		const st = this.higher ? t.statsHigher : t.stats;
		if (!st) {
			return 0;
		}

		const [num] = this.api.statsNum(key, st);
		return num;
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

	history(id: number|string|null): void {
		this.tankName = '';
		this.historyID = 0;
		if (!id) {
			return;
		}
		if (typeof id != 'number') {
			id = +id;
		}
		const name = this.srcMap[id]?.base?.name;
		if (!name) {
			return;
		}
		if (pk.includes(this.key)) {
			this.historyP = this.key;
		}
		this.historyKeyListReset();
		this.tankName = name;
		this.historyID = id;
	}

	historyKeyListReset() {
		this.historyKeyList.length = 0;
		this.yKeyList.forEach(v => {
			if (v?.id === 'p1') {
				this.historyKeyList.push({
					id: this.historyP,
					name: '1/2/3 环',
				});
			} else if (!pk.includes(v?.id)) {
				this.historyKeyList.push(v);
			}
		});
	}

	historyClick(id: number) {
		if (!id) {
			if (this.historyKey === 'battle') {
				this.byBattle = true;
			}
			this.select();
		}
		this.history(id);
		this.updateURI();
	}
}
