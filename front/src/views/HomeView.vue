<script setup>
import { ref } from 'vue';

const users = [
  {
    "id": 1,
    "test_datetime": "2021/2/15",
    "last_name": "田中",
    "first_name": "一郎",
    "gender": "男",
    "time": "60分",
    "score": "10点"
  },
  {
    "id": 2,
    "test_datetime": "2021/2/15",
    "last_name": "吉田",
    "first_name": "二郎",
    "gender": "男",
    "time": "60分",
    "score": "10点"
  },
  {
    "id": 3,
    "test_datetime": "2021/2/15",
    "last_name": "山田",
    "first_name": "春子",
    "gender": "男",
    "time": "60分",
    "score": "10点"
  },
]

let refCheckedID = ref([])

const onSubmit = () => {
  console.log("Fix処理を追加する")
}

const SelectAll = () => {
    if (users.length === refCheckedID.value.length) {
      refCheckedID.value = []
    } else {
      let idList = []
      users.forEach(user => idList.push(user.id))
      refCheckedID.value = idList
    }
}
</script>

<template>
  <main>
    <form @submit.prevent="onSubmit()">
      <input type="button" value="全選択" @click="SelectAll()" />
      <input type="submit" value="Fix" id="fix" />
      <table>
        <tr>
          <th></th>
          <th></th>
          <th>受験日</th>
          <th>氏名</th>
          <th>性別</th>
          <th>経過時間</th>
          <th>点数</th>
        </tr>
        <tr v-for="user in users" :key="user.id">
          <td><RouterLink to="#">採点</RouterLink></td>
          <td><input type="checkbox" :value="user.id" v-model="refCheckedID" /></td>
          <td>{{user.test_datetime}}</td>
          <td>{{user.last_name}} {{user.first_name}}</td>
          <td>{{user.gender}}</td>
          <td>{{user.time}}</td>
          <td>{{user.score}}</td>
        </tr>
      </table>
    </form>
  </main>
</template>
