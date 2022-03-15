<script setup>
import axios from 'axios';
import { reactive } from 'vue';
import { useRoute } from 'vue-router';
import UserForm from '../components/UserForm.vue';

const route = useRoute();
const { id } = route.params;

let user = reactive({
    first_name: "",
    last_name: "",
    gender: null,
    age: null
});

(() => {
  axios.get(`http://localhost:8080/v1/user/${id}`)
  .then(res => {
    user.first_name = res.data.first_name
    user.last_name = res.data.last_name
    user.gender = res.data.gender
    user.age = res.data.age
  });
})()

const onSubmit = (e) => {
  console.log(e)
    axios.put(`http://localhost:8080/v1/user/${id}`, e)
    .then(res => {
      console.log(res)
    })
    .catch(err => {
      console.log("error", err)
    });
}

const onDelete = () => {
    axios.put(`http://localhost:8080/v1/user/${id}`)
    .then(res => {
      console.log(res)
    })
    .catch(err => {
      console.log("error", err)
    });
}
</script>

<template>
  <form>
    <UserForm status="edit" :user="user" @updateUser="onSubmit" @deleteUser="onDelete" />
  </form>
</template>

<style>
</style>
