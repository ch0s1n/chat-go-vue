<template>
  <div class="bg-white w-full border-t-2 h-2/5">
    <textarea
      v-model="inputtext"
      placeholder="请输入内容....Enter 发送,Ctrl + Enter换行" 
      class="h-full w-full p-2 px-4"
      @keydown.enter.native.prevent="handleEnterKey($event)" 
    >
        {{ inputtext }}
    </textarea>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";


const inputtext = ref();
const handleEnterKey =(e:any) => {
    if(e.ctrlKey && e.keyCode==13) { //用户点击了ctrl+enter触发  
       inputtext.value += '\n';
    }else { //用户点击了enter触发  
       sendMessage();
    } 
}

const sendMessage=()=>{
    inputtext.value = ''
}

</script>

<style scoped></style>
