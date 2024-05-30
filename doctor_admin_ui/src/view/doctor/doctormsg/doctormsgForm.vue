<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="医生姓名:" prop="doctorName">
          <el-input v-model="formData.doctorName" :clearable="true"  placeholder="请输入医生姓名" />
       </el-form-item>
        <el-form-item label="医生年龄:" prop="age">
          <el-input v-model.number="formData.age" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:" prop="sex">
          <el-input v-model="formData.sex" :clearable="true"  placeholder="请输入性别" />
       </el-form-item>
        <el-form-item label="医生职位:" prop="position">
          <el-input v-model="formData.position" :clearable="true"  placeholder="请输入医生职位" />
       </el-form-item>
        <el-form-item label="医生标签:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true"  placeholder="请输入医生标签" />
       </el-form-item>
        <el-form-item label="医生简介:" prop="desc">
          <RichEdit v-model="formData.desc"/>
       </el-form-item>
        <el-form-item label="医生工龄:" prop="workAge">
          <el-input v-model.number="formData.workAge" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医生工作时间:" prop="workTime">
          <el-date-picker v-model="formData.workTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="科室:" prop="officeId">
        <el-select  v-model="formData.officeId" placeholder="请选择科室" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.officeId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getDoctorDataSource,
  createDoctor,
  updateDoctor,
  findDoctor
} from '@/api/doctor/doctormsg/doctormsg'

defineOptions({
    name: 'DoctorForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            doctorName: '',
            age: 0,
            sex: '',
            position: '',
            tag: '',
            desc: '',
            workAge: 0,
            workTime: new Date(),
            officeId: 0,
            status: 0,
        })
// 验证规则
const rule = reactive({
               doctorName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               age : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sex : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               position : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tag : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               desc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               workAge : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               workTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               officeId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getDoctorDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDoctor({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redoctor
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createDoctor(formData.value)
               break
             case 'update':
               res = await updateDoctor(formData.value)
               break
             default:
               res = await createDoctor(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
