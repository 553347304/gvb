<template>
	<div id="echarts_calendar" style="height: 150px"></div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts';
import {onMounted} from "vue";
import {TypeArticle} from "@/api/article_api";
import router from "@/router";
import {Message} from "@arco-design/web-vue";

type Props = { data: TypeArticle.Calendar, theme: boolean }
const props = defineProps<Props>()

onMounted(() => {
	let bg = {
		color: "#000000",
		border: "#fff",
		inRange: [
			"#ebedf0", "#c6e48b", "#7bc96f", "#32af4a",
			"#1a792c", "#0f5e1e", "#0f491a", "#02340c"
		],
	}
	if (!props.theme) {
		bg = {
			color: "#5a5a5a",
			border: "#222429",
			inRange: [
				"#404148", "#c6e48b", "#7bc96f", "#32af4a",
				"#1a792c", "#0f5e1e", "#0f491a", "#02340c"
			]
		}
	}
	const chartDom = document.getElementById('echarts_calendar')!;
	const myChart = echarts.init(chartDom, null, {locale: "zh"});
	let option: echarts.EChartsOption = {
		tooltip: {
			padding: 10,
			backgroundColor: "#555",
			borderColor: "#777",
			borderWidth: 1,
			// formatter: function (e) {
			// 	e = e.value;
			// 	return '<div style="font-size: 14px; color: white">' + e[0] + "：" + e[1] + "</div>"
			// }
		},
		visualMap: {
			show: !0,
			showLabel: !0,
			categories: [0, 1, 2, 3, 4, 5, 6, 7],
			calculable: !0,
			inRange: {
				symbol: "rect",
				color: bg.inRange
			},  // 对应上面的四个值
			itemWidth: 12,
			itemHeight: 12,
			orient: "horizontal",
			left: "center",
			bottom: 0,
			textStyle: {
				color: bg.color
			}
		},
		calendar: [{
			top: 25,
			left: "center",
			range: props.data.date,  // 时间范围 ["2023-10-22", "2024-10-22"]
			cellSize: [13, 13],
			splitLine: {show: !1},
			itemStyle: {color: "#196127", borderColor: bg.border, borderWidth: 2},
			yearLabel: {show: !1},
			monthLabel: {
				nameMap: "cn",
				fontSize: 11,
				color: bg.color
			},
			dayLabel: {
				nameMap: "cn",
				fontSize: 11,
				color: bg.color
			}
		}],
		series: [
			{
				type: "heatmap",
				coordinateSystem: "calendar",
				calendarIndex: 0,
				data: props.data.total,	// [["2023-10-22",0], ["2023-10-22",0]]
			}
		],
	};

	myChart.on("click", function (params) {
		const value = params.value as [string, number]
		Message.success(value[0])
		if (value[1] === 0) return
		router.push({query: {date: value[0]}})

	})
	option && myChart.setOption(option);
})


</script>

<style scoped lang="scss">

</style>