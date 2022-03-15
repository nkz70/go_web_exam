<script setup>
import { reactive, watchEffect } from 'vue';
const props = defineProps({
  status: {
    type: String,
    required: true,
  },
  user: {
    type: Object,
  },
});

let user = reactive({
    first_name: "",
    last_name: "",
    gender: null,
    age: null
})

watchEffect(() => {
    user.first_name = props.user.first_name
    user.last_name = props.user.last_name
    user.gender = props.user.gender
    user.age = props.user.age
})


const emit = defineEmits(['createUser', 'updateUser', 'deleteUser']);
const sendCreateUser = () => {
    emit('createUser', user);
};
const sendUpdateUser = () => {
    emit('updateUser', user);
};
const sendDeleteUser = () => {
    emit('deleteUser')
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
    <input type="submit" @click="sendUpdateUser" value="更新">
    <input type="submit" @click="sendDeleteUser" value="削除">
  </template>
</template>
