<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="用户名称" prop="userId" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.userId,scope.row.userId) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" label="appointmentId字段" prop="appointmentId" width="120" />
        <el-table-column align="left" label="医生名称" prop="doctorId" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.doctorId,scope.row.doctorId) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" label="科室名称" prop="officeId" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.officeId,scope.row.officeId) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" label="预约时间" prop="shiftId" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.shiftId,scope.row.shiftId) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" label="患者姓名" prop="realName" width="120" />
        <el-table-column align="left" label="患者手机号" prop="mobile" width="120" />
        <el-table-column align="left" label="身份证号" prop="idNumber" width="120" />
        <el-table-column align="left" label="症状详情" prop="symptoms" width="120" />
        <el-table-column align="left" label="诊断详情" prop="diagnosis" width="120" />
        <el-table-column align="left" label="处方" prop="prescription" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateAttendingphysicianFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="用户名称:"  prop="userId" >
            <el-select  v-model="formData.userId" placeholder="请选择用户名称" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="appointmentId字段:"  prop="appointmentId" >
              <el-input v-model.number="formData.appointmentId" :clearable="true" placeholder="请输入appointmentId字段" />
            </el-form-item>
            <el-form-item label="doctorId字段:"  prop="doctorId" >
            <el-select  v-model="formData.doctorId" placeholder="请选择doctorId字段" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.doctorId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="officeId字段:"  prop="officeId" >
            <el-select  v-model="formData.officeId" placeholder="请选择officeId字段" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.officeId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="shiftId字段:"  prop="shiftId" >
            <el-select  v-model="formData.shiftId" placeholder="请选择shiftId字段" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.shiftId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="realName字段:"  prop="realName" >
              <el-input v-model="formData.realName" :clearable="true"  placeholder="请输入realName字段" />
            </el-form-item>
            <el-form-item label="mobile字段:"  prop="mobile" >
              <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入mobile字段" />
            </el-form-item>
            <el-form-item label="idNumber字段:"  prop="idNumber" >
              <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入idNumber字段" />
            </el-form-item>
            <el-form-item label="symptoms字段:"  prop="symptoms" >
              <el-input v-model="formData.symptoms" :clearable="true"  placeholder="请输入symptoms字段" />
            </el-form-item>
            <el-form-item label="diagnosis字段:"  prop="diagnosis" >
              <el-input v-model="formData.diagnosis" :clearable="true"  placeholder="请输入diagnosis字段" />
            </el-form-item>
            <el-form-item label="prescription字段:"  prop="prescription" >
              <el-input v-model="formData.prescription" :clearable="true"  placeholder="请输入prescription字段" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
    getAttendingphysicianDataSource,
  createAttendingphysician,
  deleteAttendingphysician,
  deleteAttendingphysicianByIds,
  updateAttendingphysician,
  findAttendingphysician,
  getAttendingphysicianList
} from '@/api/appointmentmgmt/registermgmt/attendingphysicians'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Attendingphysician'
})

// 自动化生成的字典（可能为空）以及字段
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
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getAttendingphysicianDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAttendingphysicianList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteAttendingphysicianFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteAttendingphysicianByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAttendingphysicianFunc = async(row) => {
    const res = await findAttendingphysician({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reattendingphysician
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAttendingphysicianFunc = async (row) => {
    const res = await deleteAttendingphysician({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
