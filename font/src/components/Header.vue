<template>
  <div
    class="bg-white h-11 w-full flex items-center justify-between px-8 border-b-2"
  >
    <div>欢迎来聊天</div>
    <div>1111</div>
    <el-dropdown trigger="click">
      <div style="font-size: 23px">
        <el-icon><User /></el-icon>
      </div>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item @click="dialogVisible = true"
            >退出登陆</el-dropdown-item
          >
          <el-dropdown-item>其他功能...待开发</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>

  <el-dialog v-model="dialogVisible" width="500">
    <span>确认退出吗</span>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="(dialogVisible = false), logout()">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { deleteLocalUser } from "@/utils/localutils";
import axios from "axios";
import { ElMessage } from "element-plus";
import { ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();

const dialogVisible = ref(false);
const logout = async () => {
  await axios
    .post("/api/change-status", {
      user_id: Number(localStorage.getItem("userid")),
      status: 0,
    })
    .then(() => {
      deleteLocalUser();
      router.push("/");
      ElMessage("退出登陆");
    });
};
</script>

<style scoped></style>
