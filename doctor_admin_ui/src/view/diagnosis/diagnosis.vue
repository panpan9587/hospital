<template>
  <div class="chat-wrapper">
    <div class="sidebar">
      <ul>
        <li v-for="user in users" :key="user.id" @click="selectUser(user)">
          {{ user.name }}
        </li>
      </ul>
    </div>
    <div class="chat-container">
      <div class="messages">
        <div v-for="message in messages" :key="message.id" class="message">
          <span>{{ message.text }}</span>
        </div>
      </div>
      <div class="input-container">
        <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="回复患者的消息"/>
        <button @click="sendMessage">发送</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts" name="ChatBox">
  import { ref } from 'vue';
  import { WebDisnosis} from '@/api/diagnosis/diagnosis';
import { da } from 'element-plus/es/locale';

    const users = ref([
      { id: 1, name: 'User 1' },
      { id: 2, name: 'User 2' },
      { id: 3, name: 'User 3' }
      ]);
    const selectedUser = ref(null);
    const messages = ref([]);
    const newMessage = ref('');

    const selectUser = (user) => {
     selectedUser.value = user;
     // 清空当前聊天记录
     messages.value = [];
  };

const sendMessage = () => {
  if (newMessage.value.trim() === '') return;
  // 这里接收消息处理

    async function getLoveTalk(){
 // 发请求，下面这行的写法是：连续解构赋值+重命名
    let data = await WebDisnosis({});
    console.log(data);
    // 把请求回来的字符串，包装成一个对象
    // let obj = {id:nanoid(),title}
    // 放到数组中
    // talkList.unshift(obj)
}

// const getTableData = async() => {
//   const table = await WebDisnosis({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
//   if (table.code === 0) {
//     tableData.value = table.data.list
//     total.value = table.data.total
//     page.value = table.data.page
//     pageSize.value = table.data.pageSize
//   }
// }

  //调用后端代码获取消息
  messages.value.push({
    id: Date.now(),
    text: newMessage.value
  });

  newMessage.value = '';
};

  

</script>

<style scoped>
.chat-wrapper {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 200px;
  border-right: 1px solid #ccc;
  padding: 16px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.sidebar ul {
  list-style-type: none;
  padding: 0;
}

.sidebar li {
  padding: 8px;
  cursor: pointer;
}

.sidebar li:hover {
  background-color: #f1f1f1;
}

.chat-container {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  max-width: 600px;
  width: 100%;
  position: relative; /* 固定位置 */
  bottom: -195px;
  margin: auto;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  background: white; /* 添加背景色，防止内容重叠 */
}

.messages {
  flex-grow: 1;
  overflow-y: auto;
  margin-bottom: 16px;
  max-height: 200px; /* 设置一个最大高度，以防止过多信息超出 */
}

.message {
  background: #f1f1f1;
  border-radius: 4px;
  padding: 8px;
  margin-bottom: 8px;
}

.input-container {
  display: flex;
  gap: 8px;
}

input {
  flex-grow: 1;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  padding: 8px 16px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
