<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名称:" prop="userId">
        <el-select  v-model="formData.userId" placeholder="请选择用户名称" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="appointmentId字段:" prop="appointmentId">
          <el-input v-model.number="formData.appointmentId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="doctorId字段:" prop="doctorId">
        <el-select  v-model="formData.doctorId" placeholder="请选择doctorId字段" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.doctorId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="officeId字段:" prop="officeId">
        <el-select  v-model="formData.officeId" placeholder="请选择officeId字段" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.officeId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="shiftId字段:" prop="shiftId">
        <el-select  v-model="formData.shiftId" placeholder="请选择shiftId字段" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.shiftId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="realName字段:" prop="realName">
          <el-input v-model="formData.realName" :clearable="true"  placeholder="请输入realName字段" />
       </el-form-item>
        <el-form-item label="mobile字段:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入mobile字段" />
       </el-form-item>
        <el-form-item label="idNumber字段:" prop="idNumber">
          <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入idNumber字段" />
       </el-form-item>
        <el-form-item label="symptoms字段:" prop="symptoms">
          <el-input v-model="formData.symptoms" :clearable="true"  placeholder="请输入symptoms字段" />
       </el-form-item>
        <el-form-item label="diagnosis字段:" prop="diagnosis">
          <el-input v-model="formData.diagnosis" :clearable="true"  placeholder="请输入diagnosis字段" />
       </el-form-item>
        <el-form-item label="prescription字段:" prop="prescription">
          <el-input v-model="formData.prescription" :clearable="true"  placeholder="请输入prescription字段" />
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
    getAttendingphysicianDataSource,
  createAttendingphysician,
  updateAttendingphysician,
  findAttendingphysician
} from '@/api/appointmentmgmt/registermgmt/attendingphysicians'

defineOptions({
    name: 'AttendingphysicianForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            userId: 0,
            appointmentId: 0,
            doctorId: 0,
            officeId: 0,
            shiftId: 0,
            realName: '',
            mobile: '',
            idNumber: '',
            symptoms: '',
            diagnosis: '',
            prescription: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getAttendingphysicianDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAttendingphysician({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reattendingphysician
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
               res = await createAttendingphysician(formData.value)
               break
             case 'update':
               res = await updateAttendingphysician(formData.value)
               break
             default:
               res = await createAttendingphysician(formData.value)
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
