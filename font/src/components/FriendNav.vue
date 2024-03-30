<template>
  <div class="h-full w-1/5 flex flex-col border-r">
    <el-input class="h-10 my-2" v-model="input" placeholder="请输入账号或昵称">
      <template #append>
        <el-dropdown trigger="click">
          <span style="font-size: 22px">+</span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="searchFriend(true)"
                >通过账号添加</el-dropdown-item
              >
              <el-dropdown-item @click="searchFriend(false)"
                >通过昵称添加</el-dropdown-item
              >
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
    </el-input>
    <el-scrollbar>
      <div
        style="width: 90%"
        class="h-14 border flex items-center justify-center m-2"
        v-for="item in friendsList"
        :key="item"
      >
        {{ item.friend_nickname }}
      </div>
    </el-scrollbar>
    <el-dialog v-model="dialogVisible" title="用户列表" width="500">
      <div v-for="item in searchedFriends" :key="item.id" class="border">
        <div
          class="flex justify-between items-center px-8"
          style="font-size: 24px"
        >
          <el-icon><UserFilled /></el-icon>
          <div style="font-size: 14px">{{ item.nickname }}</div>
          <div style="font-size: 14px">{{ item.account }}</div>
          <el-icon @click="(dialogVisible = false), addFriend(item)"
            ><Select
          /></el-icon>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import service from "@/api/request";
import { ElMessage } from "element-plus";
import { ref } from "vue";

const input = ref("");
const dialogVisible = ref(false);
const searchedFriends = ref();
const friendsList = ref()
const getFriends = async () => {
  await service.post("get-friends", {
    userid: localStorage.getItem("userid"),
  })
  .then((res)=>{
    friendsList.value = res.data.data
  });
};
getFriends()

const searchFriend = async (b: boolean) => {
  if (input.value == "") {
    ElMessage("请先输入");
    return;
  }
  if (b) {
    await service
      .post("/search-friend", {
        account: input.value,
      })
      .then((res) => {
        searchedFriends.value = res.data.data;
        dialogVisible.value = true;
        input.value = "";
      });
  } else {
    await service.post("/search-friend", {
      nickname: input.value,
    });
  }
};
const addFriend = async (params: any) => {
  await service.post("add-friend", {
    user_id: Number(localStorage.getItem("userid")),
    friend_id: params.id,
    friend_account: params.account,
    friend_nickname: params.nickname,
  }).then(()=>{
    getFriends()
  });
};


</script>

<style scoped></style>
