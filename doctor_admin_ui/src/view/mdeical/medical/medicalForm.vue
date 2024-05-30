<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="过敏史:" prop="allergies">
          <el-input v-model="formData.allergies" :clearable="true"  placeholder="请输入过敏史" />
       </el-form-item>
        <el-form-item label="诊断:" prop="diagnosis">
          <el-input v-model="formData.diagnosis" :clearable="true"  placeholder="请输入诊断" />
       </el-form-item>
        <el-form-item label="身份证号:" prop="idCard">
          <el-input v-model="formData.idCard" :clearable="true"  placeholder="请输入身份证号" />
       </el-form-item>
        <el-form-item label="医学检查:" prop="medicalTests">
          <el-input v-model="formData.medicalTests" :clearable="true"  placeholder="请输入医学检查" />
       </el-form-item>
        <el-form-item label="药物:" prop="medications">
          <el-input v-model="formData.medications" :clearable="true"  placeholder="请输入药物" />
       </el-form-item>
        <el-form-item label="就诊人员:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入就诊人员" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊日期:" prop="visitDate">
          <el-date-picker v-model="formData.visitDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createMedical,
  updateMedical,
  findMedical
} from '@/api/mdeical/medical'

defineOptions({
    name: 'MedicalForm'
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
            allergies: '',
            diagnosis: '',
            idCard: '',
            medicalTests: '',
            medications: '',
            name: '',
            status: 0,
            visitDate: new Date(),
        })
// 验证规则
const rule = reactive({
               idCard : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMedical({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remedical
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
               res = await createMedical(formData.value)
               break
             case 'update':
               res = await updateMedical(formData.value)
               break
             default:
               res = await createMedical(formData.value)
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
