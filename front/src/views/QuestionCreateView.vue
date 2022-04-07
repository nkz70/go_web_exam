<script setup>
import { computed, ref, reactive } from 'vue';
import  axios  from 'axios';
import FormText from '../components/FormText.vue';
import FormCheck from '../components/FormCheck.vue';
import FormRadio from '../components/FormRadio.vue';
import FormSelect from '../components/FormSelect.vue';

const text = 1
const radio = 2
const select = 3
const check = 4

const form_type = ref(text)
const id = ref(1)
const question = reactive({
    short_name: "",
    form_type: form_type.value,
    content: "",
})

const tabs = {
    [text]: FormText,
    [radio]: FormRadio,
    [select]: FormSelect,
    [check]: FormCheck
}

const forms = reactive([
    {
        id: id.value,
        label_position: null,
        label: "",
        answer: "",
        choices: [
            "",
        ]
    },
]);

const addForm = () => {
    id.value++
    let form = {
            id: id.value,
            label_position: null,
            label: "",
            answer: "",
            choices:[
                ""
            ]
        }
    forms.push(form)
}

const submitForm = () => {
    axios.post("http://localhost:8080/v1/question",{
        content: question.content,
        short_name: question.short_name,
        form_type: question.form_type,
        number: 1,
        question_answers:forms
    }).then(res =>
        console.log("せいこう")
    )
}

const tab = computed(() => {
    question.form_type = form_type.value
    return tabs[form_type.value]
})

const popForm = id => {
    const index = forms.findIndex(form => form.id === id);
    forms.splice(index, 1);
    // console.log(form.value.id)
}

const fileSelected = (e) => {
    console.log(e.target.files[0].name)
}
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
          <input type="file" @change="fileSelected">
      </div>
      <div>
          <button @click="form_type = text">text</button>
          <button @click="form_type = check">check</button>
          <button @click="form_type = radio">radio</button>
          <button @click="form_type = select">select</button>
      </div>
        <button @click="addForm()">追加</button>
        <component
        :form="form"
        v-for="form in forms"
        :is="tab"
        @popForm="popForm"
      ></component>
      <input type="button" @click="submitForm()" value="登録">
  </form>
</template>
