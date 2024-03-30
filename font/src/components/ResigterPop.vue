<template>
  <el-dialog v-model="props.dialogVisible" title="注册账号" width="500">
    <el-form
      ref="ruleFormRef"
      class="w-full"
      :model="ruleForm"
      status-icon
      :rules="rules"
      label-width="auto"
    >
      <el-form-item label="账号" prop="account">
        <el-input v-model="ruleForm.account" />
      </el-form-item>
      <el-form-item label="昵称" prop="nickname">
        <el-input v-model="ruleForm.nickname" />
      </el-form-item>
      <el-form-item label="密码" prop="pass">
        <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
      </el-form-item>
      <el-form-item label="确认密码" prop="checkPass">
        <el-input
          v-model="ruleForm.checkPass"
          type="password"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item class="flex flex-col-reverse">
        <el-button type="primary" @click="submitForm(ruleFormRef)"
          >确认</el-button
        >
        <el-button @click="resetForm(ruleFormRef)">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import axios from "axios";
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from "element-plus";

const ruleFormRef = ref<FormInstance>();

const props = defineProps<{
  dialogVisible?: boolean;
}>();

const emit = defineEmits(["cancel"]);

const validatePass = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入密码"));
  } else {
    if (ruleForm.checkPass !== "") {
      if (!ruleFormRef.value) return;
      ruleFormRef.value.validateField("checkPass", () => null);
    }
    callback();
  }
};
const validatePass2 = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请再次输入密码"));
  } else if (value !== ruleForm.pass) {
    callback(new Error("两次输入不一致!"));
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

const validateNick = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("昵称不能为空"));
  } else {
    callback();
  }
};

const ruleForm = reactive({
  account: "",
  nickname: "",
  pass: "",
  checkPass: "",
});

const rules = reactive<FormRules<typeof ruleForm>>({
  pass: [{ validator: validatePass, trigger: "blur" }],
  checkPass: [{ validator: validatePass2, trigger: "blur" }],
  account: [{ validator: validateAccount, trigger: "blur" }],
  nickname: [{ validator: validateNick, trigger: "blur" }],
});

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid: any) => {
    if (valid) {
      let response =  register(ruleForm)
      response.catch((err)=>{
        ElMessage(err.response.data.message) 
      })
      response.then((res)=>{
        ElMessage(res.data.message)
      })
    } else {
      console.log("error submit!");
      return false;
    }
  });
  emit("cancel");
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
  emit("cancel");
};

const register = async (ruleForm: any) => {
  const params = {
    account:ruleForm.account,
    password:ruleForm.pass,
    nickname:ruleForm.nickname

  }
  const response = await axios.post("/api/register", params);
  return response
};
</script>
