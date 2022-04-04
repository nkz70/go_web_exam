<script setup>
import { computed, ref, reactive } from 'vue';
import  axios  from 'axios';
import FormText from '../components/FormText.vue';
import FormCheck from '../components/FormCheck.vue';
import FormRadio from '../components/FormRadio.vue';
import FormSelect from '../components/FormSelect.vue';

const form = ref('text')
const id = ref(1)
const question = reactive({
    short_name: "",
    content: "",
})


const tabs = {
    text: FormText,
    check: FormCheck,
    radio: FormRadio,
    select: FormSelect
}

const forms = reactive([
    {
        id: id.value,
        label_position: null,
        label: "",
        answer: "",
        form_type: 1,
        choices: [
            "",
        ]
    },
]);

const addForm = () => {
    id.value++
    let form = {
            id: id.value,
            label_position: "",
            label: "",
            answer: "",
            choices:[
                ""
            ]
        }
    forms.push(form)
}

const submitForm = () => {
    console.log(forms)
    axios.post("http://localhost:8080/v1/question",{
        content: question.content,
        short_name: question.short_name,
        number: 1,
        question_answers:forms
    }).then(res =>
        console.log("せいこう")
    )
}

const tab = computed(() => tabs[form.value])
</script>

<template>
  <form @submit.prevent>
      <div>
          <label for="">略称</label>
          <input type="text" v-model="question.short_name">
      </div>
      <div>
          <label for="">問題文</label>
          <textarea id="" cols="30" rows="10" v-model="question.content"></textarea>
      </div>
      <div>
          <label for="">画像</label>
          <input type="text">
      </div>
      <div>
          <button @click="form = 'text'">text</button>
          <button @click="form = 'check'">check</button>
          <button @click="form = 'radio'">radio</button>
          <button @click="form = 'select'">select</button>
      </div>
      <button @click="addForm()">追加</button>
        <component
        :form="form"
        v-for="form in forms"
        :is="tab"
      ></component>
      <input type="button" @click="submitForm()" value="登録">
  </form>
</template>
