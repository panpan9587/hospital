<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="账号:" prop="username">
          <el-input v-model="formData.username" :clearable="true"  placeholder="请输入账号" />
       </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true"  placeholder="请输入密码" />
       </el-form-item>
        <el-form-item label="用户名:" prop="full_name">
          <el-input v-model="formData.full_name" :clearable="true"  placeholder="请输入用户名" />
       </el-form-item>
        <el-form-item label="性别:" prop="sex">
          <el-input v-model="formData.sex" :clearable="true"  placeholder="请输入性别" />
       </el-form-item>
        <el-form-item label="年龄:" prop="age">
          <el-input v-model.number="formData.age" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机号:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入手机号" />
       </el-form-item>
        <el-form-item label="用户健康分,用于活动抽奖:" prop="health_score">
          <el-input v-model.number="formData.health_score" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
       </el-form-item>
        <el-form-item label="是否订阅健康资讯 0:否 1:是:" prop="subscription">
          <el-input v-model.number="formData.subscription" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户状态:" prop="status">
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
  createUser,
  updateUser,
  findUser
} from '@/api/patient/patientmgmt/patientmgmt'

defineOptions({
    name: 'UserForm'
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
            username: '',
            password: '',
            full_name: '',
            sex: '',
            age: 0,
            mobile: '',
            health_score: 0,
            email: '',
            subscription: 0,
            status: 0,
        })
// 验证规则
const rule = reactive({
               username : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               full_name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mobile : [{
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
      const res = await findUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reuser
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
               res = await createUser(formData.value)
               break
             case 'update':
               res = await updateUser(formData.value)
               break
             default:
               res = await createUser(formData.value)
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
