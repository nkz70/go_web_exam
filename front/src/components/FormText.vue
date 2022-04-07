<script setup>
import { ref } from 'vue';

const props = defineProps({
    form: Object
});

const form = ref(props.form)

const emit = defineEmits([
    "popForm",
    "updateForm"
]);

</script>

<template>
    <button @click="emit('popForm', id)">削除</button>
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
</template>
