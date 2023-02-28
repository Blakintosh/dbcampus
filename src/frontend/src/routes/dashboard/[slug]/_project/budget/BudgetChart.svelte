<script lang="ts">
	// @ts-nocheck

	export let chartData: Array<number>;
	export let budget: number;

	import { Line } from 'svelte-chartjs';

	import {
		Chart as ChartJS,
		Title,
		Tooltip,
		Legend,
		LineElement,
		LinearScale,
		PointElement,
		CategoryScale,
		Filler
	} from 'chart.js';

	const skipped = (ctx, value) => ctx.p0.skip || ctx.p1.skip ? value : undefined;
	const down = (ctx, value) => ctx.p0.parsed.y > ctx.p1.parsed.y ? value : undefined;

	$: data = {
		labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
		datasets: [
			{
				label: 'Total Project Spend',
				fill: {
					above: 'rgba(185, 28, 28, 0.3)',
					below: 'rgba(34, 197, 94, 0.3)',
					target: 1
				},
				lineTension: 0.3,
				backgroundColor: 'rgba(185, 28, 28, 0.3)',
				borderColor: 'rgba(245, 158, 11, 0.5)',
				borderCapStyle: 'butt',
				borderDash: [],
				borderDashOffset: 0.0,
				borderJoinStyle: 'miter',
				pointBorderColor: 'rgb(217 119 6)',
				pointBorderWidth: 10,
				pointHoverRadius: 5,
				pointHoverBackgroundColor: 'rgb(0, 0, 0)',
				pointHoverBorderColor: 'rgba(220, 220, 220,1)',
				pointHoverBorderWidth: 2,
				pointRadius: 1,
				pointHitRadius: 10,
				data: chartData,
				segment: {
					borderColor: ctx => skipped(ctx, 'rgb(0,0,0,0.2)') || down(ctx, 'rgb(192,75,75)'),
					borderDash: ctx => skipped(ctx, [6, 6]),
				},
      			spanGaps: true
			},
			{
				label: 'Expected Project Spend',
				borderColor: 'rgb(113 113 122)',
				borderDash: [5, 5],
				data: [{
					x: 0,
					y: 0
				},
				{
					x: 'July',
					y: budget
				}]
			}
		]
	};

	const options = {
		responsive: true,
		interaction: {
			intersect: false
		},
		plugins: [
			{
				beforeRender: function (x, options) {
					var c = x.chart
					var dataset = x.data.datasets[0];
					var yScale = x.scales['primary'];
					var yPos = yScale.getPixelForValue(0);

					var gradientFill = c.ctx.createLinearGradient(0, 0, 0, c.height);
					gradientFill.addColorStop(0, 'green');
					gradientFill.addColorStop(yPos / c.height - 0.01, 'green');
					gradientFill.addColorStop(yPos / c.height + 0.01, 'red');
					gradientFill.addColorStop(1, 'red');

					var model = x.data.datasets[0]._meta[Object.keys(dataset._meta)[0]].dataset._model;
					model.backgroundColor = gradientFill;
				}
			}
		],
		scales: {
			y: {
				ticks: {
					callback: function(value, index, ticks) {
                        return '$' + value.toLocaleString("en", { minimumFractionDigits: 0, maximumFractionDigits: 2 });
                    }
				}
			}
		}
	};

	ChartJS.defaults.font.family = "Open Sans";

	ChartJS.register(Title, Tooltip, Legend, LineElement, LinearScale, PointElement, CategoryScale, Filler);
</script>

<div class="w-full">
	<Line {data} {options} />
</div>