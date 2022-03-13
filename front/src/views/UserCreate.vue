<script setup>
import axios from 'axios';
import UserForm from '../components/UserForm.vue';

const onSubmit = (e) => {
  console.log(e)
    axios.post('http://localhost:8080/v1/user', e)
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
    <UserForm status="create" @createUser="onSubmit"/>
  </form>
</template>

<style>
</style>
