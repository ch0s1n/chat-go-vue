<template>
  <div class="w-full h-full">
    <Header></Header>
    <div class="w-full h-full flex border-b-2">
      <FriendNav></FriendNav>
      <div class="h-full w-full flex flex-col">
        <Talk></Talk> 
        <Footer></Footer>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import service from "@/api/request";
import Footer from "@/components/Footer.vue";
import FriendNav from "@/components/FriendNav.vue";
import Header from "@/components/Header.vue";
import Talk from "@/components/Talk.vue";

const getFriends = async() =>{
  return await service.post("/get-friends",{
    "userid":localStorage.getItem("userid")
  })

}


</script>





<style scoped></style>
