<script setup>
import { ref, reactive, computed } from 'vue';

const user = reactive({
    "id": 1,
    "last_name": "山田",
    "first_name": "太郎",
    "gender": 1,
    "age": 28,
    "score": 6
})

const scores = ref([
    {
        "id": 1,
        "short_name": "数字1",
        "user_answers":[
            {
                "answer": "A",
                "correct": 0
            }
        ]
    }
])

const gender = computed(() => user.gender === 1 ? "男" : "女")

</script>

<template>
  <table>
      <tr>
        <td>{{user.last_name}}{{user.first_name}}</td>
        <td>{{gender}}</td>
        <td>{{user.age}}歳</td>
        <td>{{user.score}}点</td>
      </tr>
  </table>
  <table>
      <tr>
          <th></th>
          <th>問題</th>
          <th>正答</th>
          <th>正解</th>
      </tr>
      <tr v-for="score,index in scores" :key="score.id">
        <td>{{index + 1}}</td>
        <td>{{score.short_name}}</td>
        <td>{{score.user_answers[0].answer}}</td>
        <td>
            <input type="radio" value="2" :checked="score.user_answers[0].correct === 2">◯
            <input type="radio" value="0" :checked="score.user_answers[0].correct === 0">❌
            <input type="radio" value="1" :checked="score.user_answers[0].correct === 1">△
        </td>
      </tr>
        <input type="button" value="fix">
  </table>
</template>

<style>
</style>
