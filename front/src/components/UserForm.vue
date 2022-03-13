<script setup>
import { reactive } from 'vue';
const props = defineProps({
  status:{
    type: String,
    required: true,
  }
});

const user = reactive({
    "first_name": "",
    "last_name": "",
    "gender": null,
    "age": null,
})

const emit = defineEmits(['createUser']);
const sendCreateUser = () => {
    emit('createUser', user);
};
</script>

<template>
  <div>
    <label for="first_name">性</label>
    <input type="text" id="first_name" v-model="user.first_name" />
  </div>

  <div>
    <label for="last_name">名</label>
    <input type="text" id="last_name" v-model="user.last_name" />
  </div>

  <div>
    <label for="gender">性別</label>
    <input type="radio" :value="1" v-model.number="user.gender" />男性
    <input type="radio" :value="2" v-model.number="user.gender" />女性
  </div>

  <div>
    <label for="age">年齢</label>
    <input type="number" id="age" v-model.number="user.age" />
  </div>

  <template v-if="props.status === 'create'">
      <input type="button" @click="sendCreateUser" value="登録">
  </template>
  <template v-else>
    <input type="submit" value="更新">
    <input type="submit" value="削除">
  </template>
</template>
