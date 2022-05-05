import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import * as d3 from 'd3';

export interface Data {
	date: string;
	num: number;
}

@Component({
	selector: 'app-history',
	templateUrl: './history.component.html',
	styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {

	@ViewChild('box') box!: ElementRef;

	constructor() {
		this.draw();
	}

	ngOnInit(): void {
	}

	draw() {

		const aapl: Data[] = [
			{
				date: '2007-04-23',
				num: 93.24,
			},
			{
				date: '2007-04-24',
				num: 91.24,
			},
		];

		const X = d3.map(aapl, v => Date.parse(v.date));
		const Y = d3.map(aapl, v => v.num);
		const D = d3.map(aapl, v => !!v?.num);
		console.log(typeof X[0]);

		const width = 1000;
		const height = 1000;

		const marginTop = 20;
		const marginRight = 30;
		const marginBottom = 30;
		const marginLeft = 40;

		const xRange = [marginLeft, width - marginRight];
		const yRange = [height - marginBottom, marginTop];
		console.log('range', xRange, yRange);

		const xDomain = d3.extent(X);
		const yDomain = [0, d3.max(Y) as number];
		console.log('domain', xDomain, yDomain);

		const xScale = d3.scaleLinear([
			xDomain[0] as number,
			xDomain[1] as number,
		], xRange);
		const yScale = d3.scaleLinear(yDomain, yRange);
		console.log('scale', xScale, yScale);
		const xAxis = d3.axisBottom(xScale).ticks(width / 80).tickSizeOuter(0);
		const yAxis = d3.axisLeft(yScale).ticks(height / 40);

		console.log('D', D);

  // Construct a line generator.
		/*
  const line = d3.line()
      .curve(d3.curveLinear)
      .x(i => xScale(X[i]))
      .y(i => yScale(Y[i]));
	   */

	  /*
  const svg = d3.create("svg")
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [0, 0, width, height])
      .attr("style", "max-width: 100%; height: auto; height: intrinsic;");

  svg.append("g")
      .attr("transform", `translate(0,${height - marginBottom})`)
      .call(xAxis);

  svg.append("g")
      .attr("transform", `translate(${marginLeft},0)`)
      .call(yAxis)
      .call(g => g.select(".domain").remove())
      .call(g => g.selectAll(".tick line").clone()
          .attr("x2", width - marginLeft - marginRight)
          .attr("stroke-opacity", 0.1))
      .call(g => g.append("text")
          .attr("x", -marginLeft)
          .attr("y", 10)
          .attr("fill", "currentColor")
          .attr("text-anchor", "start")
          .text(yLabel));

  svg.append("path")
      .attr("fill", "none")
      .attr("stroke", color)
      .attr("stroke-width", strokeWidth)
      .attr("stroke-linecap", strokeLinecap)
      .attr("stroke-linejoin", strokeLinejoin)
      .attr("stroke-opacity", strokeOpacity)
      .attr("d", line(I));
	 */
	}
}
