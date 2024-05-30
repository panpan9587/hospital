<template>
  <div class="dashboard">
    <div class="chart-container" v-for="(chart, index) in charts" :key="index">
      <div :id="chart.id" :style="{ height: '200px', width: '100%' }"></div>
    </div>
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';
import * as echarts from 'echarts';

export default {
  name: 'Dashboard',
  setup() {
    // const charts = ref([
    //   { id: 'temperature', title: '温度 (°C)', data: [22, 23, 21, 24, 22] },
    //   { id: 'humidity', title: '湿度 (%)', data: [55, 57, 56, 54, 55] },
    //   { id: 'operating-status', title: '运行状态', data: [1, 0, 1, 1, 1], categories: ['正常', '异常'] },
    //   { id: 'voltage', title: '透析结果 (V)', data: [220, 218, 219, 221, 220] }
    // ]);
    const charts = ref([
      { id: 'temperature', title: '设备温度 (°C)', data: [35, 36, 35, 37, 36] },
      { id: 'humidity', title: '设备湿度 (%)', data: [40, 42, 41, 43, 42] },
      { id: 'operating-status', title: '运行状态', data: [1, 0, 1, 1, 1], categories: ['正常', '异常'] },
      { id: 'voltage', title: '透析结果 (V)', data: [220, 218, 219, 221, 220] },
      // 新增图表1
      { id: 'power-usage', title: '电力消耗 (kWh)', data: [150, 155, 152, 158, 154] },
      // 新增图表2
      { id: 'scan-time', title: '扫描时间 (秒)', data: [30, 28, 32, 31, 29] },
      // 新增图表3
      { id: 'tube-current', title: '管电流 (mA)', data: [200, 210, 205, 215, 210] }
    ]);
    onMounted(() => {
      charts.value.forEach(chart => {
        const chartInstance = echarts.init(document.getElementById(chart.id));
        const option = {
          title: {
            text: chart.title,
            left: 'center'
          },
          xAxis: {
            type: 'category',
            data: ['09:07', '09:19', '09:31', '09:43', '09:55']
          },
          yAxis: {
            type: chart.id === 'operating-status' ? 'category' : 'value',
            data: chart.id === 'operating-status' ? chart.categories : undefined
          },
          series: [{
            data: chart.data,
            type: 'line'
          }]
        };
        chartInstance.setOption(option);
      });
    });

    return {
      charts
    };
  }
};
</script>

<style scoped>
.chart-container {
  margin: 20px 0;
}
</style>