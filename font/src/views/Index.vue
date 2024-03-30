<template>
  <ResigterPop :dialogVisible="dialogVisible" @cancel="dialogVisible = false">
  </ResigterPop>
  <div class="w-full h-full flex flex-col items-center justify-center">
    <div class="mb-6" style="font-size: 25px">请登陆</div>
    <el-form
      ref="ruleFormRef"
      class="w-1/4"
      :model="ruleForm"
      status-icon
      :rules="rules"
      label-width="auto"
    >
      <el-form-item label="账号" prop="account">
        <el-input v-model="ruleForm.account" />
      </el-form-item>
      <el-form-item label="密码" prop="pass">
        <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
      </el-form-item>
      <el-form-item class="flex flex-col-reverse">
        <el-button type="primary" @click="submitForm(ruleFormRef)"
          >确认</el-button
        >
        <el-button @click="pop()">注册</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import ResigterPop from "@/components/ResigterPop.vue";
import { reactive, ref } from "vue";
import { ElMessage, type FormInstance, type FormRules } from "element-plus";
import { useRouter } from "vue-router";
import axios from "axios";
import { saveLocalUser } from "@/utils/localutils";

const ruleFormRef = ref<FormInstance>();
const router = useRouter();

const validatePass = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入密码"));
  } else {
    callback();
  }
};

const validateAccount = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("账号不能为空"));
  } else {
    callback();
  }
};

const ruleForm = reactive({
  account: "",
  pass: "",
});

const rules = reactive<FormRules<typeof ruleForm>>({
  pass: [{ validator: validatePass, trigger: "blur" }],
  account: [{ validator: validateAccount, trigger: "blur" }],
});

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid: any) => {
    if (valid) {
      let response = login(ruleForm);
      response.catch((err) => {
        ElMessage(err.response.data.message);
      });
      response.then(async (res) => {
        await axios
          .post("/api/change-status", {
            user_id: parseInt(res.data.data.userid),
            status: 1,
          })
          .then(() => {
            saveLocalUser(
              res.data.data.token,
              res.data.data.userid,
              ruleForm.account,
              res.data.data.nickname
            );
            router.push("chat");
            ElMessage(res.data.message);
          });
      });
    } else {
      console.log("error submit!");
      return false;
    }
  });
};

const dialogVisible = ref(false);

const pop = () => {
  dialogVisible.value = true;
};
const login = async (ruleForm: any) => {
  const params = {
    account: ruleForm.account,
    password: ruleForm.pass,
  };
  const response = await axios.post("/api/login", params);
  return response;
};
</script>

<style scoped></style>
