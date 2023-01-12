<template>
  <router-view/>asda
</template>

<script>
  export default {
    name: 'App',
    mounted () {
      window.notice = this.$Notice
    }
  }
</script>

<style lang="less">
  .size{
    width: 100%;
    height: 100%;
  }
  html,body{
    .size;
    overflow: hidden;
    margin: 0;
    padding: 0;
  }
  #app {
    .size;
  }
</style>
