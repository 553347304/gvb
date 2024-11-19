<template>
	<Gvb_card class="gvb_statistics_view" title="数据统计">
		<div id="login_data"/>
	</Gvb_card>
</template>

<script setup lang="ts">
import * as echarts from 'echarts';
import {onMounted} from "vue";
import Gvb_card from "@/components/admin/card/gvb_card.vue";
import {TypeDate} from "@/api/date_api";

type Props = { data: TypeDate.List }
const props = defineProps<Props>()
onMounted(() => {
	const chartDom = document.getElementById('login_data')!;
	const myChart = echarts.init(chartDom);

	let bg = getComputedStyle(document.body).getPropertyValue("--bg")
	let color_text_1 = getComputedStyle(document.body).getPropertyValue("--color_text_1")

	let option: echarts.EChartsOption = {
		color: ["red", "yellow"],
		title: {text: '7日用户数据', textStyle: {color: color_text_1},},
		tooltip: {
			trigger: 'axis',
			axisPointer: {
				type: 'cross',
				label: {
					backgroundColor: color_text_1,
				}
			}
		},
		legend: {
			textStyle: {color: color_text_1},
			data: ['登录', '注册']
		},
		grid: {
			left: '3%',
			right: '4%',
			bottom: '3%',
			containLabel: true
		},
		xAxis: [
			{
				type: 'category',
				boundaryGap: false,
				data: props.data.date_list,

			}
		],
		yAxis:
			{
				type: 'value',
				splitLine: {
					lineStyle: {
						color: bg,
					}
				}
			},
		series: [
			{
				name: '登录',
				type: 'line',
				// stack: 'Total',
				areaStyle: {},
				smooth: true,
				emphasis: {
					focus: 'series'
				},
				label: {show: true, position: "top"},
				data: props.data.login_data,
			},
			{
				name: '注册',
				type: 'line',
				// stack: 'Total',
				areaStyle: {},
				emphasis: {
					focus: 'series'
				},
				data: props.data.sign_data,
			},
		]
	};

	option && myChart.setOption(option);
})
</script>

<style scoped lang="scss">
.gvb_statistics_view {
	margin-top: 20px;
	position: relative;

	.select {
		position: absolute;
		right: 10px;
		top: 70px;
		z-index: 1;
		width: 160px;
	}

	#login_data {
		height: 300px;
	}

}
</style>