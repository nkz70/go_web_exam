<script setup>
import { computed, reactive } from 'vue';
import  axios  from 'axios';

let users = reactive([]);

(async() => {
  await axios.get(`http://localhost:8080/v1/users`)
  .then(res => {
    res.data.forEach(user => {
        users.push(user)
    });
  });
})()

const statusName = computed(() => {
    return (status) => {
        let statusName
        switch (status) {
            case 0:
                statusName = "未受験"
                break;
            case 1:
                statusName = "受験済み"
                break;
            case 2:
                statusName = "採点済み"
                break;
        }

        return statusName
    }
});
</script>

<template>
    <div>
        <table>
            <tr>
                <label for="last_name">氏名</label>
                <input type="text" placeholder="苗字">
                <input type="text" placeholder="名前">
            </tr>
            <tr>
                <label for="test_datetime">受験日</label>
                <input type="text" id="text-datetime">
            </tr>
            <tr>
                <label for="status">ステータス</label>
                <select>
                    <option>fix</option>
                </select>
                <input type="button" value="検索">
            </tr>
        </table>
    </div>

    <table>
        <tr>
            <th></th>
            <th></th>
            <th>状態</th>
            <th>更新日</th>
            <th>試験日</th>
            <th>氏名</th>
            <th>点数</th>
            <th>性別</th>
            <th>年齢</th>
            <th>経過</th>
        </tr>
        <tr v-for="user in users" :key="user.id">
            <td><router-link :to="{name:'UserEdit', params: {id: user.id}}">編集</router-link></td>
            <td><router-link :to="{name:'Score', params:{id: user.id}}">採点</router-link></td>
            <td>{{statusName(user.status)}}</td>
            <td>{{user.updated_at}}</td>
            <td>{{user.test_datetime}}</td>
            <td>{{user.last_name}} {{user.first_name}}</td>
            <td>{{user.score}}</td>
            <td>{{user.gender}}</td>
            <td>{{user.age}}</td>
            <td>{{user.time}}</td>
        </tr>
    </table>
</template>
