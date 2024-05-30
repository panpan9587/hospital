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
    //   { id: 'temperature', title: '心率 ', data: [22, 23, 21, 24, 22] },
    //   { id: 'humidity', title: '心电 ', data: [55, 57, 56, 54, 55] },
    //   { id: 'operating-status', title: '运行状态', data: [1, 0, 1, 1, 1], categories: ['正常', '异常'] },
    //   { id: 'voltage', title: '血压 ', data: [220, 218, 219, 221, 220] }
    // ]);
    const charts = ref([
      { id: 'temperature', title: '心率', data: [22, 23, 21, 24, 22] },
      { id: 'humidity', title: '心电', data: [55, 57, 56, 54, 55] },
      { id: 'operating-status', title: '运行状态', data: [1, 0, 1, 1, 1], categories: ['正常', '异常'] },
      { id: 'voltage', title: '血压', data: [220, 218, 219, 221, 220] },
      // 新增图表1
      { id: 'respiration-rate', title: '呼吸频率', data: [16, 17, 16, 18, 17] },
      // 新增图表2
      { id: 'body-temperature', title: '体温', data: [36.5, 36.7, 36.6, 36.8, 36.6] },
      // 新增图表3
      { id: 'oxygen-saturation', title: '血氧饱和度', data: [98, 97, 99, 98, 97] }
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