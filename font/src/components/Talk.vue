<template>
  <div class="h-4/6 w-full">
    <el-scrollbar class="pb-5">
      <div
        v-for="item in chatArray"
        :key="item.id"
        class="w-full border flex px-5"
        :class="[isSelf(item)?'flex-row':'flex-row-reverse']"
      >
       <div class="flex flex-col">
        <div class="w-full flex gap-5">
          <span
            style="font-size: 8px; color: green"
            class="flex items-center"
            >{{ item.time }}</span
          >
          <span
            style="font-size: 16px; color: blueviolet"
            class="flex items-center"
            >{{ item.nickname }}</span
          >
        </div>
        <div class="pl-3" style="width: 400px; word-break: break-all;"      
        >
          {{ item.content }}
        </div>
       </div>
        
      </div>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { isReturnStatement } from "typescript";
import { reactive, ref } from "vue";

interface chatObj {
  id: number;
  userid: number;
  nickname: string;
  time: Date;
  content: string;
}

const chatArray = ref<chatObj[]>();
const isSelf = (item: chatObj) => {
  return item.userid == 111;
};

chatArray.value = [
  {
    id: 1,
    userid: 111,
    nickname: "lzy",
    time: new Date("2023-04-01T12:30:00"),
    content: "sadf",
  },
  {
    id: 2,
    userid: 111,
    nickname: "lzy",
    time: new Date("2023-04-01T12:31:00"),
    content: "world",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 3,
    userid: 222,
    nickname: "zhangsan",
    time: new Date("2023-04-01T12:32:00"),
    content:
      "sadkjfasdjkfhjdsahfjsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshfsdafjhjksdhafjkdshf",
  },
  {
    id: 1,
    userid: 111,
    nickname: "lzy",
    time: new Date("2023-04-01T12:30:00"),
    content: "sadf",
  },
];


</script>

<style scoped>
.self {
  left: 100%;
}
</style>
