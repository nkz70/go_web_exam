<script setup>
import { reactive, ref } from 'vue';

const props = defineProps({
    form: Object
})

const emit = defineEmits(['popForm'])

const form = ref(props.form)

const sendPopForm = () => emit('popForm', form.value.id)

</script>

<template>
    <button @click="sendPopForm">削除</button>
    <div>
        <label for="">ラベル位置</label>
        <input type="text" v-model.number="form.label_position">
    </div>
    <div>
        <label for="">ラベル</label>
        <input type="text" v-model="form.label">
    </div>
    <div>
        <label for="">答え</label>
        <input type="text" v-model="form.answer">
    </div>
    <div>
        <label for="">選択肢</label>
        <input v-for="(choice, index) in form.choices" type="text" v-model="form.choices[index]">
    </div>
</template>
