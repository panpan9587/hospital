<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="name字段:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入name字段" />
       </el-form-item>
        <el-form-item label="cover字段:" prop="cover">
          <el-input v-model="formData.cover" :clearable="true"  placeholder="请输入cover字段" />
       </el-form-item>
        <el-form-item label="goodsType字段:" prop="goodsType">
          <el-input v-model="formData.goodsType" :clearable="true"  placeholder="请输入goodsType字段" />
       </el-form-item>
        <el-form-item label="price字段:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="status字段:" prop="status">
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
  createGoods,
  updateGoods,
  findGoods
} from '@/api/goods/goods'

defineOptions({
    name: 'GoodsForm'
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
            name: '',
            cover: '',
            goodsType: '',
            price: 0,
            status: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGoods({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reproduct
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
               res = await createGoods(formData.value)
               break
             case 'update':
               res = await updateGoods(formData.value)
               break
             default:
               res = await createGoods(formData.value)
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
