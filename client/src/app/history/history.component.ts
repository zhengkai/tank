import { Component, OnChanges, Input, ElementRef, ViewChild } from '@angular/core';
import { ApiService } from '../common/api.service';
import { pb } from '../../pb';
import * as d3 from 'd3';

export interface Data {
	date: Date;
	num: number;
	numShow: string;
}

@Component({
	selector: 'app-history',
	templateUrl: './history.component.html',
	styleUrls: ['./history.component.scss'],
})
export class HistoryComponent implements OnChanges {

	@Input() tankID = 0;
	@Input() higher = false;
	@Input() zeroStart = true;
	@Input() key = '';

	@ViewChild('box') box!: ElementRef;

	max: Data|null = null;
	min: Data|null = null;
	median: Data|null = null;
	loading = false;

	serial = 0;
	svg: any = null;
	tooltip: any = null;

	emptyData = false;

	constructor(
		public api: ApiService,
	) {
	}

	AfterViewInit() {
		this.ngOnChanges();
	}

	ngOnChanges() {
		if (!this.key || !this.tankID) {
			return;
		}
		this.chart();
	}

	async chart(): Promise<void> {

		this.serial++;
		const serial = this.serial;

		if (!this.api.historyExists(this.tankID)) {
			this.max = null;
			this.min = null;
			this.svg?.remove();
			this.loading = true;
		}
		const re = await this.api.history(this.tankID);
		this.loading = false;

		if (serial !== this.serial) {
			console.log('skip', serial);
			return;
		}
		const d = this.transData(re);
		this.draw(d);
	}

	transData(d: pb.TankStatHistory|null): Data[][] {

		const li = d?.list;
		if (!li?.length) {
			return [];
		}

		const full: Data[][] = [];

		let key = [this.key];
		const pk = ['p3', 'p2', 'p1'];
		if (pk.includes(this.key)) {
			key = pk;
		}

		for (const k of key) {
			const re: Data[] = [];
			for (const a of li) {
				const t = '' + a?.date;

				const st = this.higher ? a.statsHigher : a.stats;
				if (!st) {
					continue;
				}

				const [num, numShow] = this.api.statsNum(k, st);
				if (!num) {
					continue;
				}
				const o = {
					num,
					numShow,
					date: new Date(`${t.substring(0, 4)}-${t.substring(4, 6)}-${t.substring(6, 8)}`),
				} as Data;
				re.push(o);
			}
			if (re.length) {
				full.push(re);
			}
		}

		return full;
	}

	draw(data: Data[][]) {

		if (!data.length) {
			this.emptyData = true;
			return;
		}
		this.emptyData = false;

		this.text(data[0]);

		const X = d3.map(data[0], v => new Date(v.date));
		const Y = d3.map(data[0], v => v.num);
		const I = d3.range(X.length);

		const width = this.box.nativeElement.offsetWidth;

		const height = 500;

		const marginTop = 20;
		const marginRight = 60;
		const marginBottom = 30;
		const marginLeft = 80;

		const xRange = [marginLeft, width - marginRight];
		const yRange = [height - marginBottom, marginTop];

		const xDomain = d3.extent(X);

		let dy = Y;
		if (data.length > 1) {
			dy = [];
			data.forEach(col => {
				dy.push(...d3.map(col, v => v.num));
			});
		}

		const yDomain = [
			(d3.min(dy) as number),
			(d3.max(dy) as number),
		];

		if (this.zeroStart) {
			yDomain[0] = 0;
			yDomain[1] *= 1.08;
		} else {
			const diff = (yDomain[1] - yDomain[0]) / 5;
			yDomain[0] -= diff;
			yDomain[1] += diff;
		}

		const xScale = d3.scaleUtc(<[Date, Date]>xDomain, xRange);
		const yScale = d3.scaleLinear(yDomain, yRange);
		const xAxis = d3.axisBottom(xScale).ticks(width / 80).tickSizeOuter(0);
		const yAxis = d3.axisLeft(yScale).ticks(height / 80);

		const svg = d3.select(this.box.nativeElement).append('svg');
		if (this.svg) {
			this.svg.remove();
		}
		this.svg = svg;

		svg.attr('width', width)
			.attr('height', height)
			.attr('viewBox', [0, 0, width, height])
			.attr('style', 'max-width: 100%; height: auto; height: intrinsic;')
			.on('pointerenter pointermove', (event: PointerEvent) => {

				const d = xScale.invert(d3.pointer(event)[0]);
				const i = d3.bisectCenter(X, d);

				const x = Math.round(xScale(X[i])) - 0.5;
				const y = yScale(Y[i]) + 5;

				this.pointermoved(event, data.map(col => col[i]));
				this.tooltip.style('display', null)
					.attr('transform', `translate(${x},${y})`);
				this.svg.property('value', I[i]).dispatch('input', { bubbles: true });
			})
			.on('pointerleave', () => {
				this.tooltip.style('display', 'none');
				this.svg.node().value = null;
				this.svg.dispatch('input', { bubbles: true });
			});

		svg.append('g')
			.attr('transform', `translate(0,${height - marginBottom})`)
			.call(xAxis);

		svg.append('g')
			.attr('transform', `translate(${marginLeft},0)`)
			.call(yAxis)
			.call(g => g.select('.domain').remove())
			.call(g => g.selectAll('.tick line').clone()
				.attr('x2', width - marginLeft - marginRight)
				.attr('stroke-opacity', 0.1));

		for (const col of data) {

			const yt = d3.map(col, v => v.num);

			const line = d3.line()
				.curve(d3.curveLinear)
				.x((_, i) => xScale(X[i]))
				.y((_, i) => yScale(yt[i]));

			const li = line(col.map(v => [+v.date, v.num]));
			svg.append('path')
				.attr('fill', 'none')
				.attr('stroke', 'currentColor')
				.attr('stroke-width', 1.5)
				.attr('stroke-linecap', 'round')
				.attr('stroke-linejoin', 'round')
				.attr('stroke-opacity', 1)
				.attr('d', li);
		}

		this.tooltip = svg.append('g')
			.style('pointer-events', 'none')
			.style('display', 'none');

		this.tooltip
			.append('line')
			.style('stroke', 'rgba(0, 128, 128, 0.5)')
			.style('stroke-width', 3)
			.attr('x1', 0)
			.attr('y1', -1000)
			.attr('x2', 0)
			.attr('y2', 1000);
	}

	pointermoved(event: PointerEvent, r: Data[]) {

		const path = this.tooltip.selectAll('path')
			.data([,]) // eslint-disable-line
			.join('path')
			.attr('fill', 'white')
			.attr('stroke', 'black');

		const d = [r[0].date.toISOString().substring(0, 10)];
		d.push(...r.map(row => row.numShow));

		const text = this.tooltip.selectAll('text')
			.data([,]) // eslint-disable-line
			.join('text')
			.call((text: any) => text
				.selectAll('tspan')
				.data(d)
				.join('tspan')
				.attr('x', 0)
				.attr('y', (_: any, i: number) => `${i * 1.1}em`)
				.attr('font-weight', (_: any, i: number) => i ? null : 'bold')
				.text((d: string) => d));

		const { y, width: w, height: h } = text.node().getBBox();
		text.attr('transform', `translate(${-w / 2},${15 - y})`);
		path.attr('d', `M${-w / 2 - 10},5H-5l5,-5l5,5H${w / 2 + 10}v${h + 20}h-${w + 20}z`);
	}

	text(data: Data[]) {
		const sort = data.slice().sort((a, b) => {
			let d = (a?.num || 0) - (b?.num || 0);
			if (!d) {
				d = (+a?.date || 0) - (+b?.date || 0);
			}
			return d;
		});
		this.min = sort[0];
		this.median = sort[Math.round(sort.length / 2)];
		this.max = sort.pop() || null;
	}
}
